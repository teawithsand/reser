package reser

// Marshaler, which fits most interfaces exposed by golang already.
// For instance, json.Marshal with SerializerFunc.
type Marshaler interface {
	Serialize(data interface{}) (res []byte, err error)
}

type MarshalerFunc func(data interface{}) (res []byte, err error)

func (f MarshalerFunc) Serialize(data interface{}) (res []byte, err error) {
	return f(data)
}

// Unmarshaler, which fits most interfaces exposed by golang already.
// For instance, json.Unmarshal with SerializerFunc.
type Unmarshaler interface {
	Deserialize(data []byte, dst interface{}) (err error)
}

type UnmarshalerFunc func(data []byte, dst interface{}) (err error)

func (f UnmarshalerFunc) Deserialize(data []byte, dst interface{}) (err error) {
	return f(data, dst)
}
