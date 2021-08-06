package reser_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/teawithsand/reser"
)

func makeSTPolySerializer(ty reflect.Type) *reser.STPolySerializer {
	return &reser.STPolySerializer{
		Type:         ty,
		Serializer:   reser.SerializerFunc(json.Marshal),
		Deserializer: reser.DeserializerFunc(json.Unmarshal),
	}
}

func makeFacSTPolySerializer(fac func() interface{}) *reser.STPolySerializer {
	return &reser.STPolySerializer{
		TypeFactory:  fac,
		Serializer:   reser.SerializerFunc(json.Marshal),
		Deserializer: reser.DeserializerFunc(json.Unmarshal),
	}
}

func Test_PolySingleSerializer_CanSerialize(t *testing.T) {
	serializer := makeSTPolySerializer(reflect.TypeOf(pt1{}))
	v := pt1{
		ValOne: 42,
	}
	res, err := serializer.PolySerialize(v)
	if err != nil {
		t.Error(err)
		return
	}

	v2, err := serializer.PolyDeserialize(res)
	if err != nil {
		t.Error(err)
		return
	}

	if v != v2.(pt1) {
		t.Error("Value mismatch")
	}
}

func Test_PolySingleSerializer_CanSerializePtr(t *testing.T) {
	serializer := makeSTPolySerializer(reflect.TypeOf(&pt1{}))
	v := &pt1{
		ValOne: 42,
	}
	res, err := serializer.PolySerialize(v)
	if err != nil {
		t.Error(err)
		return
	}

	v2, err := serializer.PolyDeserialize(res)
	if err != nil {
		t.Error(err)
		return
	}

	if *v != *(v2.(*pt1)) {
		t.Error("Value mismatch")
	}
}

func Test_PolySingleSerializer_CanSerializeTypeFactoryPtr(t *testing.T) {
	serializer := makeFacSTPolySerializer(func() interface{} {
		return &pt1{}
	})
	v := &pt1{
		ValOne: 42,
	}
	res, err := serializer.PolySerialize(v)
	if err != nil {
		t.Error(err)
		return
	}

	v2, err := serializer.PolyDeserialize(res)
	if err != nil {
		t.Error(err)
		return
	}

	if *v != *(v2.(*pt1)) {
		t.Error("Value mismatch")
	}
}
