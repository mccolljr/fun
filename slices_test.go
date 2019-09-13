package fun

import (
	"reflect"
	"strings"
	"testing"
	"unicode"
)

func TestSliceEach(t *testing.T) {
	type testCase struct{ slice, predicate interface{} }

	for i, c := range []testCase{
		{slice: []int{1, 2, 3}, predicate: func(_ int, _ int) {}},
		{slice: []string{"a", "b", "c"}, predicate: func(_ int, _ string) {}},
		{slice: []interface{}{"a", 1, false}, predicate: func(_ int, _ interface{}) {}},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("case %d: %v", i, err)
				}
			}()
			Each(c.slice, c.predicate)
		}()
	}
}

func TestSliceMap(t *testing.T) {
	type testCase struct{ slice, want, predicate interface{} }

	for i, c := range []testCase{
		{slice: []int{1, 2, 3}, want: []int{2, 4, 6}, predicate: func(_ int, val int) int { return val * 2 }},
		{slice: []string{"a", "b", "c"}, want: []string{"A", "B", "C"}, predicate: func(_ int, val string) string { return strings.ToUpper(val) }},
		{slice: []interface{}{"a", 1, false}, want: []interface{}{"a", nil, false}, predicate: func(i int, val interface{}) interface{} {
			if i%2 != 0 {
				return nil
			}
			return val
		}},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("case %d: %v", i, err)
				}
			}()
			if got := Map(c.slice, c.predicate); !reflect.DeepEqual(got, c.want) {
				t.Errorf("case %d: expected %v, got %v", i, c.want, got)
			}
		}()
	}
}

func TestSliceFlatMap(t *testing.T) {
	type testCase struct{ slice, want, predicate interface{} }

	for i, c := range []testCase{
		{
			slice:     []string{"a,b", "c,d", "e,f"},
			want:      []string{"a", "b", "c", "d", "e", "f"},
			predicate: func(_ int, val string) []string { return strings.Split(val, ",") }},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("case %d: %v", i, err)
				}
			}()
			if got := FlatMap(c.slice, c.predicate); !reflect.DeepEqual(got, c.want) {
				t.Errorf("case %d: expected %v, got %v", i, c.want, got)
			}
		}()
	}
}

func TestSliceFilter(t *testing.T) {
	type testCase struct{ slice, want, predicate interface{} }

	for i, c := range []testCase{
		{slice: []int{1, 2, 3}, want: []int{1, 3}, predicate: func(_ int, val int) bool { return val%2 != 0 }},
		{slice: []string{"a", "b", "c"}, want: []string{"a", "c"}, predicate: func(i int, _ string) bool { return i%2 == 0 }},
		{slice: []interface{}{"a", 1, false}, want: []interface{}{"a"}, predicate: func(_ int, val interface{}) bool {
			_, ok := val.(string)
			return ok
		}},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("case %d: %v", i, err)
				}
			}()
			if got := Filter(c.slice, c.predicate); !reflect.DeepEqual(got, c.want) {
				t.Errorf("case %d: expected %v, got %v", i, c.want, got)
			}
		}()
	}
}

func TestSliceCollect(t *testing.T) {
	type testCase struct{ slice, want, agg, predicate interface{} }

	for i, c := range []testCase{
		{
			slice: []int{1, 2, 3},
			want:  []int{1, 3},
			agg:   new([]int),
			predicate: func(agg *[]int, _ int, val int) {
				if val%2 != 0 {
					*agg = append(*agg, val)
				}
			},
		},
		{
			slice: []string{"a", "b", "c"},
			want:  []string{"b"},
			agg:   new([]string),
			predicate: func(agg *[]string, i int, val string) {
				if i%2 != 0 {
					*agg = append(*agg, val)
				}
			},
		},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("case %d: %v", i, err)
				}
			}()
			Collect(c.agg, c.slice, c.predicate)
			got := reflect.ValueOf(c.agg).Elem().Interface()
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("case %d: expected %v, got %v", i, c.want, got)
			}
		}()
	}
}

func TestSliceSome(t *testing.T) {
	type testCase struct {
		slice, predicate interface{}
		wantSome         bool
	}

	for i, c := range []testCase{
		{slice: []int{1, 2, 3}, wantSome: true, predicate: func(_ int, val int) bool { return val%2 != 0 }},
		{slice: []string{"a", "b", "c"}, wantSome: false, predicate: func(_ int, val string) bool { return unicode.IsUpper(rune(val[0])) }},
		{slice: []interface{}{"a", 1, false}, wantSome: false, predicate: func(_ int, val interface{}) bool {
			_, ok := val.(float64)
			return ok
		}},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("case %d: %v", i, err)
				}
			}()
			if got := Some(c.slice, c.predicate); !reflect.DeepEqual(got, c.wantSome) {
				t.Errorf("case %d: expected Some to return %v, got %v", i, c.wantSome, got)
			}
		}()
	}
}

func TestSliceCount(t *testing.T) {
	type testCase struct {
		slice, predicate interface{}
		wantCount        int
	}

	for i, c := range []testCase{
		{slice: []int{1, 2, 3}, wantCount: 1, predicate: func(_ int, val int) bool { return val%2 == 0 }},
		{slice: []string{"a", "b", "c"}, wantCount: 0, predicate: func(_ int, val string) bool { return unicode.IsUpper(rune(val[0])) }},
		{slice: []interface{}{"a", 1, false}, wantCount: 3, predicate: func(_ int, _ interface{}) bool { return true }},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("case %d: %v", i, err)
				}
			}()
			if got := Count(c.slice, c.predicate); !reflect.DeepEqual(got, c.wantCount) {
				t.Errorf("case %d: expected Count to return %v, got %v", i, c.wantCount, got)
			}
		}()
	}
}

func TestSliceUnique(t *testing.T) {
	type testCase struct{ slice, want interface{} }

	for i, c := range []testCase{
		{slice: []int{1, 1, 2, 2, 3, 3}, want: []int{1, 2, 3}},
		{slice: []string{"a", "b", "b", "b", "c"}, want: []string{"a", "b", "c"}},
		{slice: []interface{}{"a", "a", 1, 1, 1, true, false}, want: []interface{}{"a", 1, true, false}},
	} {
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("case %d: %v", i, err)
				}
			}()
			if got := Unique(c.slice); !reflect.DeepEqual(got, c.want) {
				t.Errorf("case %d: expected %v, got %v", i, c.want, got)
			}
		}()
	}
}
