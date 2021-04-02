package reser

import "reflect"

type UnsupportedValueError struct {
	Val interface{}
}

func (uve *UnsupportedValueError) Error() string {
	if uve == nil {
		return "<nil>"
	}
	return "unsupported value provided for (de)serialization"
}

type UnsupportedTypeError struct {
	Type reflect.Type
}

func (ute *UnsupportedTypeError) Error() string {
	if ute == nil {
		return "<nil>"
	}
	return "unsupported type provided for (de)serialization"
}
