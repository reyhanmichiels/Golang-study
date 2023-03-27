package main

func main() {
	//declare and create slice 
	var testSlice = []int{
		1,
		2,
		3,
	}

	//create slice with slicing existed slice or array
	testSlicing := testSlice[0:]

	//create slice with make function 
	testMake := make([]int, 3)
	testMake[0] = testSlicing[0]
	testMake[1] = testSlicing[1]
	testMake[2] = testSlicing[2]

	//append capacity slice
	testMake = append(testMake, 4)

	testCopy := make([]int, 2)

	//copying slice data
	copy(testCopy, testMake)
}


