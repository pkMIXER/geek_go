package main

import (
	"fmt"
	"math/rand"
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
		if b == 0 {
			fmt.Printf("На ноль делить в этом калькуляторе нельзя.")
			break
		}
		res = a / b
	case "%":
		fmt.Printf("Остаток от деления: %d", c%d)
		break
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	if !(op == "%") && !(b == 0) {
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
		for f = 2; f <= i; f++ {
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

func startfizzbuzz() {
	var i int
	for i = 0; i <= 100; i++ {
		fmt.Printf("Текущее число %s ", strconv.Itoa(i))
		if i%15 == 0 {
			fmt.Printf("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Printf("Fizz")
		} else if i%5 == 0 {
			fmt.Printf("Buzz")
		}
		fmt.Printf("\n")
	}
}

func startfizzbuzz2() {
	var i int
	var str string
	for i = 0; i <= 100; i++ {
		str = "Текущее число " + strconv.Itoa(i)
		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		fmt.Printf(str + "\n")

	}
}
func dosortbubble() {
	var count, i, j, k int
	var arr, arr2 []int

	fmt.Printf("Введите количество рандомных чисел:")
	fmt.Scanln(&count)
	for i = 0; i < count; i++ {
		min := 1
		max := 500
		arr = append(arr, rand.Intn(max-min+1)+min)

	}
	arr2 = make([]int, count)
	copy(arr2, arr)
	//fmt.Printf("Было: %v \n\n\n", arr2)
	//fmt.Printf("Было: %v \n\n\n", arr)
	for k = 0; k < count-1; k++ {
		for j = 0; j < count-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr[j], arr[j+1], j, arr)
			}
		}
	}

	/*for j=1, ele := range arr {
	if ele > arr[j]{
		swap
	}*/
	//fmt.Print(strconv.Itoa(ele) + " ")
	//}
	fmt.Printf("Было: %v \n\n\n", arr2)
	fmt.Printf("Стало: %v \n", arr)
}

//Меняет местами 2 смежных элемента массива
func swap(first int, second int, pos int, array []int) {
	array[pos] = second
	array[pos+1] = first

}
func doinsertionsort() {
	var count int
	var arr, arr2 []int
	fmt.Printf("Введите количество рандомных чисел:")
	fmt.Scanln(&count)
	for i := 0; i < count; i++ {
		min := 1
		max := 500
		arr = append(arr, rand.Intn(max-min+1)+min)
	}
	arr2 = make([]int, count)
	copy(arr2, arr)
	var n = len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}
	}
	fmt.Printf("Было: %v \n\n\n", arr2)
	fmt.Printf("Стало: %v \n", arr)
}
