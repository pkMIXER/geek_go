package main

import (
	"fmt"
	"os"
)

//Fibonacci recursive
func fib(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	var a int
	var err error
	fmt.Print("Введите число для поиска числа фибоначи: ")
	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Printf("Введено не число")
		os.Exit(0)
	}
	for i := 0; i < a; i++ {
		fmt.Printf("%d, ", fib(i))
	}
}
