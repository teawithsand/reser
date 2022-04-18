package reser

// Type, which is both TagWrapper and DataWrapper.
type TagDataWrapper[T TypeTag, D any] interface {
	TagWrapper[T]
	DataWrapper[D]
}

type TagWrapper[T TypeTag] interface {
	GetTag() T
	SetTag(T)
}

type DataWrapper[T any] interface {
	GetData() T
	SetData(T)
}

type WrapperFactory[T TypeTag, D any] interface {
	NewTagWrapper() TagWrapper[T]
	NewDataWrapper() DataWrapper[D]
	NewTagDataWrapper() TagDataWrapper[T, D]
}

// DEFAULT IMPLS HERE //

// Default implementation of TaggedDataWrapper, which uses "tag" and "data" fields.
type DefaultTagDataWrapper[T TypeTag, D any] struct {
	DefaultTagWrapper[T]
	DefaultDataWrapper[D]
}

type DefaultTagWrapper[T TypeTag] struct {
	Tag T `json:"tag" xml:"tag" yaml:"tag" toml:"tag"`
}

func (dtw *DefaultTagWrapper[T]) GetTag() T {
	return dtw.Tag
}
func (dtw *DefaultTagWrapper[T]) SetTag(tag T) {
	dtw.Tag = tag
}

type DefaultDataWrapper[T any] struct {
	Data T `json:"data" xml:"data" yaml:"data" toml:"data"`
}

func (ddw *DefaultDataWrapper[T]) GetData() T {
	return ddw.Data
}

func (ddw *DefaultDataWrapper[T]) SetData(data T) {
	ddw.Data = data
}

type DefaultWrapperFactory[T TypeTag, D any] struct{}

func (DefaultWrapperFactory[T, D]) NewTagWrapper() TagWrapper[T] {
	return &DefaultTagWrapper[T]{}
}
func (DefaultWrapperFactory[T, D]) NewDataWrapper() DataWrapper[D] {
	return &DefaultDataWrapper[D]{}
}
func (DefaultWrapperFactory[T, D]) NewTagDataWrapper() TagDataWrapper[T, D] {
	return &DefaultTagDataWrapper[T, D]{}
}
