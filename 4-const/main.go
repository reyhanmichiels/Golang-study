package main

import "fmt"

//constant declaration
const firstName string = "reyhan"
const lastName = "rusyard"

//multi constant declaration
const (
	age     int = 20
	college     = "Brawijaya University"
)

func main() {
	fmt.Print("Hello my name is ", firstName, " ", lastName, ", I'm ", age, " years old, and im study at ", college, "\n")
}
