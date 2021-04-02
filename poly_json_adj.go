package reser

type adjacentTagSerialize struct {
	Type interface{}
	Data interface{}
}
type adjacentTagDeserialize struct {
	Data interface{}
}

// struct used for applying actual external tagging.
type adjacentTag struct {
	Type interface{}
}

func (ad *adjacentTag) GetTypeTag() TypeTag {
	return ad.Type
}

type adjacentTagInt struct {
	Type int
}

func (ad *adjacentTagInt) GetTypeTag() TypeTag {
	return ad.Type
}

type adjacentTagInt8 struct {
	Type int8
}

func (ad *adjacentTagInt8) GetTypeTag() TypeTag {
	return ad.Type
}

type adjacentTagInt16 struct {
	Type int16
}

func (ad *adjacentTagInt16) GetTypeTag() TypeTag {
	return ad.Type
}

type adjacentTagInt32 struct {
	Type int32
}

func (ad *adjacentTagInt32) GetTypeTag() TypeTag {
	return ad.Type
}

type adjacentTagInt64 struct {
	Type int64
}

func (ad *adjacentTagInt64) GetTypeTag() TypeTag {
	return ad.Type
}

type adjacentTagUInt struct {
	Type uint
}

func (ad *adjacentTagUInt) GetTypeTag() TypeTag {
	return ad.Type
}

type adjacentTagUInt8 struct {
	Type uint8
}

func (ad *adjacentTagUInt8) GetTypeTag() TypeTag {
	return ad.Type
}

type adjacentTagUInt16 struct {
	Type uint16
}

func (ad *adjacentTagUInt16) GetTypeTag() TypeTag {
	return ad.Type
}

type adjacentTagUInt32 struct {
	Type uint32
}

func (ad *adjacentTagUInt32) GetTypeTag() TypeTag {
	return ad.Type
}

type adjacentTagUInt64 struct {
	Type uint64
}

func (ad *adjacentTagUInt64) GetTypeTag() TypeTag {
	return ad.Type
}

type adjacentTagString struct {
	Type string
}

func (ad *adjacentTagString) GetTypeTag() TypeTag {
	return ad.Type
}

// TODO(teaiwthsand): bytearray tag support

type tagContainer interface {
	GetTypeTag() TypeTag
}
