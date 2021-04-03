package reser_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/teawithsand/reser"
)

func makePrefixPolySerializer() *reser.PrefixPolySerializer {
	ttr := makeTTR()

	s := reser.SerializerFunc(json.Marshal)
	de := reser.DeserializerFunc(json.Unmarshal)
	tts := &reser.STPolySerializer{
		Type:         reflect.TypeOf(""),
		Serializer:   s,
		Deserializer: de,
	}
	ts := &reser.DefaultTagSerializer{
		Serializer:       s,
		Deserializer:     de,
		TagTypeResgistry: ttr,
	}
	return &reser.PrefixPolySerializer{
		TypeTagSerializer:   tts,
		TypeTagDeserializer: tts,
		DataSerializer:      ts,
		DataDeserializer:    ts,
	}
}

func Test_PrefixPolySerializer_Serialize(t *testing.T) {
	v := pt1{
		ValOne: 42,
	}
	s := makePrefixPolySerializer()
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
	actualRes := res.(pt1)
	if actualRes != v {
		t.Error("invalid result value")
	}
}
