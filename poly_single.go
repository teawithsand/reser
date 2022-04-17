package reser

import "io"

// Creates poly encoder, which is able to marshal only single type.
type SinglePolyEncoderFactory[T any] struct {
	EncoderFactory func(w io.Writer) (enc Encoder, err error)
}

func (f *SinglePolyEncoderFactory[T]) NewEncoder(w io.Writer) (e PolyEncoder[T], err error) {
	enc, err := f.EncoderFactory(w)
	if err != nil {
		return
	}

	e = &singlePolyEncoder[T]{
		enc: enc,
	}
	return
}

type singlePolyEncoder[T any] struct {
	enc Encoder
}

func (e *singlePolyEncoder[T]) PolyEncode(data T) (err error) {
	err = e.enc.Encode(data)
	if err != nil {
		return
	}
	return
}

// Creates poly decoder, which is able to decode only single type.
type SinglePolyDecoderFactory[T any] struct {
	DecoderFactory func(r io.Reader) (dec Decoder, err error)
	Factory        func() *T // must return non nil ptr to new instances of T
}

func (f *SinglePolyDecoderFactory[T]) NewDecoder(r io.Reader) (e PolyDecoder[T], err error) {
	dec, err := f.DecoderFactory(r)
	if err != nil {
		return
	}

	e = &singlePolyDecoder[T]{
		dec: dec,
		fac: f.Factory,
	}
	return
}

type singlePolyDecoder[T any] struct {
	fac func() *T
	dec Decoder
}

func (e *singlePolyDecoder[T]) PolyDecode() (data T, err error) {
	dst := e.fac()

	err = e.dec.Decode(dst)
	if err != nil {
		return
	}

	data = *dst
	return
}
