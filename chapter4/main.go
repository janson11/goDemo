package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 100; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Println("the sum is", sum)

	array := [5]string{"a", "b", "c", "d", "e"}
	array1 := [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[0] + array1[2])

	for i, v := range array {
		fmt.Printf("数组索引：%d,对应值：%s\n", i, v)
	}

	for _, v := range array {
		fmt.Printf("数组对应值：%s\n", v)
	}

}
