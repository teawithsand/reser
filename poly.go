package reser

// PolyMarshaler handles polymorphic serialization without any external hints about type of data serialized.
type PolyMarshaler[T any] interface {
	PolyMarshal(data T) (res []byte, err error)
}

type PolyMarshalerFunc[T any] func(data T) (res []byte, err error)

func (f PolyMarshalerFunc[T]) PolyMarshal(data T) (res []byte, err error) {
	return f(data)
}

type AnyPolyMarshal = PolyMarshaler[any]

// PolyUnmarshaler handles polymorphic deserialization without any external hints about type of data deserialized
// provided by caller.
//
// Note: This deserializer always returns struct values rather than pointer values.
type PolyUnmarshaler[T any] interface {
	PolyUnmarshal(data []byte) (res T, err error)
}

type AnyPolyUnmarshaler = PolyUnmarshaler[any]

type PolyUnmarshalerFunc[T any] func(data []byte) (res T, err error)

func (f PolyUnmarshalerFunc[T]) PolyUnmarshal(data []byte) (res T, err error) {
	return f(data)
}
