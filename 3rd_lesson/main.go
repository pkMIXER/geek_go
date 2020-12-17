package main

import (
	"fmt"
	"os"
)

//Fibonacci recursive
func fibr(n int) int {
	if n < 2 {
		return 1
	}
	return fibr(n-2) + fibr(n-1)
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
		fmt.Printf("%d, ", fibr(i))
	}
}
