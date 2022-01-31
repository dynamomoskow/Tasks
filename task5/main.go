package main

import "fmt"

func main() {
	var bilet int
	fmt.Scanln(&bilet)

	var c1, c2, c3, c4, c5, c6 int //цифры билета

	c1 = bilet / 100000
	c2 = bilet % 100000 / 10000
	c3 = bilet % 10000 / 1000
	c4 = bilet % 1000 / 100
	c5 = bilet % 100 / 10
	c6 = bilet % 10
	summaFirst3 := c1 + c2 + c3
	summaLast3 := c4 + c5 + c6
	if summaFirst3 == summaLast3 {
		fmt.Println("YES")
	}else {
		fmt.Println("NO")
	}
}
