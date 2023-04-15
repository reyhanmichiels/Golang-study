package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"study/GoProgrammingLanguage/ch2/ex2.2/lengthConv"
)

func main() {
	if len(os.Args[1:]) == 0 {

		input := bufio.NewScanner(os.Stdin)

		for input.Scan() {

			number, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "lengthconv: %v", err)
			}

			calculate(number)

		}

	} else {

		for _, val := range os.Args[1:] {

			number, err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "lengthconv: %v", err)
			}

			calculate(number)

		}
	}
}

func calculate(val float64) {

	km := lengthConv.KiloMeter(val)
	m := lengthConv.Meter(val)

	fmt.Printf("%s = %s , %s = %s\n", km, lengthConv.KmtoM(km), m, lengthConv.MtoKm(m))

}
