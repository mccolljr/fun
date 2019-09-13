package fun

import "reflect"

var (
	reflectTypBoolean = reflect.TypeOf(false)
)

func Each(slice, fn interface{}) {
	sv, fv := reflect.ValueOf(slice), reflect.ValueOf(fn)
	each(sv, func(i int, v reflect.Value) bool {
		callRet0(fv, reflect.ValueOf(i), v)
		return false
	})
}

func Map(slice, fn interface{}) interface{} {
	sv, fv := reflect.ValueOf(slice), reflect.ValueOf(fn)
	length := sv.Len()
	result := reflect.MakeSlice(reflect.SliceOf(fv.Type().Out(0)), length, length)
	each(sv, func(i int, v reflect.Value) bool {
		result.Index(i).Set(callRet1(fv, reflect.ValueOf(i), v))
		return false
	})
	return result.Interface()
}

func FlatMap(slice, fn interface{}) interface{} {
	sv, fv := reflect.ValueOf(slice), reflect.ValueOf(fn)
	retTyp := fv.Type().Out(0)
	retTypEltTyp := retTyp.Elem()
	result := reflect.MakeSlice(reflect.SliceOf(retTypEltTyp), 0, 0)
	each(sv, func(i int, v reflect.Value) bool {
		toFlatten := fv.Call([]reflect.Value{reflect.ValueOf(i), v})[0]
		each(toFlatten, func(_ int, v reflect.Value) bool {
			result = reflect.Append(result, v)
			return false
		})
		return false
	})
	return result.Interface()
}

func Some(slice, fn interface{}) bool {
	sv, fv := reflect.ValueOf(slice), reflect.ValueOf(fn)
	hasSome := false
	each(sv, func(i int, v reflect.Value) bool {
		hasSome = callRet1(fv, reflect.ValueOf(i), v).Bool()
		return hasSome
	})
	return hasSome
}

func Count(slice, fn interface{}) int {
	sv, fv := reflect.ValueOf(slice), reflect.ValueOf(fn)
	count := 0
	each(sv, func(i int, v reflect.Value) bool {
		if callRet1(fv, reflect.ValueOf(i), v).Bool() {
			count++
		}
		return false
	})
	return count
}

func Unique(slice interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	result := reflect.MakeSlice(sv.Type(), 0, sv.Len())
	seen := reflect.MakeMap(reflect.MapOf(sv.Type().Elem(), reflectTypBoolean))
	each(sv, func(i int, v reflect.Value) bool {
		if got := seen.MapIndex(v); got.IsValid() && got.Bool() {
			return false
		}
		result = reflect.Append(result, v)
		seen.SetMapIndex(v, reflect.ValueOf(true))
		return false
	})
	return result.Interface()
}

func Filter(slice, fn interface{}) interface{} {
	sv, fv := reflect.ValueOf(slice), reflect.ValueOf(fn)
	result := reflect.MakeSlice(sv.Type(), 0, sv.Len())
	each(sv, func(i int, v reflect.Value) bool {
		if callRet1(fv, reflect.ValueOf(i), v).Bool() {
			result = reflect.Append(result, v)
		}
		return false
	})
	return result.Interface()
}

func Collect(agg, slice, fn interface{}) {
	sv, fv := reflect.ValueOf(slice), reflect.ValueOf(fn)
	aggv := reflect.ValueOf(agg)
	each(sv, func(i int, v reflect.Value) bool {
		callRet0(fv, aggv, reflect.ValueOf(i), v)
		return false
	})
}
