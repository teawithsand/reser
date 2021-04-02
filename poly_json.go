package reser

import (
	"errors"
	"reflect"
)

// ETPolySerializer uses external field called `type`.
// Actual data is stored in `data` json field.
// This approach is known as adjacent tagging in most(?) serialization frameworks(for instance serde_json).
// It works with serializers like JSON one build in golang stdlib.
//
// Note: in order for this to work deserializer must accept unknown fields in data, otherwise it won't work.
// Note#2: It's implementation is somwhat magic and it's expected to work only against JSON serializer in go stdlib.
// No unit testing is performed against any other serialization implementation.
type ETPolySerializer struct {
	Serializer       Serializer
	Deserializer     Deserializer
	TagTypeResgistry TagTypeResgistry
}

func (ets *ETPolySerializer) PolySerialize(data interface{}) (res []byte, err error) {
	tt, err := ets.TagTypeResgistry.GetTag(reflect.TypeOf(data))
	if err != nil {
		return
	}

	// TODO(teawithsand): add check here, so that serialization would fail before deserialization
	//  check about whether or not is data struct or at least struct pointer, which is not nil

	return ets.Serializer.Serialize(adjacentTagSerialize{
		Type: tt,
		Data: data,
	})
}

func (ets *ETPolySerializer) PolyDeserializer(data []byte) (res interface{}, err error) {
	ttt := ets.TagTypeResgistry.GetTypeTagType()
	var tagContainerType tagContainer
	if ttt.Kind() == reflect.Struct {
		tagContainerType = &adjacentTag{
			Type: reflect.New(ttt),
		}
	} else if ttt.Kind() == reflect.Ptr && ttt.Elem().Kind() == reflect.Struct {
		tagContainerType = &adjacentTag{
			Type: reflect.New(ttt.Elem()),
		}
	} else if ttt.Kind() == reflect.String {
		tagContainerType = &adjacentTagString{}
	} else if ttt.Kind() == reflect.Uint {
		tagContainerType = &adjacentTagUInt{}
	} else if ttt.Kind() == reflect.Uint8 {
		tagContainerType = &adjacentTagUInt8{}
	} else if ttt.Kind() == reflect.Uint16 {
		tagContainerType = &adjacentTagUInt16{}
	} else if ttt.Kind() == reflect.Uint32 {
		tagContainerType = &adjacentTagUInt32{}
	} else if ttt.Kind() == reflect.Uint64 {
		tagContainerType = &adjacentTagUInt64{}
	} else if ttt.Kind() == reflect.Int {
		tagContainerType = &adjacentTagInt{}
	} else if ttt.Kind() == reflect.Int8 {
		tagContainerType = &adjacentTagInt8{}
	} else if ttt.Kind() == reflect.Int16 {
		tagContainerType = &adjacentTagInt16{}
	} else if ttt.Kind() == reflect.Int32 {
		tagContainerType = &adjacentTagInt32{}
	} else if ttt.Kind() == reflect.Int64 {
		tagContainerType = &adjacentTagInt64{}
	} else {
		err = errors.New("invalid type provided as tag type")
		return
		// TODO(teawithsand): beter error - unsupported tag type
	}

	err = ets.Deserializer.Deserialize(data, tagContainerType)
	if err != nil {
		return
	}
	ty, err := ets.TagTypeResgistry.GetType(tagContainerType.GetTypeTag())
	if err != nil {
		return
	}
	if ty.Kind() != reflect.Struct {
		err = &UnsupportedTypeError{
			Type: ty,
		}
		return
	}
	resultData := reflect.New(ty).Interface()
	resContainer := adjacentTagDeserialize{
		Data: resultData,
	}
	err = ets.Deserializer.Deserialize(data, &resContainer)
	if err != nil {
		return
	}
	res = resContainer.Data
	return
}
