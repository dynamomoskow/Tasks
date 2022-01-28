package main

import "fmt"

func main() {
	var chislo int
	fmt.Scan(&chislo)
	switch {
	case chislo == 10000:
		fmt.Println(chislo /10000)
	case chislo >= 1000:
		fmt.Println(chislo / 1000)
	case chislo >= 100:
		fmt.Println(chislo / 100)
	case chislo >= 10:
		fmt.Println(chislo / 10)
	case chislo < 10:
		fmt.Println(chislo)
	}
}