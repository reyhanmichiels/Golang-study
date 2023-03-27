package main

import "fmt"

func main() {
	//make map without initial value
	var mapWithOutInitialValue map[string]int
	mapWithOutInitialValue = map[string]int{}

	mapWithOutInitialValue["Reyhan"] = 20

	//make map with initial value
	mapWithInitialValue := make(map[string]int)
	mapWithInitialValue["Reyhan"] = 20

	mapWithInitialValue2 := map[string]int{}
	mapWithInitialValue2["Reyhan"] = 20

	mapWithInitialValue["James"] = 25

	//delete item in map
	delete(mapWithInitialValue, "James")

	//check map item is exist
	_, isExist := mapWithInitialValue["James"]
	fmt.Println(isExist)

	
}
