package main

import "fmt"

func reverse_slice(slice [][]int) [][]int {
	var tempSlice [][]int

	for i := len(slice) - 1; i >= 0; i-- {
		var temp []int
		for j := len(slice[i]) - 1; j >= 0; j-- {
			temp = append(temp, slice[i][j])
		}
		tempSlice = append(tempSlice, temp)
	}

	return tempSlice
}

func add_slices(slice1 [][]int, slice2 [][]int) [][]int {
	var tempSlice [][]int

	if len(slice1) != len(slice2) || len(slice1[0]) != len(slice2[0]) {
		panic("Slices are not the same size")
	}

	for i := 0; i < len(slice1); i++ {
		var temp []int
		for j := 0; j < len(slice1[0]); j++ {
			temp = append(temp, slice1[i][j]+slice2[i][j])
		}
		tempSlice = append(tempSlice, temp)
	}

	return tempSlice
}

func main() {
	slice := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println("slice = ", slice)
	fmt.Println("reverse_slice(slice) = ", reverse_slice(slice))
	fmt.Println("slice = ", slice)

	fmt.Println("add_slices(slice, slice) = ", add_slices(slice, reverse_slice(slice)))

}
