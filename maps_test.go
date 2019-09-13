package fun

import (
	"reflect"
	"testing"
)

func TestMapEach(t *testing.T) {
	type testCase struct{ m, predicate interface{} }

	for i, c := range []testCase{
		{m: map[string]int{"a": 1, "b": 2, "c": 3}, predicate: func(_ string, _ int) {}},
		{m: map[string]string{"a": "A", "b": "B", "c": "C"}, predicate: func(_, _ string) {}},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("case %d: %v", i, err)
				}
			}()
			EachKV(c.m, c.predicate)
		}()
	}
}

// TODO FIXME this test will fail if the ordering on the map iteration changes...
func TestMapKeys(t *testing.T) {
	type testCase struct{ m, wantKeys interface{} }

	for i, c := range []testCase{
		{m: map[string]int{"a": 1, "b": 2, "c": 3}, wantKeys: []string{"a", "b", "c"}},
		{m: map[int]string{1: "a", 2: "b", 3: "c"}, wantKeys: []int{1, 2, 3}},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("case %d: %v", i, err)
				}
			}()
			if got := Keys(c.m); !reflect.DeepEqual(got, c.wantKeys) {
				t.Errorf("case %d: expected %v, got %v", i, c.wantKeys, got)
			}
		}()
	}
}

// TODO FIXME this test will fail if the ordering on the map iteration changes...
func TestMapValues(t *testing.T) {
	type testCase struct{ m, wantValues interface{} }

	for i, c := range []testCase{
		{m: map[string]int{"a": 1, "b": 2, "c": 3}, wantValues: []int{1, 2, 3}},
		{m: map[int]string{1: "a", 2: "b", 3: "c"}, wantValues: []string{"a", "b", "c"}},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("case %d: %v", i, err)
				}
			}()
			if got := Values(c.m); !reflect.DeepEqual(got, c.wantValues) {
				t.Errorf("case %d: expected %v, got %v", i, c.wantValues, got)
			}
		}()
	}
}
