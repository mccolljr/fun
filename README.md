[![Documentation](https://godoc.org/github.com/mccolljr/fun?status.svg)](http://godoc.org/github.com/mccolljr/fun)

Fun
-----
Chainable functional wrappers for iteration of slices and maps in go.

Installation
-----
```
go get -u github.com/mccolljr/fun
```

Usage
-----
```go
slice := []int{1,2,3,4}
fun.Each(slice, func(index int, val int) {
    fmt.Printf("slice[%d]=%d\n", index, val)
})

slice2 := fun.Map(slice, func(index int, val int) int {
    return val*2
}).([]int) //[]int{2,4,6,8}

slice3 := fun.Filter(slice2, func(index int, val int) bool {
    return val > 3
}) // []int{4,6,8}
```