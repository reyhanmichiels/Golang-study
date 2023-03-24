package main

import "fmt"

func main() {
	//basic for loop
	count := 5
	for i := 1; i <= count; i++ {
		fmt.Println("angka", i)
	}

	//for loop only condition
	count = 6
	for count <= 10 {
		fmt.Println("angka", count)
		count++
	}

	//for loop without argument
	count = 11
	for {
		fmt.Println("angka", count)
		if count == 20 {
			break
		}
		count++
	}
}
