package enums

import (
	"reflect"
)

type innerElement interface {
	Name() string
	Ordinal() int
	setName(name string)
	setOrdinal(ordinal int)
}

type Element struct {
	name    string
	ordinal int
}

func (e *Element) setOrdinal(ordinal int) {
	e.ordinal = ordinal
}

func (e *Element) setName(name string) {
	e.name = name
}

func (e *Element) Name() string {
	return e.name
}
func (e *Element) Ordinal() int {
	return e.ordinal
}

type EnumHolder[T any] interface {
	Values() []T
	Names() []string
	ValueOf(name string) *T
}

type innerEnum[T any] interface {
	EnumHolder[T]

	// Add
	//
	// Internal:
	add(name string, v T)
}

type Enum[T innerElement] struct {
	values []T
	ofName map[string]T
}

func (e *Enum[T]) ValueOf(name string) *T {
	if e == nil {
		return nil
	}
	v, ok := e.ofName[name]
	if !ok {
		return nil
	}
	return &v
}

func (e *Enum[T]) Values() []T {
	if e == nil {
		return []T{}
	}
	return e.values
}

func (e *Enum[T]) Names() []string {
	if e == nil {
		return []string{}
	}
	names := make([]string, 0, len(e.ofName))
	for name := range e.ofName {
		names = append(names, name)
	}
	return names
}

func (e *Enum[T]) add(name string, v T) {
	if e == nil {
		return
	}
	if e.ofName == nil {
		e.ofName = make(map[string]T)
	}
	e.ofName[name] = v
	v.setName(name)
	v.setOrdinal(len(e.values))
	e.values = append(e.values, v)
}

func Of[T innerElement, E innerEnum[T]](v E) E {
	rv := reflect.ValueOf(v)

	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	rt := rv.Type()

	targetType := reflect.TypeOf(new(T)).Elem()

	for i := 0; i < rt.NumField(); i++ {
		fieldT := rt.Field(i)

		if fieldT.Type == targetType {
			v.add(fieldT.Name, rv.Field(i).Interface().(T))
		}
	}

	return v
}
