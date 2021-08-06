package reser_test

import (
	"encoding/json"
	"testing"

	"github.com/teawithsand/reser"
)

func makeETPolySerializer(ttr *reser.TypeTagRegistry) *reser.ETPolySerializer {
	if ttr == nil {
		ttr = makeTTR()
	}

	return &reser.ETPolySerializer{
		Serializer:      reser.SerializerFunc(json.Marshal),
		Deserializer:    reser.DeserializerFunc(json.Unmarshal),
		TagTypeRegistry: ttr,
	}
}

func Test_ETPolySerializer_CanSerializeDeserialize(t *testing.T) {
	v := pt1{
		ValOne: 42,
	}
	s := makeETPolySerializer(nil)
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
	nv := res.(pt1)
	if nv != v {
		t.Error("Value mismatch")
		return
	}
}

func Test_ETPolySerializer_CanPointerSerializeDeserialize(t *testing.T) {
	v := &pt1{
		ValOne: 42,
	}
	s := makeETPolySerializer(makePointerTTR())
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
	if *nv != *v {
		t.Error("Value mismatch")
		return
	}
}
