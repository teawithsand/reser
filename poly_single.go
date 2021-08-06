package reser

import "reflect"

// STPolySerializer is poly serializer and deserializer, which is able to serialize only single type.
type STPolySerializer struct {
	Type        reflect.Type       // type of data serializer
	TypeFactory func() interface{} // OR factory of types to return

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
	if ser.TypeFactory != nil {
		dst := ser.TypeFactory()
		reflectDst := reflect.TypeOf(dst)
		if reflectDst.Kind() != reflect.Ptr {
			panic("type factory must return pointer type")
		}
		err = ser.Deserializer.Deserialize(data, dst)
		if err != nil {
			return
		}

		res = dst
	} else {
		v := reflect.New(ser.Type).Interface()

		err = ser.Deserializer.Deserialize(data, v)
		if err != nil {
			return
		}

		res = reflect.ValueOf(v).Elem().Interface()
	}

	return
}
