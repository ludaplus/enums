package enums

import (
	"reflect"
	"strings"
	"sync"
)

var unmashals = make(map[reflect.Type]func(name string) any)
var mu sync.Mutex

type innerElement interface {
	Name() string
	Ordinal() int
	initialize(name string, ordinal int)
	IsValid() bool
}

type Element[T innerElement] struct {
	name        string
	ordinal     int
	initialized bool
}

func (e *Element[T]) Equals(other T) bool {
	return e != nil && e.Ordinal() == other.Ordinal()
}

func (e *Element[T]) IsValid() bool {
	return e.initialized
}

func (e *Element[T]) initialize(name string, ordinal int) {
	e.name = name
	e.ordinal = ordinal
	e.initialized = true
}

func (e *Element[T]) Name() string {
	return e.name
}
func (e *Element[T]) Ordinal() int {
	return e.ordinal
}

func (e *Element[T]) String() string {
	return e.Name()
}

func (e *Element[T]) MarshalJSON() ([]byte, error) {
	return []byte(`"` + e.Name() + `"`), nil
}

func (e *Element[T]) UnmarshalHelper(name string) *T {
	return unmashals[reflect.TypeFor[T]()](name).(*T)
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

	func() {
		mu.Lock()
		defer mu.Unlock()

		unmashals[targetType] = func(name string) any {
			return v.ValueOf(name)
		}
	}()

	return v
}
