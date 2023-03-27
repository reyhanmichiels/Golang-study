package main

import "fmt"

func main() {
	//declaring array
	var introduction [5]string

	//initializing array
	introduction[0] = "Hello"
	introduction[1] = "my"
	introduction[2] = "name"
	introduction[3] = "is"
	introduction[4] = "Reyhan"

	fmt.Println(introduction)

	//declaring and initializing array
	var intro = [5]string{
		"Hello",
		"my",
		"name",
		"is",
		"Reyhan",
	}

	//for range in array
	for _, data := range intro {
		fmt.Print(data + " ")
	}
	fmt.Println()
}
