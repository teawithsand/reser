package reser

import (
	"bytes"
	"io"
)

type polyEncoderMarshaler[T any] struct {
	encoderFactory func(w io.Writer) (PolyEncoder[T], error)
}

func (em *polyEncoderMarshaler[T]) PolyMarshal(data T) (res []byte, err error) {
	b := bytes.NewBuffer(nil)
	e, err := em.encoderFactory(b)
	if err != nil {
		return
	}
	err = e.PolyEncode(data)
	if err != nil {
		return
	}
	res = b.Bytes()
	return
}

type polyDecoderUnmarshaler[T any] struct {
	decoderFactory func(r io.Reader) (PolyDecoder[T], error)
}

func (em *polyDecoderUnmarshaler[T]) PolyUnmarshal(data []byte) (res T, err error) {
	d, err := em.decoderFactory(bytes.NewReader(data))
	if err != nil {
		return
	}

	res, err = d.PolyDecode()
	if err != nil {
		return
	}

	return
}

func PolyEncoderToMarshaler[T any](encoderFactory func(w io.Writer) (PolyEncoder[T], error)) (marshaler PolyMarshaler[T]) {
	marshaler = &polyEncoderMarshaler[T]{
		encoderFactory: encoderFactory,
	}
	return
}

func PolyDecoderToUnmarshaler[T any](decoderFactory func(r io.Reader) (PolyDecoder[T], error)) (marshaler PolyUnmarshaler[T]) {
	marshaler = &polyDecoderUnmarshaler[T]{
		decoderFactory: decoderFactory,
	}
	return
}
