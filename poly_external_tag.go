package reser

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

type ExternalTagPolyEncoderFactory[T TypeTag, D any] struct {
	WrapperFactory  WrapperFactory[T, D]
	TypeTagRegistry TypeTagRegistry[T]
	EncoderFactory  func(w io.Writer) (Encoder, error)
}

func (f *ExternalTagPolyEncoderFactory[T, D]) NewEncoder(w io.Writer) (e PolyEncoder[D], err error) {
	enc, err := f.EncoderFactory(w)
	if err != nil {
		return
	}

	e = &taggedDataPolyEncoder[T, D]{
		typeTagRegistry: f.TypeTagRegistry,
		wrapperFactory:  f.WrapperFactory,
		encoder:         enc,
	}
	return
}

type taggedDataPolyEncoder[T TypeTag, D any] struct {
	typeTagRegistry TypeTagRegistry[T]
	wrapperFactory  WrapperFactory[T, D]
	encoder         Encoder
}

func (e *taggedDataPolyEncoder[T, D]) PolyEncode(data D) (err error) {
	tag, ok := e.typeTagRegistry.GetTagForType(reflect.TypeOf(data))
	if !ok {
		err = fmt.Errorf("reser: type %T is not registered in tag registry given", data)
		return
	}

	w := e.wrapperFactory.NewTagDataWrapper()
	w.SetData(data)
	w.SetTag(tag)

	err = e.encoder.Encode(w)
	if err != nil {
		return
	}
	return
}

type ExternalTagPolyDecoderFactory[T TypeTag, D any] struct {
	TypeTagRegistry TypeTagRegistry[T]
	WrapperFactory  WrapperFactory[T, D]
	DecoderFactory  func(r io.Reader) (Decoder, error)
}

func (f *ExternalTagPolyDecoderFactory[T, D]) NewDecoder(r io.Reader) (e PolyDecoder[D], err error) {
	e = &taggedDataPolyDecoder[T, D]{
		typeTagRegistry:   f.TypeTagRegistry,
		wrapperFactory:    f.WrapperFactory,
		tagDecoderFactory: f.DecoderFactory,
	}
	return
}

// Note: this decoder does not decode data in streamming way.
// It buffers all data before decoding, since information about type may be anywhere in data structure.
// So first it has to decode data for tag, to determine its type and then  it has to decode for data itself.
type taggedDataPolyDecoder[T TypeTag, D any] struct {
	typeTagRegistry    TypeTagRegistry[T]
	wrapperFactory     WrapperFactory[T, D]
	tagDecoderFactory  func(r io.Reader) (Decoder, error)
	dataDecoderFactory func(r io.Reader) (Decoder, error)
	r                  io.Reader
}

func (d *taggedDataPolyDecoder[T, D]) PolyDecode() (data D, err error) {
	// HACK(teawithsand): use TeeReader in order to memorize all data decoded.
	// This should not be used if reader incoming is seekable and just could be rewinded.
	// OR if input is just buffer of bytes.
	// Then no need to copy them one more time

	b := bytes.NewBuffer(nil)
	tagReader := io.TeeReader(d.r, b)

	tagDecoder, err := d.tagDecoderFactory(tagReader)
	if err != nil {
		return
	}

	tagWrapper := d.wrapperFactory.NewTagWrapper()
	err = tagDecoder.Decode(tagWrapper)
	if err != nil {
		return
	}

	ty := d.typeTagRegistry.GetTypeForTag(tagWrapper.GetTag())
	if ty == nil {
		err = fmt.Errorf("reser: tag %+#v is not registered in tag registry given", tagWrapper.GetTag())
		return
	}

	dataWrapper := d.wrapperFactory.NewDataWrapper()
	dataWrapper.SetData(reflect.New(ty).Elem().Interface().(D))

	dataDecoder, err := d.dataDecoderFactory(io.MultiReader(bytes.NewReader(b.Bytes()), d.r))
	if err != nil {
		return
	}

	err = dataDecoder.Decode(&dataWrapper)
	if err != nil {
		return
	}

	data = dataWrapper.GetData()
	return
}
