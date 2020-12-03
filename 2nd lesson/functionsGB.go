package main

import (
	"fmt"
	"os"
	"strconv"
)

func startcalc() {
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
func startsearch() {
	var a, i, f, founded, kolvodeistvii int
	var err error
	kolvodeistvii = 0
	fmt.Print("Добро пожаловать в поиск простых чисел :) \n")
	fmt.Print("Введите целое число-предел числа для поиска \n")
	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Printf("Введено не число")
		os.Exit(0)
	}
	for i = 0; i <= a; i++ {
		//fmt.Printf("Текущее число %s\n", strconv.Itoa(i))
		founded = 0
		for f = 1; f <= i; f++ {
			if i%f == 0 {
				founded++
				kolvodeistvii++
			}
		}
		if founded == 2 {
			fmt.Printf("Текущее число %s\n", strconv.Itoa(i))
		}
	}
	fmt.Printf("Для поиска совершено %s действий(я)\n", strconv.Itoa(kolvodeistvii))
}
