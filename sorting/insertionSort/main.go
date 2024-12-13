package main

import "fmt"

func main(){

	num:=[]int{4,5,7,9}
	fmt.Println(insertionSort(num,6))
}
func insertionSort(inp []int,num int)[]int{
	inp = append(inp, num)
		for j:=len(inp)-1;j>0;j--{
			if inp[j]<inp[j-1]{
				inp[j],inp[j-1]=inp[j-1],inp[j]
			}
		}
	return inp
}