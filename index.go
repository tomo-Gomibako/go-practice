package main

import (
	"fmt"
	"practice/fibonacci"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 40; i++ {
		fmt.Println(i, fibonacci.Fib(i))
	}
	fmt.Println((time.Now().UnixNano() - start.UnixNano()), "ns")
	start = time.Now()
	for i := 0; i < 40; i++ {
		fmt.Println(i, fibonacci.Memo(i))
	}
	fmt.Println((time.Now().UnixNano() - start.UnixNano()), "ns")
}
