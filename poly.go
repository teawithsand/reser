package reser

// PolySerializer handles polymorphic serialization without any external hints about type of data serialized.
type PolySerializer interface {
	PolySerialize(data interface{}) (res []byte, err error)
}

// TODO(teawithsand): cleanup return type here in type registry with additonal options

// PolyDeserializer handles polymprphic deserialization without any external hints about type of data deserialized
// provided by caller.
//
// Note: This deserializer always returns struct values rather than pointer values.
type PolyDeserializer interface {
	PolyDeserialize(data []byte) (res interface{}, err error)
}
