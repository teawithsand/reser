package reser

import (
	"encoding/json"
	"reflect"
)

// ETPolySerializer uses external field called `Type`.
// Actual data is stored in `Data` json field.
// This approach is known as adjacent tagging in most(?) serialization frameworks(for instance serde_json).
// It works with serializers like JSON one build in golang stdlib.
//
// Note: in order for this to work deserializer must accept unknown fields in data, otherwise it won't work.
// Note#2: It's implementation is somwhat magic and it's expected to work only against JSON serializer in go stdlib.
// No unit testing is performed against any other serialization implementation.
type ETPolySerializer struct {
	Serializer      Serializer
	Deserializer    Deserializer
	TagTypeRegistry *TypeTagRegistry
}

func (ets *ETPolySerializer) getSerializer() (ser Serializer) {
	ser = ets.Serializer
	if ser == nil {
		ser = SerializerFunc(json.Marshal)
	}
	return
}
func (ets *ETPolySerializer) getDeserializer() (des Deserializer) {
	des = ets.Deserializer
	if des == nil {
		des = DeserializerFunc(json.Unmarshal)
	}
	return
}

func (ets *ETPolySerializer) PolySerialize(data interface{}) (res []byte, err error) {
	tt, err := ets.TagTypeRegistry.GetTag(reflect.TypeOf(data))
	if err != nil {
		return
	}

	// TODO(teawithsand): add check here, so that serialization would fail before deserialization
	//  check about whether or not is data struct or at least struct pointer, which is not nil

	return ets.getSerializer().Serialize(adjacentTagSerialize{
		Type: tt,
		Data: data,
	})
}

func (ets *ETPolySerializer) PolyDeserialize(data []byte) (res interface{}, err error) {
	ttt := ets.TagTypeRegistry.GetTypeTagType()
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
	} else if ttt.Kind() == reflect.Slice && ttt.Elem().Kind() == reflect.Uint8 {
		tagContainerType = &adjacentTagBytes{}
	} else {
		err = &UnsupportedTypeTagTypeError{
			Type: ttt,
		}
		return
	}

	err = ets.getDeserializer().Deserialize(data, tagContainerType)
	if err != nil {
		return
	}
	ty, err := ets.TagTypeRegistry.GetType(tagContainerType.GetTypeTag())
	if err != nil {
		return
	}
	/*
		if ty.Kind() != reflect.Struct {
			err = &UnsupportedTypeError{
				Type: ty,
			}
			return
		}
	*/
	resultData := reflect.New(ty).Interface()
	resContainer := adjacentTagDeserialize{
		Data: resultData,
	}
	err = ets.getDeserializer().Deserialize(data, &resContainer)
	if err != nil {
		return
	}
	res = reflect.ValueOf(resContainer.Data).Elem().Interface()
	return
}
