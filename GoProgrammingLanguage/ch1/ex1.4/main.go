//dup 2 Prints the count and text of line that appears more than once in input. It Reads from stdin or from a list of named files.
//task -> modify dup 2 to prints the names of all files in which each duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	occurs := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "Standard Input", occurs)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, occurs)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("there is %d %q at:\n", n, line)
			fmt.Println(occurs[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, name string, occurs map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		if ! strings.Contains(occurs[input.Text()], name) {
			temp := fmt.Sprintf(" %s\n", name)
			occurs[input.Text()] += temp
		}
	}
}

// func countLines(f *os.File, counts map[string]int){
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++;
// 	}
// }
