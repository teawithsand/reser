package reser

import "reflect"

// STPolySerializer is poly serializer and deserializer, which is able to serialize only single type.
type STPolySerializer struct {
	Type         reflect.Type
	Serializer   Serializer
	Deserializer Deserializer
}

func (ser *STPolySerializer) PolySerialize(data interface{}) (res []byte, err error) {
	// Just assume that type is compatibile with what's given
	/*
		typeValid := false
		if data != nil {
			ty := reflect.TypeOf(data)
			if !reflect.ValueOf(data).IsZero() && ty.Kind() == reflect.Ptr && ty.Elem() == ser.Type {
				typeValid = true
			}
			if reflect.TypeOf(data) != ser.Type {
				typeValid = true
			}
		}
		if !typeValid {
			err = &UnsupportedValueError{
				Val: data,
			}
			return
		}
	*/

	return ser.Serializer.Serialize(data)
}

func (ser *STPolySerializer) PolyDeserialize(data []byte) (res interface{}, err error) {
	v := reflect.New(ser.Type).Interface()

	err = ser.Deserializer.Deserialize(data, v)
	if err != nil {
		return
	}

	res = reflect.ValueOf(v).Elem().Interface()
	return
}
