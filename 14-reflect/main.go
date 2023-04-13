package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	Age  int
}

func (p person) info() {
	ref := reflect.ValueOf(p)

	for i := 0; i < ref.NumField(); i++ {
		fmt.Println("Field Name\t ", ref.Type().Field(i).Name)
		fmt.Println("Field Type\t ", ref.Type().Field(i).Type)
		fmt.Println("Field Value\t ", ref.Field(i).Interface())
		fmt.Println()
	}
}

func (p *person) SetName(name string) {
	p.Name = name
}

func main() {
	temp := 10

	ref := reflect.ValueOf(temp)

	//check value type
	if ref.Kind() == reflect.Int {
		fmt.Println("variable type :", ref.Type())
		fmt.Println("variable value :", ref.Int())
		fmt.Printf("%T %[1]v\n", ref.Interface())
	}

	person1 := person{
		Name: "Reyhan",
		Age:  20,
	}

	person1.info()

	reflect.ValueOf(&person1).MethodByName("SetName").Call([]reflect.Value{
		reflect.ValueOf("Reyhan Hafiz Rusyard"),
	})

	person1.info()
}
