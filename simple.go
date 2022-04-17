package reser

// Marshaler, which fits most interfaces exposed by golang already.
// For instance, json.Marshal with MarshalerFunc.
type Marshaler interface {
	Marshal(data interface{}) (res []byte, err error)
}

type MarshalerFunc func(data interface{}) (res []byte, err error)

func (f MarshalerFunc) Marshal(data interface{}) (res []byte, err error) {
	return f(data)
}

// Unmarshaler, which fits most interfaces exposed by golang already.
// For instance, json.Unmarshal with MarshalerFunc.
type Unmarshaler interface {
	Unmarshal(data []byte, dst interface{}) (err error)
}

type UnmarshalerFunc func(data []byte, dst interface{}) (err error)

func (f UnmarshalerFunc) Unmarshal(data []byte, dst interface{}) (err error) {
	return f(data, dst)
}
