package reser

type PolyEncoder[T any] interface {
	Encode(data T) (err error)
}

type PolyEncoderFunc[T any] func(data T) (err error)

func (f PolyEncoderFunc[T]) Encode(data T) (err error) {
	return f(data)
}

type PolyDecoder[T any] interface {
	Decode() (res T, err error)
}

type PolyDecoderFunc[T any] func() (res T, err error)

func (f PolyDecoderFunc[T]) Decode() (res T, err error) {
	return f()
}
