package reser

import "reflect"

// TypeTagRegistry maps type to type tag and vice versa.
type TypeTagRegistry struct {
	tagType reflect.Type

	tagToType map[TypeTag]reflect.Type
	typeToTag map[reflect.Type]TypeTag
}

// NewTypeTagRegistry creates new tag type registry accepting tags with specified type.
func NewTypeTagRegistry(tagType reflect.Type) *TypeTagRegistry {
	return &TypeTagRegistry{
		tagType:   tagType,
		tagToType: map[TypeTag]reflect.Type{},
		typeToTag: map[reflect.Type]TypeTag{},
	}
}

// GetTypeTagType returns type of type tags used in this registry.
func (ttr *TypeTagRegistry) GetTypeTagType() reflect.Type {
	return ttr.tagType
}

// Registers specified type, so that it can be deserialized using specified type tag.
func (ttr *TypeTagRegistry) RegisterType(ty reflect.Type, et TypeTag) (err error) {
	if reflect.TypeOf(et) != ttr.tagType {
		err = &TypeTagTypeError{
			ExpectedType: ttr.tagType,
			Tag:          et,
		}
		return
	}

	_, ok := ttr.tagToType[et]
	if ok {
		err = &TagRegisterError{
			Type:    ty,
			TypeTag: et,
		}
		return
	}
	_, ok = ttr.typeToTag[ty]
	if ok {
		err = &TagRegisterError{
			Type:    ty,
			TypeTag: et,
		}
		return
	}

	ttr.tagToType[et] = ty
	ttr.typeToTag[ty] = et

	return
}

func (ttr *TypeTagRegistry) GetTag(ty reflect.Type) (tt TypeTag, err error) {
	if ttr.tagToType == nil {
		err = &TagNotFoundError{
			Type: ty,
		}
		return
	}

	tt, ok := ttr.typeToTag[ty]
	if !ok {
		// TODO(teaiwthsand): better error
		err = &TagNotFoundError{
			Type: ty,
		}
		return
	}
	return
}

func (ttr *TypeTagRegistry) GetType(tt TypeTag) (ty reflect.Type, err error) {
	if ttr.tagToType == nil {
		err = &TypeNotFoundError{
			TypeTag: tt,
		}
		return
	}

	ty, ok := ttr.tagToType[tt]
	if !ok {
		err = &TypeNotFoundError{
			TypeTag: tt,
		}
		return
	}
	return
}
