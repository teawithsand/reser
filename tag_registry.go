package reser

import "reflect"

// TagTypeResgistry maps type to type tag and vice versa.
type TagTypeResgistry struct {
	tagType reflect.Type

	registeredTypes        map[TypeTag]reflect.Type
	reverseRegisteredTypes map[reflect.Type]TypeTag
}

// NewTagTypeResgistry creates new tag type registry accepting tags with specified type.
func NewTagTypeResgistry(tagType reflect.Type) *TagTypeResgistry {
	return &TagTypeResgistry{
		tagType:                tagType,
		registeredTypes:        map[TypeTag]reflect.Type{},
		reverseRegisteredTypes: map[reflect.Type]TypeTag{},
	}
}

// GetTypeTagType returns type of type tags used in this registry.
func (ttr *TagTypeResgistry) GetTypeTagType() reflect.Type {
	return ttr.tagType
}

// Registers specified type, so that it can be deserialized using specified type tag.
func (ttr *TagTypeResgistry) RegisterType(ty reflect.Type, et TypeTag) (err error) {
	if reflect.TypeOf(et) != ttr.tagType {
		err = &TypeTagTypeError{
			ExpectedType: ttr.tagType,
			Tag:          et,
		}
	}

	if ty.Kind() == reflect.Ptr {
		err = &TagRegisterError{
			Type:    ty,
			TypeTag: et,
		}
		return
	}

	if ttr.registeredTypes == nil {
		ttr.registeredTypes = map[TypeTag]reflect.Type{}
	}
	if ttr.reverseRegisteredTypes == nil {
		ttr.reverseRegisteredTypes = map[reflect.Type]TypeTag{}
	}

	_, ok := ttr.registeredTypes[et]
	if ok {
		err = &TagRegisterError{
			Type:    ty,
			TypeTag: et,
		}
		return
	}
	_, ok = ttr.reverseRegisteredTypes[ty]
	if ok {
		err = &TagRegisterError{
			Type:    ty,
			TypeTag: et,
		}
		return
	}

	ttr.registeredTypes[et] = ty
	ttr.reverseRegisteredTypes[ty] = et

	return
}

func (ttr *TagTypeResgistry) GetTag(ty reflect.Type) (tt TypeTag, err error) {
	if ttr.registeredTypes == nil {
		err = &TagNotFoundError{
			Type: ty,
		}
		return
	}

	tt, ok := ttr.registeredTypes[ty]
	if !ok {
		// TODO(teaiwthsand): better error
		err = &TagNotFoundError{
			Type: ty,
		}
		return
	}
	return
}

func (ttr *TagTypeResgistry) GetType(tt TypeTag) (ty reflect.Type, err error) {
	if ttr.registeredTypes == nil {
		err = &TypeNotFoundError{
			TypeTag: tt,
		}
		return
	}

	tt, ok := ttr.registeredTypes[ty]
	if !ok {
		err = &TypeNotFoundError{
			TypeTag: tt,
		}
		return
	}
	return
}
