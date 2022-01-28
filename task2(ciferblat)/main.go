package main

import "fmt"

func main() {
	var d int
	fmt.Scanln(&d) // считываем заданные градусы
	//1 градус на цыферблате равен 2 минуты,переведем заданные градусы в минуты, для определния часов поделим на 60
	//для определения минут возьмем остаток от деления на 60
	hours := (2 * d) / 60 
	minutes := (2 * d) % 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}