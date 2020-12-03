package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, b, res float32
	var c, d int
	var op string
	var err error

	fmt.Print("Добро пожаловать в калькулятор :) \n")
	fmt.Print("Введите первое число: ")
	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Printf("введено не число")
		os.Exit(0)
	}

	//fmt.Println(os.Stderr)

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Printf("введено не число")
		os.Exit(0)
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /,%-взятие остатка.): ")
	fmt.Scanln(&op)

	if op == "%" {
		c = int(a)
		d = int(b)
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "%":
		fmt.Printf("Остаток от деления: %s", strconv.Itoa(c%d))
		break
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	if !(op == "%") {
		fmt.Printf("Результат выполнения операции: %f\n", res)
	}
}
