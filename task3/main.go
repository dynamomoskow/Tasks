package main

import "fmt"

func main() {
	var chislo int
	fmt.Scan(&chislo)
	if chislo == 10000{
		fmt.Println(chislo /10000)
	}else if chislo >= 1000 {
		fmt.Println(chislo / 1000)
	}else if chislo >= 100 {
		fmt.Println(chislo / 100)
	}else if chislo >= 10 {
		fmt.Println(chislo / 10)
	}else if chislo < 10 {
		fmt.Println(chislo)
	}

}