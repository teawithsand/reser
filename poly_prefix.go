package reser

import (
	"fmt"
	"io"
	"reflect"
)

type PrefixPolyEncoderFactory[T TypeTag, D any] struct {
	TagRegistry TypeTagRegistry[T]

	TagEncoderFactory  func(w io.Writer) (e PolyEncoder[T], err error)
	DataEncoderFactory func(w io.Writer) (e Encoder, err error)
}

func (f *PrefixPolyEncoderFactory[T, D]) NewEncoder(w io.Writer) (e PolyEncoder[D], err error) {
	tagEncoder, err := f.TagEncoderFactory(w)
	if err != nil {
		return
	}

	dataEncoder, err := f.DataEncoderFactory(w)
	if err != nil {
		return
	}

	e = &prefixPolyEncoder[T, D]{
		tagEncoder:  tagEncoder,
		dataEncoder: dataEncoder,
		tagRegistry: f.TagRegistry,
	}

	return
}

type prefixPolyEncoder[T TypeTag, D any] struct {
	tagEncoder  PolyEncoder[T]
	dataEncoder Encoder
	tagRegistry TypeTagRegistry[T]
}

func (ppe *prefixPolyEncoder[T, D]) PolyEncode(data D) (err error) {
	tag, ok := ppe.tagRegistry.GetTagForType(reflect.TypeOf(data))
	if !ok {
		err = fmt.Errorf("reser: type %T is not registered in tag registry given", data)
		return
	}

	err = ppe.tagEncoder.PolyEncode(tag)
	if err != nil {
		return
	}

	err = ppe.dataEncoder.Encode(data)
	if err != nil {
		return
	}

	return
}

type PrefixPolyDecoderFactroy[T TypeTag, D any] struct {
	TagRegistry TypeTagRegistry[T]

	// Note: this decoder must not read more bytes than required from reader given.
	TagDecoderFactory func(r io.Reader) (e PolyDecoder[T], err error)

	// Note: this decoder must not read more bytes than required from reader given.
	DataDecoderFactory func(r io.Reader) (d Decoder, err error)
}

func (f *PrefixPolyDecoderFactroy[T, D]) NewDecoder(r io.Reader) (e PolyDecoder[D], err error) {
	tagDecoder, err := f.TagDecoderFactory(r)
	if err != nil {
		return
	}

	dataDecoder, err := f.DataDecoderFactory(r)
	if err != nil {
		return
	}

	e = &prefixPolyDecoder[T, D]{
		tagDecoder:  tagDecoder,
		dataDecoder: dataDecoder,
		tagRegistry: f.TagRegistry,
	}

	return
}

type prefixPolyDecoder[T TypeTag, D any] struct {
	tagDecoder  PolyDecoder[T]
	dataDecoder Decoder
	tagRegistry TypeTagRegistry[T]
}

func (ppe *prefixPolyDecoder[T, D]) PolyDecode() (data D, err error) {
	tag, err := ppe.tagDecoder.PolyDecode()
	if err != nil {
		return
	}

	ty := ppe.tagRegistry.GetTypeForTag(tag)
	if ty == nil {
		err = fmt.Errorf("reser: tag %+#v is not registered in tag registry given", tag)
		return
	}

	res := reflect.New(ty).Interface()

	err = ppe.dataDecoder.Decode(res)
	if err != nil {
		return
	}

	data = *res.(*D)
	return
}
