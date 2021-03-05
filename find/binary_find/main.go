package main

import "fmt"

func BinaryFind(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println(" cannot find it")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("find it %v \n ", middle)
	}
}
func main() {
	arr := []int{1, 8, 10, 19, 25, 1000, 1652}
	BinaryFind(&arr, 0, len(arr)-1, 1653)
}
