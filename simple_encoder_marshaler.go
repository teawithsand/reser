package reser

import (
	"bytes"
	"io"
)

type EncoderMarshaler struct {
	EncoderFactory func(w io.Writer) (Encoder, error)
}

func (em *EncoderMarshaler) Marshal(data any) (res []byte, err error) {
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

type DecoderUnmarshaler struct {
	DecoderFactory func(r io.Reader) (Decoder, error)
}

func (em *DecoderUnmarshaler) Unmarshal(data []byte, res any) (err error) {
	d, err := em.DecoderFactory(bytes.NewReader(data))
	if err != nil {
		return
	}
	err = d.Decode(data)
	if err != nil {
		return
	}
	return
}
