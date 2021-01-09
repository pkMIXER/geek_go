package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	prognum = flag.Int("prognum", 0, "Какую программу хотим использовать 1-7")
)

func main() {
	flag.Parse()
	var e int
	var err error
	//Проверяем задан ли флаг программы
	if *prognum != 0 {
		if *prognum < 1 || *prognum > 7 {
			log.Fatalf("Значение prognum должно быть от 1 до 7")
		} else {
			e = *prognum
		}
	} else {
		fmt.Print("Что вы хотите использовать?:\n" +
			"1. Калькулятор \n" +
			"2. Поиск простого числа\n" +
			"3.&4. FizzBuzz\n" +
			"5.Сортировка bubble\n" +
			"6.Сортировка insertion\n" +
			"7.Чтение конфига\n")
		_, err = fmt.Scanln(&e)
		if err != nil {
			fmt.Printf("Введено не число")
			os.Exit(0)
		}
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
	case 7:
		readconf()
	}

}
