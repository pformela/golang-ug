package main

import (
	"fmt"
	"math/big"
	"strings"
)

func factorial(n int) big.Int {
	result := big.NewInt(1)

	for i := 1; i <= n; i++ {
		result = result.Mul(result, big.NewInt(int64(i)))
	}

	fmt.Println("Current factorial: ", result, " for number: ", n)

	return *result
}

func checkIfNumbersInFactorial(nums []int, factorial big.Int) bool {
	var formattedFactorial string = factorial.String()

	for _, num := range nums {
		if !strings.Contains(formattedFactorial, fmt.Sprintf("%d", num)) {
			return false
		}
	}

	fmt.Println(formattedFactorial, " contains all numbers from: ", nums)
	return true
}

func main() {
	var nickname string = "patfor"
	var number int
	var nums [6]int

	for i, letter := range nickname {
		number += int(nickname[i])
		nums[i] = number
		fmt.Printf("Letter %d is %c\n", nickname[i], letter)
	}

	var result bool = false
	var fact int = 333
	var currFactorial big.Int = factorial(fact)

	// for n := 1; n < 500; n++ {
	// 	currFactorial = factorial(n)
	// 	result = checkIfNumbersInFactorial(nums[:], currFactorial)
	// 	if result {
	// 		break
	// 	}
	// 	number++
	// }

	result = checkIfNumbersInFactorial(nums[:], currFactorial)

	fmt.Println(result)

	if result {
		fmt.Println("Found factorial:", fact, "!")
	} else {
		fmt.Println("No factorial found")
	}

}
