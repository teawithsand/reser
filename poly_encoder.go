package reser

type PolyEncoder[T any] interface {
	PolyEncode(data T) (err error)
}

type PolyEncoderFunc[T any] func(data T) (err error)

func (f PolyEncoderFunc[T]) PolyEncode(data T) (err error) {
	return f(data)
}

type PolyDecoder[T any] interface {
	PolyDecode() (res T, err error)
}

type PolyDecoderFunc[T any] func() (res T, err error)

func (f PolyDecoderFunc[T]) PolyDecode() (res T, err error) {
	return f()
}
