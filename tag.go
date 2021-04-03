package reser

import (
	"reflect"
)

type TypeTag interface{}

// TagSerializer uses external information(so called tag) to make polymorphic marshalling work.
type TagSerializer interface {
	GetTypeTag(data interface{}) (tt TypeTag, err error)
	Serialize(data interface{}) (res []byte, err error)
}

// TagDeserializer uses tag and data in order to deserialize actual type.
//
// Note: This deserializer always returns struct pointers rather than struct values.
type TagDeserializer interface {
	Deserialize(data []byte, tag TypeTag) (res interface{}, err error)
}

// DefaultTagSerializer is default implementation of TagSerializer
// and TagDeserializer.
type DefaultTagSerializer struct {
	Serializer       Serializer
	Deserializer     Deserializer
	TagTypeResgistry *TagTypeResgistry

	// TODO(teaiwthsand): add lock to this type?
}

func (ds *DefaultTagSerializer) GetTypeTag(data interface{}) (tt TypeTag, err error) {
	ty := reflect.TypeOf(data)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}
	tt, err = ds.TagTypeResgistry.GetTag(ty)
	return
}

func (ds *DefaultTagSerializer) Serialize(event interface{}) (res []byte, err error) {
	f := ds.Serializer
	return f.Serialize(event)
}

func (ds *DefaultTagSerializer) Deserialize(data []byte, tt TypeTag) (res interface{}, err error) {
	f := ds.Deserializer

	ty, err := ds.TagTypeResgistry.GetType(tt)
	if err != nil {
		return
	}

	rawRes := reflect.New(ty).Interface()
	err = f.Deserialize(data, rawRes)
	if err != nil {
		return
	}
	res = reflect.ValueOf(rawRes).Elem().Interface()
	return
}
