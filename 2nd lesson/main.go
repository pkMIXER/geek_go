package main

import (
	"fmt"
	"os"
)

func main() {
	var e int
	var err error
	fmt.Print("Что вы хотите использовать?:\n 1. Калькулятор \n 2. Поиск простого числа\n ")
	_, err = fmt.Scanln(&e)
	if err != nil {
		fmt.Printf("Введено не число")
		os.Exit(0)
	}
	switch e {
	case 1:
		startcalc()
	case 2:
		startsearch()
	}
}
