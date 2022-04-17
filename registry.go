package reser

import "reflect"

type TypeTag interface {
	comparable
}

func canonizeTypeForRegistry(ty reflect.Type) reflect.Type {
	for ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	return ty
}

type TypeTagRegistry[T TypeTag] interface {
	GetTypeForTag(tag T) (ty reflect.Type)
	GetTagForType(ty reflect.Type) (tag T, ok bool)
}

// DefaultTypeTagRegistry maps type to type tag and vice versa.
// Mutating it is not goroutine safe.
type DefaultTypeTagRegistry[T TypeTag] struct {
	tagToType map[T]reflect.Type
	typeToTag map[reflect.Type]T
}

func (ttr *DefaultTypeTagRegistry[T]) init() {
	if ttr.tagToType == nil {
		ttr.tagToType = map[T]reflect.Type{}
	}
	if ttr.typeToTag == nil {
		ttr.typeToTag = map[reflect.Type]T{}
	}
}

func (ttr *DefaultTypeTagRegistry[T]) AddMapping(tag T, ty reflect.Type) {
	ttr.init()

	ty = canonizeTypeForRegistry(ty)

	ttr.tagToType[tag] = ty
	ttr.typeToTag[ty] = tag
}

func (ttr *DefaultTypeTagRegistry[T]) AddMappingSimple(tag T, instance any) {
	ttr.init()

	ty := reflect.TypeOf(instance)
	ty = canonizeTypeForRegistry(ty)

	ttr.tagToType[tag] = ty
	ttr.typeToTag[ty] = tag
}

func (ttr *DefaultTypeTagRegistry[T]) GetTypeForTag(tag T) (ty reflect.Type) {
	if ttr.tagToType == nil {
		return
	}
	ty = ttr.tagToType[tag]
	return
}

func (ttr *DefaultTypeTagRegistry[T]) GetTagForType(ty reflect.Type) (tag T, ok bool) {
	if ttr.typeToTag == nil {
		return
	}

	ty = canonizeTypeForRegistry(ty)

	tag, ok = ttr.typeToTag[ty]
	return
}
