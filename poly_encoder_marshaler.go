package reser

import (
	"bytes"
	"io"
)

type PolyEncoderMarshaler[T any] struct {
	EncoderFactory func(w io.Writer) (PolyEncoder[T], error)
}

func (em *PolyEncoderMarshaler[T]) PolyMarshal(data T) (res []byte, err error) {
	b := bytes.NewBuffer(nil)
	e, err := em.EncoderFactory(b)
	if err != nil {
		return
	}
	err = e.Encode(data)
	if err != nil {
		return
	}
	res = b.Bytes()
	return
}

type PolyDecoderUnmarshaler[T any] struct {
	DecoderFactory func(r io.Reader) (PolyDecoder[T], error)
}

func (em *PolyDecoderUnmarshaler[T]) PolyUnmarshal(data []byte) (res T, err error) {
	d, err := em.DecoderFactory(bytes.NewReader(data))
	if err != nil {
		return
	}

	res, err = d.Decode()
	if err != nil {
		return
	}

	return
}
