package main

import (
	"fmt"
	"os"
)

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

func fibr(i int, fn map[int]int) int {
	var f int

	if i <= 2 {
		fn[i] = 1
		f = 1
	} else {
		//Проверяем если уже есть рассчитанный элемент при обращении в рекурсию
		//выдаем существующий элемент, а не считаем заново.
		if len(fn) == i {
			//fmt.Printf("\nПошла рекурсия для %d\n", len(fn))
			f = fibr(i-2, fn) + fibr(i-1, fn)
		} else {
			f = fn[i-2] + fn[i-1]
		}
	}
	return f
}

func main() {
	var b uint
	var err error
	fmt.Print("Введите количество для поиска числа фибоначи: ")
	_, err = fmt.Scanln(&b)
	if (err != nil) || (b == 0) {
		fmt.Printf("Введено неверное количество")
		os.Exit(0)
	}
	a := int(b)
	for i := 0; i < a; i++ {
		fmt.Printf("%d, ", fib(i))
	}
	//Поиск фибо по рекурссии
	fn := make(map[int]int)
	fmt.Printf("\nФибо по рекурсии:\n")
	for j := 0; j < a; j++ {
		f := fibr(j, fn)
		fn[j] = f
		fmt.Printf("%d, ", f)

	}

}
