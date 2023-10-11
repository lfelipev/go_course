package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, element := range numbers {
		if element%2 == 0 {
			fmt.Println(element, "is even")
		} else {
			fmt.Println(element, "is odd")
		}
	}
}
