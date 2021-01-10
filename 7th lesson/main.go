package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/pkMIXER/geek_go/7th_lesson/confread"
	"github.com/pkMIXER/geek_go/7th_lesson/functionsGB"
)

var (
	prognum = flag.Int("prognum", 0, "Какую программу хотим использовать 1-7")
)

func main() {
	flag.Parse()
	var e int
	var err error

	fmt.Print("Запущено на: " + runtime.GOOS + "\n")
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
		functionsGB.Startcalc()
	case 2:
		functionsGB.Startsearch()
	case 3:
		functionsGB.Startfizzbuzz()
	case 4:
		functionsGB.Startfizzbuzz2()
	case 5:
		functionsGB.Dosortbubble()
	case 6:
		functionsGB.Doinsertionsort()
	case 7:
		confread.ReadConf()
	}

}
