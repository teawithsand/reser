package reser

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// PrefixPolySerializer serializes and stores tag as prefix in result byte slice in order
// to mix tag along with data stored with it.
//
// It adds minimal overhead to serialized data.
// When tag is shorter than 127 bytes then only one additional byte(apart from tag) is added to serialized data.
type PrefixPolySerializer struct {
	TypeTagSerializer   PolySerializer
	TypeTagDeserializer PolyDeserializer

	DataSerializer   TagSerializer
	DataDeserializer TagDeserializer
}

func (pps *PrefixPolySerializer) PolySerialize(data interface{}) (res []byte, err error) {
	tt, err := pps.DataSerializer.GetTypeTag(data)
	if err != nil {
		return
	}
	serializedTT, err := pps.TypeTagSerializer.PolySerialize(tt)
	if err != nil {
		return
	}
	serializedData, err := pps.DataSerializer.Serialize(data)
	if err != nil {
		return
	}
	sizeTag := [9]byte{}
	sz := binary.PutUvarint(sizeTag[:], uint64(len(serializedTT)))
	res = append(res, sizeTag[:sz]...)
	res = append(res, serializedTT...)
	res = append(res, serializedData...)
	return
}

func (pps *PrefixPolySerializer) PolyDeserialize(data []byte) (res interface{}, err error) {
	const MaxInt = int(^uint(0) >> 1)

	rd := bytes.NewReader(data)
	sz, err := binary.ReadUvarint(rd)
	if err != nil {
		return
	}
	sizeSize := len(data) - rd.Len()
	if sz > uint64(len(data)) || sz > uint64(MaxInt) || sz+uint64(sizeSize) > uint64(len(data)) || sz+uint64(sizeSize) > uint64(MaxInt) {
		err = errors.New("invalid tag size or out of bounds or data corrupted") // TODO(teaiwthsand): better error here
		return
	}
	tagData := data[sizeSize : sizeSize+int(sz)]
	serializedData := data[sizeSize+int(sz):]

	tag, err := pps.TypeTagDeserializer.PolyDeserialize(tagData)
	if err != nil {
		return
	}
	res, err = pps.DataDeserializer.Deserialize(serializedData, tag)
	if err != nil {
		return
	}
	return
}
