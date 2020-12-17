package main

import (
	"fmt"
	"os"
)

func main() {
	var e int
	var err error
	fmt.Print("Что вы хотите использовать?:\n 1. Калькулятор \n 2. Поиск простого числа\n 3.&4. FizzBuzz\n 5.Сортировка bubble\n 6.Сортировка insertion\n")
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
	case 3:
		startfizzbuzz()
	case 4:
		startfizzbuzz2()
	case 5:
		dosortbubble()
	case 6:
		doinsertionsort()
	}
}
