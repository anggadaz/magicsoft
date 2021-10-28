package main

import (
	"fmt"
)

func findMax(a []int) (max int) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func printBar(num []int) {
	max := findMax(num)
	for i := 0; i < max; i++ {
		for j := 0; j < len(num); j++ {
			if i >= (max - num[j]) {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	for i := 0; i < len(num); i++ {
		fmt.Print(num[i], " ")
	}
	fmt.Println()
}

func insertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
			printBar(arr)
		}
		arr[j+1] = key
	}
}

func revInsertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j = j - 1
			printBar(arr)
		}
		arr[j+1] = key
	}
}

func main() {
	num := []int{1, 4, 5, 6, 8, 2}
	//insertSort(num)
	revInsertSort(num)
	printBar(num)
}
