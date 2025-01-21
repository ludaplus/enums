package enums

import (
	"reflect"
	"strings"
)

type Element struct {
}

type enum interface {
	Values() []any
	Add(v any)
}

type Enum[T any] struct {
	values []*T
}

func (e *Enum[T]) Values() []*T {
	if e == nil {
		return []*T{}
	}
	return e.values
}

// Add
//
// Internal:
func (e *Enum[T]) Add(v *T) {
	if e == nil {
		return
	}
	e.values = append(e.values, v)
}

func unPtrValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func Of[T any](v T) T {
	rv := reflect.ValueOf(v)
	rv = unPtrValue(rv)

	rt := rv.Type()

	var enumField reflect.Value
	for i := 0; i < rt.NumField(); i++ {
		fieldT := rt.Field(i)
		fieldV := rv.Field(i)

		if fieldT.Anonymous && strings.HasPrefix(fieldT.Type.String(), "*enums.Enum[") {
			if fieldV.IsNil() {
				newEnum := reflect.New(fieldT.Type.Elem())
				fieldV.Set(newEnum)
			}
			enumField = fieldV
			break
		}
	}

	if !enumField.IsValid() {
		panic("Enum field not found")
	}

	addMethod, ok := enumField.Type().MethodByName("Add")
	if !ok {
		panic("Add method not found")
	}

	for i := 0; i < rt.NumField(); i++ {
		fieldT := rt.Field(i)
		fieldV := rv.Field(i)

		if fieldT.Type == enumField.Type().Elem().Field(0).Type.Elem() {
			addMethod.Func.Call([]reflect.Value{enumField, fieldV})
		}
	}

	return v
}
