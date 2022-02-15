// RECURSIVE AND ITERATIVE BINARY SEARCH OF A NUMBER IN A ARRAY BY DANI95RICO

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	find := binarySearch(slice, 5, 0, len(slice)-1)
	fmt.Println("Elemento encontrado por función recursiva de binarySearch en la posición: ", find)

	find = iterBinarySearch(slice, 5, 0, len(slice)-1)
	fmt.Println("Elemento encontrado por función iterativa de binarySearch en la posición: ", find)

}

func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	start := lowIndex
	end := highIndex
	var mid int
	for start <= end {
		mid = int((start + end) / 2)
		if array[mid] > target {
			end = mid
		} else if array[mid] < target {
			start = mid
		} else {
			return mid
		}
	}
	return -1
}
