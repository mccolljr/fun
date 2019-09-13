package fun

import "reflect"

func EachKV(m, fn interface{}) {
	mv, fv := reflect.ValueOf(m), reflect.ValueOf(fn)
	eachKV(mv, func(k, v reflect.Value) bool {
		callRet0(fv, k, v)
		return false
	})
}

func MapKV(m, fn interface{}) interface{} {
	mv, fv := reflect.ValueOf(m), reflect.ValueOf(fn)
	result := reflect.MakeMap(reflect.MapOf(mv.Type().Key(), fv.Type().Out(0)))
	iter := mv.MapRange()
	for iter.Next() {
		result.SetMapIndex(
			iter.Key(),
			fv.Call([]reflect.Value{
				iter.Key(), iter.Value(),
			})[0],
		)
	}
	return result.Interface()
}

func SomeKV(m, fn interface{}) bool {
	mv, fv := reflect.ValueOf(m), reflect.ValueOf(fn)
	hasSome := false
	eachKV(mv, func(k, v reflect.Value) bool {
		hasSome = callRet1(fv, k, v).Bool()
		return hasSome
	})
	return hasSome
}

func FilterKV(m, fn interface{}) interface{} {
	mv, fv := reflect.ValueOf(m), reflect.ValueOf(fn)
	result := reflect.MakeMap(mv.Type())
	eachKV(mv, func(k, v reflect.Value) bool {
		if callRet1(fv, k, v).Bool() {
			result.SetMapIndex(k, v)
		}
		return false
	})
	return result.Interface()
}

func Keys(m interface{}) interface{} {
	mv := reflect.ValueOf(m)
	out := reflect.MakeSlice(reflect.SliceOf(mv.Type().Key()), 0, mv.Len())
	for _, k := range mv.MapKeys() {
		out = reflect.Append(out, k)
	}
	return out.Interface()
}

func Values(m interface{}) interface{} {
	mv := reflect.ValueOf(m)
	out := reflect.MakeSlice(reflect.SliceOf(mv.Type().Elem()), 0, mv.Len())
	iter := mv.MapRange()
	for iter.Next() {
		out = reflect.Append(out, iter.Value())
	}
	return out.Interface()
}
