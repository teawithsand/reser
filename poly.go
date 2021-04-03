package reser

// PolySerializer handles polymorphic serialization without any external hints about type of data serialized.
type PolySerializer interface {
	PolySerialize(data interface{}) (res []byte, err error)
}

// PolyDeserializer handles polymprphic deserialization without any external hints about type of data deserialized
// provided by caller.
//
// Note: This deserializer always returns struct values rather than pointer values.
type PolyDeserializer interface {
	PolyDeserialize(data []byte) (res interface{}, err error)
}
