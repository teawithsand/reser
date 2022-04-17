package reser

// Encodes data of any type into some output, which is abstracted away.
type Encoder interface {
	Encode(data any) (err error)
}

type EncoderFunc func(data any) (err error)

func (f EncoderFunc) Encode(data any) (err error) {
	return f(data)
}

// Decodes data from some input, which is abstracted away.
type Decoder interface {
	Decode(data any) (err error)
}

type DecoderFunc func(data any) (err error)

func (f DecoderFunc) Decode(data any) (err error) {
	return f(data)
}
