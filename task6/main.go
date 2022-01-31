package main

import "fmt"

func main() {
	var chislo float64
	fmt.Scanln(&chislo)
	switch {
	case chislo <= 0:
		fmt.Printf("число %2.2f не подходит", chislo)
	case chislo > 10000:
		fmt.Printf("%e", chislo)
	default:
		fmt.Printf("%0.4f", chislo * chislo)
	}
}
