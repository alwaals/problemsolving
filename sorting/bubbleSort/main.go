package main

import "fmt"

func main(){

	inp:=[]int{6,3,4,2,7}
	fmt.Println(bubble(inp))
}

func bubble(inp []int)[]int{
	 for i:=0;i<len(inp);i++{
		for j:=len(inp)-1;j>0;j--{
			if inp[j]<inp[j-1]{
				inp[j],inp[j-1]=inp[j-1],inp[j]
			}
		}
	 }
	 fmt.Println(inp)
	return inp
}