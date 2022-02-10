package main

import (
	"fmt"
)

func AscendingSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func DescendingSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] < s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func main() {
	slice := []int{}
	//fmt.Printf("%T\n", slice)
	var n int
	fmt.Print("Enter the size of slice : ")
	fmt.Scan(&n)
	var ele int
	for i := 0; i < n; i++ {
		fmt.Print("Enter a number :- ")
		fmt.Scan(&ele)
		slice = append(slice, ele)

	}
	fmt.Print("Before sorting : ")
	fmt.Println(slice)
	fmt.Println("After Sorting in Ascending Order")
	result1 := AscendingSort(slice)
	fmt.Println(result1)
	fmt.Println()
	fmt.Println("After Sorting in Descending Order")
	result2 := DescendingSort(slice)
	fmt.Println(result2)

}
