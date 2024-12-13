package main

import (
	"fmt"
)

func main() {
	inp:=[]int{10,5,2,7,9,1}
	fmt.Println(mergeSort(inp))
	//fmt.Println(merge([]int{2, 5, 10}, []int{1, 7, 9}))
}
func mergeSort(inp []int) []int {
	if len(inp) <= 1 {
		return inp
	}
	mid := len(inp) / 2
	left := mergeSort(inp[:mid])
	right := mergeSort(inp[mid:])
	return merge(left, right)
}
func merge(inp1, inp2 []int) []int {
	l := len(inp1) + len(inp2)
	arr := make([]int, l)
	lp, rp := 0, 0
	for i := 0; i < len(arr); i++ {
		if lp == len(inp1) {
			arr[i] = inp2[rp]
			rp++
		} else if rp == len(inp2) {
			arr[i] = inp1[lp]
			lp++
		} else {
			if inp1[lp] < inp2[rp] {
				arr[i] = inp1[lp]
				lp++
			} else {
				arr[i] = inp2[rp]
				rp++
			}
		}
	}
	return arr
}
