package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var i int = 10
	fmt.Println(i)

	var f32 float32 = 2.2
	var f64 float64 = 10.345
	fmt.Println("f32 is", f32, ",f64 is", f64)

	var bf bool = false
	var bt bool = true
	fmt.Println("bf is", bf, "bf is ", bt)

	var s1 string = "Hello"
	var s2 string = "世界"

	fmt.Println("s1 is ", s1, "s2 is ", s2)
	fmt.Println("s1+s2=", s1+s2)
	s1 += s2
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)

	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println(zi, zf, zb, zs)

	fmt.Println("-------------")
	test()

	pi := &i
	fmt.Println(*pi)

	const name = "小雪"
	const (
		one = iota + 1
		two
		three
		four
	)

	fmt.Println(name, one, two, three, four)

	i2s := strconv.Itoa(i)
	fmt.Println(i2s)

	fmt.Println("=============" + s1)
	fmt.Println(strings.HasPrefix(s1, "H"))
	fmt.Println(strings.Index(s1, "o"))
	fmt.Println(strings.ToUpper(s1))

}

func test() {
	i := 11
	bf := false
	s1 := "Hello1"
	fmt.Println(i, bf, s1)
}
