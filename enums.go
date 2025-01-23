package enums

import (
	"reflect"
	"strings"
)

type innerElement interface {
	Name() string
	Ordinal() int
	initialize(name string, ordinal int)
	IsValid() bool
}

type Element struct {
	name        string
	ordinal     int
	initialized bool
}

func (e *Element) IsValid() bool {
	return e.initialized
}

func (e *Element) initialize(name string, ordinal int) {
	e.name = name
	e.ordinal = ordinal
	e.initialized = true
}

func (e *Element) Name() string {
	return e.name
}
func (e *Element) Ordinal() int {
	return e.ordinal
}

func (e *Element) String() string {
	return e.Name()
}

func (e *Element) MarshalJSON() ([]byte, error) {
	return []byte(`"` + e.Name() + `"`), nil
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
	v.initialize(name, len(e.values))
	e.values = append(e.values, v)
}

func (e *Enum[T]) String() string {
	s := strings.Builder{}
	s.WriteString("Enum[")
	s.WriteString(reflect.TypeFor[T]().String())
	s.WriteString("]{")
	for i, v := range e.values {
		s.WriteString(v.Name())
		if i != len(e.values)-1 {
			s.WriteString(", ")
		}
	}
	s.WriteString("}")
	return s.String()
}

func Of[T innerElement, E innerEnum[T]](v E) E {
	rv := reflect.ValueOf(v)

	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	targetType := reflect.TypeFor[T]()

	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		fieldT := rt.Field(i)
		if fieldT.Type == targetType {
			v.add(fieldT.Name, rv.Field(i).Interface().(T))
		}
	}

	return v
}
