package main

import "fmt"

func main() {
	//if-else-if
	reyhanScore := 80

	if reyhanScore >= 80 {
		fmt.Println("A")
	} else if reyhanScore >= 60 {
		fmt.Println("B")
	} else {
		fmt.Println("not pass")
	}

	//switch-case
	color := 1
	switch color {
	case 1:
		fmt.Println("Red")

	case 2:
		fmt.Println("Green")

	case 3, 4, 5:
		fmt.Println("Blue")

	default:
		fmt.Println("Black")
	}


	reyhanScore = 75
	//switch-case if else style
	switch {
	case reyhanScore >= 80:
		fmt.Println("Awesome")
	case reyhanScore >= 60:
		fmt.Println("enough")
		//next case will be executed
		fallthrough
	case reyhanScore < 60:
		fmt.Println("Learn again")
	}

}
