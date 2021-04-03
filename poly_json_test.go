package reser_test

import (
	"encoding/json"
	"testing"

	"github.com/teawithsand/reser"
)

func makeETPolySerializer() *reser.ETPolySerializer {
	ttr := makeTTR()

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
	s := makeETPolySerializer()
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
