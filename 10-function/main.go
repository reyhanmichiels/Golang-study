package main

import "fmt"

func main() {
	//void function
	greeting("Reyhan")

	//return function
	fmt.Printf("Max number between 10 and 20 is %d\n", max(10, 20))

	//multiple return function
	add, sub, multi, div := basicMathOperation(20, 10)
	fmt.Printf("20 + 10 = %d\n20 - 10 = %d\n20 x 10 = %d\n20 / 10 = %d\n", add, sub, multi, div)

	add2, sub2, multi2, div2 := basicMathOperationWithPredefinedReturnValue(15, 10)
	fmt.Printf("15 + 10 = %d\n15 - 10 = %d\n15 x 10 = %d\n15 / 10 = %d\n", add2, sub2, multi2, div2)

	//variadic function
	temp := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Total =", temp)

	numberToSum := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	temp = sum(numberToSum...)
	fmt.Println("Total =", temp)

	//closure function in variable
	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	fmt.Println("Min value between 82 and 18 is", min(82, 18))

	//iife closure function
	sort := func(nums []int) []int {
		//selection sort with nested for loop
		for index, value := range nums {
			minIndex := index
			minValue := value

			for i := index + 1; i < len(nums); i++ {
				if nums[i] < minValue {
					minValue = nums[i]
					minIndex = i
				}
			}

			nums[index], nums[minIndex] = minValue, value
		}

		//selection sort with nested range loop
		// for index, value := range nums {
		// 	minValue := value
		// 	minIndex := index

		// 	indexInnerLoop := index + 1
		// 	for in, val := range nums[indexInnerLoop:] {
		// 		if val < minValue {
		// 			minValue = val
		// 			minIndex = in + indexInnerLoop
		// 		}
		// 	}
		// 	nums[index], nums[minIndex] = minValue, value
		// }

		return nums
	}([]int{8, 3, 5, 0, 2, 7})

	fmt.Println("Sorting slice", sort)

	//closure function as a return value
	value := increament()
	fmt.Println(value())
	fmt.Println(value())

	//callback function
	dataLengthBiggerThan3 := filter([]string{"Hello", "Tes", "World"}, func(data string) bool {
		return len(data) > 3
	})
	fmt.Println(dataLengthBiggerThan3)
}

func greeting(name string) {
	fmt.Println("Hello " + name)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func basicMathOperation(a, b int) (int, int, int, int) {
	add := a + b
	sub := a - b
	multi := a * b
	div := a / b

	return add, sub, multi, div
}

func basicMathOperationWithPredefinedReturnValue(a, b int) (add, sub, multi, div int) {
	add = a + b
	sub = a - b
	multi = a * b
	div = a / b

	return
}

func sum(a ...int) (total int) {
	for _, value := range a {
		total += value
	}

	return
}

func increament() func() int {
	num := 0

	return func() int {
		num += 1
		return num
	}
}

func filter(data []string, callback func(string) bool) []string {
	var filteredData []string

	for _, value := range data {
		if filtered := callback(value); filtered {
			filteredData = append(filteredData, value)
		}
	}

	return filteredData
}
