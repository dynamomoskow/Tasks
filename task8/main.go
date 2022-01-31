package main

import "fmt"

func main() {
	var N int
	// var a int
	// s := make([]int, N)
	fmt.Scan(&N)
	s := make([]int, N)
	for i :=0; i < N; i++ {
		fmt.Scan(&s[i])
	}
	fmt.Println(s[3])
}