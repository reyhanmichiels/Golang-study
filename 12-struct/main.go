package main

import "fmt"

//declared struct
type faculty struct {
	name string
}

//declared embedeed strutct
type university struct {
	name string
	faculty
}

func main() {
	//implement struct
	ub := university{
		name: "Brawijaya University",
		faculty: faculty{
			name: "Computer Science",
		},
	}

	//anonymous struct
	var major = struct {
		name string
		university
		faculty
	}{
		name:       "Information Technology",
		faculty:    ub.faculty,
		university: ub,
	}

	fmt.Println(major)
	ub.getUniversityInfo()

	fmt.Println("Name before change\t =", ub.name)
	ub.setName("Brawijaya's University")
	fmt.Println("Name after change\t =", ub.name)
}

//method
func (u university) getUniversityInfo() {
	fmt.Println("University name\t : " + u.name)
	fmt.Println("Faculty name\t : " + u.faculty.name)
}

//pointer method (will reference to struct property)
func (u *university) setName(name string) {
	u.name = name
}