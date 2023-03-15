package main

import (
	"fmt"
)

func hadamardProduct(m1 [3][3]int, m2 [3][3]int) [3][3]int {
	var result [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = m1[i][j] * m2[i][j]
		}
	}

	return result
}

func main() {
	var m1 [3][3]int = [3][3]int{{1, 3, 5}, {7, 9, 11}, {13, 15, 17}}
	var m2 [3][3]int = [3][3]int{{2, 4, 6}, {8, 10, 12}, {14, 16, 18}}

	fmt.Println("m1 = ", m1)
	fmt.Println("m2 = ", m2)

	var result = hadamardProduct(m1, m2)
	fmt.Println("m1 * m2 = ", result)
}
