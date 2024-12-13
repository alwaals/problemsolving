package main

import "fmt"
func main(){
	inp:=[]int{5,10,3,2,8,1}
	fmt.Println(selectionSort(inp))
}

func selectionSort(inp []int)[]int{
	for i:=0;i<len(inp);i++{
		min:=inp[i]
		for j:=i+1;j<len(inp);j++{
			if inp[j]<min{
				min=inp[j]
				inp[j]=inp[i]
				inp[i]=min
			}
		}
	}
	return inp
}