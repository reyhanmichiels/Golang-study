package main

import "fmt"

//manifest typing
//deklarasi variabel menggunakan tipe data
var name string = "Reyhan"

//type inference
//tipe data dari variabel ditentukan oleh nilainya
var major = "Technology Information"

func main() {
	//short variable declaration
	//declaration variable without var
	college := "Brawijaya University"

	fmt.Printf("Hello my name is %s, Im study at %s major %s\n", name, college, major)
}