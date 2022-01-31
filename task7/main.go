package main

import "fmt"

func main() {
	var a,b int
	workArray := [10]uint8{}
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
		
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&a, &b)
		a,b = b,a
		workArray[a], workArray[b] = workArray[b],workArray[a]
	}
	for i := range workArray {
		fmt.Print(workArray[i], " ")
	}
	// workSlice := workArray[:]
	// workSlice[1], workSlice[2] = workSlice[2], workSlice[1]
	// workSlice[3], workSlice[5] = workSlice[5], workSlice[3]
	// workSlice[7], workSlice[8] = workSlice[8], workSlice[7]
	// for _, k := range workSlice {
	// 	fmt.Println(k, "")
	// }

	// workSlice := workArray[:]
	// workSlice[1], workSlice[2] = workSlice[2], workSlice[1]
	// workSlice[3], workSlice[5] = workSlice[5], workSlice[3]
	// workSlice[7], workSlice[8] = workSlice[8], workSlice[7]
	// fmt.Println(workSlice)
	// for j := 0; j < 3; j++{
	// 	fmt.Scan(&a)
	// 	fmt.Scan(&b)
	// 	workArray[a], workArray[b] = workArray[b], workArray[a]
	// 	fmt.Println(workArray)
	// }
	// for a, v := range workArray{
	// 	fmt.Println(a, v)
	// }

	// fmt.Println(workArray)
	// for j := 0; j < 3; j++ {
	// 	for i, key := range workArray{
		
	// 		fmt.Println(i, key)
	// }
	// }
		// var a, b int
		// for i := 0; i < 3; i ++ {
		// 	fmt.Scan(&a,&b)
		// 	workArray[a], workArray[b] = workArray[b], workArray[a]
		// 	fmt.Println(workArray[i])
		// }
	// fmt.Println(workArray[i])
}