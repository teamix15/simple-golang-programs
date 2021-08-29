package main

import (
	"fmt"
	"math/rand"
)

func generateSlice(count, max, min int) []int {
	var slice []int
	for i := 0; i < count; i++ {
		slice = append(slice, rand.Intn(max-min)+min)
	}
	return slice
}

func filterSlice(slice []int) []int {
	var filteredSlice []int
	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 {
			filteredSlice = append(filteredSlice, slice[i])
		}
	}
	return filteredSlice
}

func main() {
	max := 100
	min := 9
	var arrayCount = 150
	for i := 0; i < arrayCount; i++ {
		var slice = generateSlice(rand.Intn(40-9)+9, max, min)
		filteredSlice := filterSlice(slice)
		fmt.Println(slice)
		fmt.Println(filteredSlice)
		fmt.Println("-----------------------------------------------------------------------------------------")
	}
}
