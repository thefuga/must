package reflect

import "reflect"

func CastErr(v any) error {
	err, ok := v.(error)
	if !ok {
		return nil
	}
	return err
}

func Set(dst any, src any) {
	ValueOf(dst).Set(ValueOfInterfaceElem(src))
}

type Value struct {
	v reflect.Value
}

func (v Value) CanSet(t Value) bool {
	return v.v.IsValid() && t.v.IsValid() && v.v.CanSet() && t.v.CanConvert(v.v.Type())
}

func (v Value) Set(t Value) {
	if elem := v.Elem(); elem.CanSet(t) {
		elem.v.Set(t.v.Convert(t.v.Type()))
	}
}

func ValueOf(v any) Value {
	return Value{reflect.ValueOf(v)}
}

func ValueOfInterfaceElem(v any) Value {
	return ValueOf(ValueOf(&v).Elem().v.Interface())
}

func (v Value) Elem() Value {
	if !v.v.IsValid() || v.v.Kind() != reflect.Ptr {
		return v
	}

	return Value{v.v.Elem()}
}
