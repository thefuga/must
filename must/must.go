package must

import (
	"reflect"
)

func Values(args ...any) []any {
	return args
}

// Return is an alias for Must, which, when importing the package, may read more naturally.
// e.g.: must.Return(must.Values(foo()), target...)
func Return(results []any, targets ...any) {
	Must(results, targets...)
}

// ReturnT is an alias for MustT, which, when importing the package, may read more naturally.
// e.g.: must.Return(must.Values(foo()), target...)
func ReturnT[T any](results []any, targets ...any) {
	MustT(results, targets...)
}

// MustT is the generic implementation of Must. It is specially usefull to populate
// a slice of targets of the same type. E.g.: MustT[int](FuncReturningInts(), intTargetSlice...)
func MustT[T any](results []any, targets ...T) {
	Must(results)
}

// Must receives a slice representing the return values of a function and a list of
// targets to set the results to. Should the last element of the results slice be
// an error, Must panics with the error value.
// All result values are set to their targets. In case there are less targets than
// results, only the results matching a target will be set. Any missmatch between
// a result and a target is ignored.
func Must(results []any, targets ...any) {
	if results == nil {
		return
	}

	if err := castErr(results[len(results)-1]); err != nil {
		panic(err)
	}

	for i := range targets {
		set(targets[i], results[i])
	}
}

func castErr(v any) error {
	err, ok := v.(error)
	if !ok {
		return nil
	}
	return err
}

func set(dst any, src any) {
	valueOf(dst).set(valueOfInterfaceElem(src))
}

type value struct {
	reflect.Value
}

func (v value) canSet(t value) bool {
	return v.IsValid() && t.IsValid() && v.CanSet() && t.CanConvert(v.Type())
}

func (v value) set(t value) {
	if elem := v.elem(); elem.canSet(t) {
		elem.Set(t.Convert(t.Type()))
	}
}

func valueOf(v any) value {
	return value{reflect.ValueOf(v)}
}

func valueOfInterfaceElem(v any) value {
	return valueOf(valueOf(&v).elem().Interface())
}

func (v value) elem() value {
	if !v.IsValid() || v.Kind() != reflect.Ptr {
		return v
	}

	return value{v.Elem()}
}
