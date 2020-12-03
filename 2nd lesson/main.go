package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, res float32
	var op string

	fmt.Print("Добро пожаловать в калькулятор :) \n")
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Println(os.Stderr)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %f\n", res)
}
