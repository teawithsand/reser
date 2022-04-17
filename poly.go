package reser

// PolySerializer handles polymorphic serialization without any external hints about type of data serialized.
type PolySerializer[T any] interface {
	PolySerialize(data T) (res []byte, err error)
}

type PolySerializerFunc[T any] func(data T) (res []byte, err error)

func (f PolySerializerFunc[T]) PolySerialize(data T) (res []byte, err error) {
	return f(data)
}

// PolySerializer, which can process any type.
type AnyPolySerialzier = PolySerializer[any]

// PolyDeserializer handles polymprphic deserialization without any external hints about type of data deserialized
// provided by caller.
//
// Note: This deserializer always returns struct values rather than pointer values.
type PolyDeserializer[T any] interface {
	PolyDeserialize(data []byte) (res T, err error)
}

// PolyDeserializer, which returns any type.
type AnyPolyDeserializer = PolyDeserializer[any]

type PolyDeserializerFunc[T any] func(data []byte) (res T, err error)

func (f PolyDeserializerFunc[T]) PolyDeserialize(data []byte) (res T, err error) {
	return f(data)
}
