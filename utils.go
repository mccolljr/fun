package fun

import "reflect"

func each(slice reflect.Value, fn func(int, reflect.Value) bool) {
	for i, size := 0, slice.Len(); i < size; i++ {
		if fn(i, slice.Index(i)) {
			break
		}
	}
}

func eachKV(m reflect.Value, fn func(k, v reflect.Value) bool) {
	for iter := m.MapRange(); iter.Next(); {
		if fn(iter.Key(), iter.Value()) {
			break
		}
	}
}

func callRet0(fn reflect.Value, args ...reflect.Value) {
	fn.Call(args)
}

func callRet1(fn reflect.Value, args ...reflect.Value) reflect.Value {
	return fn.Call(args)[0]
}

func callRet2(fn reflect.Value, args ...reflect.Value) (a, b reflect.Value) {
	outs := fn.Call(args)
	return outs[0], outs[1]
}
