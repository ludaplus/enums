package zero

import (
	"reflect"
)

type (
	internalEnum[T any] interface {
		Values() []T
	}
	internalStructElement interface {
		setOrdinal(int)
		setName(string)
	}
	internalElementUnion interface {
		~string | ~int
	}

	Element struct {
		ordinal int
		name    string
	}
	Enum[T any] struct {
	}
)

func (e *Element) setName(name string) {
	e.name = name
}

func (e *Element) setOrdinal(ordinal int) {
	e.ordinal = ordinal
}

func (e *Element) Name() string {
	return e.name
}

func (e *Element) Ordinal() int {
	return e.ordinal
}

func (e Enum[T]) Values() []T {
	return nil
}

func OfStruct[T internalStructElement, E internalEnum[T]](v E) E {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		panic("Of: v must be a struct")
	}

	targetType := reflect.TypeFor[T]()
	var ordinal int
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if field.Type() != targetType {
			continue
		}
		ordinal++
	}
	return v
}

type PostTypeString string

func Of[T internalElementUnion, E internalEnum[T]](v E) E {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		panic("Of: v must be a struct")
	}

	targetType := reflect.TypeFor[T]()
	var ordinal int
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if field.Type() != targetType {
			continue
		}

		if field.IsZero() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(rv.Type().Field(i).Name)
			case reflect.Int:
				field.SetInt(int64(ordinal))
			case reflect.Bool:
			case reflect.Float32:
			case reflect.Float64:
			case reflect.Complex64:
			case reflect.Complex128:
			case reflect.Uint:
			case reflect.Uint8:
			case reflect.Uint16:
			case reflect.Uint32:
			case reflect.Uint64:
			case reflect.Uintptr:
			case reflect.Int8:
			case reflect.Int16:
			case reflect.Int32:
			case reflect.Int64:
			case reflect.Array:
			case reflect.Chan:
			case reflect.Func:
			case reflect.Interface:
			case reflect.Map:
			case reflect.Ptr:
			default:
				panic("unhandled default case")
			}
		}
		ordinal++
	}
	return v
}

var StringPostTypes = Of(&struct {
	Enum[PostTypeString]
	Unknown,
	Post,
	Page,
	Note PostTypeString
}{})

type PostTypeNumber int

var NumberPostTypes = Of(&struct {
	Enum[PostTypeNumber]
	Unknown,
	Post,
	Page,
	Note PostTypeNumber
}{})

type PostTypeStruct struct {
	Element
}

var StructPostTypes = OfStruct(struct {
	Enum[*PostTypeStruct]
	Unknown,
	Post,
	Page,
	Note PostTypeStruct
}{})
