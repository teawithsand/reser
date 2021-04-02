package reser_test

import (
	"reflect"
	"testing"

	"github.com/teawithsand/reser"
)

func Test_TypeRegistry_CanQueryForTags(t *testing.T) {
	type t1 struct{}
	type t2 struct{}
	type t3 struct{}
	ttr := reser.NewTagTypeResgistry(reflect.TypeOf(""))
	ttr.RegisterType(reflect.TypeOf(t1{}), "1")
	ttr.RegisterType(reflect.TypeOf(t2{}), "2")
	ttr.RegisterType(reflect.TypeOf(t3{}), "3")

	tt, err := ttr.GetTag(reflect.TypeOf(t1{}))
	if err != nil {
		t.Error(err)
		return
	}

	if tt.(string) != "1" {
		t.Error("invalid type tag", tt)
		return
	}
}

func Test_TypeRegistry_CanQueryTypes(t *testing.T) {
	type t1 struct{}
	type t2 struct{}
	type t3 struct{}
	ttr := reser.NewTagTypeResgistry(reflect.TypeOf(""))
	ttr.RegisterType(reflect.TypeOf(t1{}), "1")
	ttr.RegisterType(reflect.TypeOf(t2{}), "2")
	ttr.RegisterType(reflect.TypeOf(t3{}), "3")

	ty, err := ttr.GetType("1")
	if err != nil {
		t.Error(err)
		return
	}
	if ty != reflect.TypeOf(t1{}) {
		t.Error("invalid type returned", ty)
		return
	}
}
