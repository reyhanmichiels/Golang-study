package main

import "fmt"

func main() {
	a := 5

	//declaring pointer
	var pointerA *int
	pointerA = &a

	fmt.Println("Variable a value =", a)
	fmt.Println("Variable a address =", &a)
	fmt.Println("Variable pointerA value =", *pointerA)
	fmt.Println("Variable pointerA address =", pointerA)

	*pointerA = 10
	fmt.Println("\n after value change\n")
	
	fmt.Println("Variable a value =", a)
	fmt.Println("Variable a address =", &a)
	fmt.Println("Variable pointerA value =", *pointerA)
	fmt.Println("Variable pointerA address =", pointerA)
}