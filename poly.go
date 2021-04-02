package reser

// PolySerializer handles polymorphic serialization without any external hints about type of data serialized.
type PolySerializer interface {
	PolySerialize(data interface{}) (res []byte, err error)
}

// PolyDeserializer handles polymprphic deserialization without any external hints about type of data deserialized
// provided by caller.
//
// Note: This deserializer always returns struct pointers rather than struct values.
type PolyDeserializer interface {
	PolyDeserializer(data []byte) (res interface{}, err error)
}
