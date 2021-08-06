package reser_test

import (
	"reflect"

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

func makeTTR() *reser.TypeTagRegistry {
	ttr := reser.NewTypeTagRegistry(reflect.TypeOf(""))
	ttr.RegisterType(reflect.TypeOf(pt1{}), "1")
	ttr.RegisterType(reflect.TypeOf(pt2{}), "2")
	ttr.RegisterType(reflect.TypeOf(pt3{}), "3")
	return ttr
}

func makePointerTTR() *reser.TypeTagRegistry {
	ttr := reser.NewTypeTagRegistry(reflect.TypeOf(""))
	ttr.RegisterType(reflect.TypeOf(&pt1{}), "1")
	ttr.RegisterType(reflect.TypeOf(&pt2{}), "2")
	ttr.RegisterType(reflect.TypeOf(&pt3{}), "3")
	return ttr
}
