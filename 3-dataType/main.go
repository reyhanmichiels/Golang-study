package main

import "fmt"

func main() {

	//NUMERIC DATA TYPE
	//non-decimal : integer and unsigned integer

	//8 bit integer and unsigned integer
	var integer8 int8 = 1
	var uinteger8 uint8 = 1
	
	fmt.Printf("%T %[1]d\n", integer8)
	fmt.Printf("%T %[1]d\n", uinteger8)

	//16 bit integer and unsigned integer
	var integer16 int16 = 1
	var uinteger16 uint16 = 1
	
	fmt.Printf("%T %[1]d\n", integer16)
	fmt.Printf("%T %[1]d\n", uinteger16)

	//32 bit integer and unsigned integer
	var integer32 int32 = 1
	var uinteger32 uint32 = 1
	
	fmt.Printf("%T %[1]d\n", integer32)
	fmt.Printf("%T %[1]d\n", uinteger32)

	//64 bit integer and unsigned integer
	var integer64 int64 = 1
	var uinteger64 uint64 = 1
	
	fmt.Printf("%T %[1]d\n", integer64)
	fmt.Printf("%T %[1]d\n", uinteger64)

	//32 or 64 bit integer and unsigned integer
	var integer int = 1
	var uinteger int = 1
	fmt.Printf("%T %[1]d\n", integer)
	fmt.Printf("%T %[1]d\n", uinteger)

	//rune, same with 32 bit integer, usesd for represent character 
	var rune rune = 'A'
	fmt.Printf("%T %[1]d %[1]c\n", rune)

	//byte, same with 8 bit uint, used for represent character 
	var byte byte = 'A'
	fmt.Printf("%T %[1]d %[1]c\n", byte)

	//decimal data type
	//32 bit float 
	var float32 float32 = 2.5
	fmt.Printf("%T %.2f\n", float32, float32)

	//64 bit float
	var float64 float64 = 10.5
	fmt.Printf("%T %.2f\n", float64, float64)
	
	//BOOL
	//value just true and false
	var college bool = true
	var work bool = false
	fmt.Printf("%T %[1]t\n", college)
	fmt.Printf("%T %[1]t\n", work)

	//STRING
	var testString string = "Hello World!"
	fmt.Printf("%T %[1]s\n", testString)
}
