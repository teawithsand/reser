package reser_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/teawithsand/reser"
)

type pt1 struct {
	ValOne int
}
type pt2 struct {
	ValTwo int
}
type pt3 struct {
	ValThere int
}

func makeSerializer() *reser.ETPolySerializer {
	ttr := reser.NewTagTypeResgistry(reflect.TypeOf(""))
	ttr.RegisterType(reflect.TypeOf(pt1{}), "1")
	ttr.RegisterType(reflect.TypeOf(pt2{}), "2")
	ttr.RegisterType(reflect.TypeOf(pt3{}), "3")

	return &reser.ETPolySerializer{
		Serializer:       reser.SerializerFunc(json.Marshal),
		Deserializer:     reser.DeserializerFunc(json.Unmarshal),
		TagTypeResgistry: ttr,
	}
}

func Test_ETPolySerializer_CanSerializeDeserialize(t *testing.T) {
	v := pt1{
		ValOne: 42,
	}
	s := makeSerializer()
	data, err := s.PolySerialize(v)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := s.PolyDeserialize(data)
	if err != nil {
		t.Error(err)
		return
	}
	nv := res.(*pt1)
	if *nv != v {
		t.Error("Value mismatch")
		return
	}
}
