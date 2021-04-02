package reser

import "reflect"

type TagRegisterError struct {
	Type    reflect.Type
	TypeTag TypeTag
}

func (e *TagRegisterError) Error() string {
	return "specified type tag is in use"
}

type TypeNotFoundError struct {
	TypeTag TypeTag
}

func (e *TypeNotFoundError) Error() string {
	return "specified type tag was not found"
}

type TagNotFoundError struct {
	Type reflect.Type
}

func (e *TagNotFoundError) Error() string {
	return "specified type was not found"
}

type TypeTagTypeError struct {
	ExpectedType reflect.Type
	Tag          TypeTag
}

func (e *TypeTagTypeError) Error() string {
	return "given type tag has invalid type for given registry"
}
