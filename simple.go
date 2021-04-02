package reser

// Serializer, which fits most interfaces exposed by golang already.
// For instance, json.Marshal with SerializerFunc.
type Serializer interface {
	Serialize(data interface{}) (res []byte, err error)
}
type SerializerFunc func(data interface{}) (res []byte, err error)

func (f SerializerFunc) Serialize(data interface{}) (res []byte, err error) {
	return f(data)
}

// Deserializer, which fits most interfaces exposed by golang already.
// For instance, json.Unmarshal with SerializerFunc.
type Deserializer interface {
	Deserialize(data []byte, dst interface{}) (err error)
}
type DeserializerFunc func(data []byte, dst interface{}) (err error)

func (f DeserializerFunc) Deserialize(data []byte, dst interface{}) (err error) {
	return f(data, dst)
}
