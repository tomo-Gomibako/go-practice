package main

import (
	"fmt"
	"math"
)

func hello() {
	const (
		sun = iota
		mon
		tue
		wed
		thu
		fri
		sat
	) // iota: sequential
	fmt.Println(sun, mon, tue)
	var target string // target == ""
	target = "World"
	msg := "Hello, " + target + "!"
	fmt.Println(msg)
	number := 3
	var pointer *int
	pointer = &number
	fmt.Println(pointer)
	fmt.Println(*pointer)
	fmt.Println(isPrime(101))
	array := [5]int{0, 1, 1, 3, 4}
	fmt.Println(array[1:3])
	array[1:3][1] = 2
	fmt.Println(array)
	slice := []int{0, 1, 2}
	fmt.Println(slice)
	fmt.Println(cap(slice))
	slice = append(slice, 3, 4)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	s := make([]int, 5, 10)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	copy(s, slice)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	// i := 2
	// for i < 10000000 {
	// 	if isPrime(i) {
	// 		fmt.Print(i, " ")
	// 	}
	// 	i++
	// }
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	i := 2
	for i <= int(math.Sqrt(float64(n))) {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true
}
