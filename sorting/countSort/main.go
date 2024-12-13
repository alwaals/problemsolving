package main

import (
	"fmt"
)

func main(){

	inp:="banana"
	fmt.Println(count(inp))
}

func count(inp string)string{
	freq:=[26]int{}
	for _,v:=range inp{
		index:=int(v-'a')
		freq[index]+=1
	}
	newStr:=""
	for i,v:=range freq{
		for j:=0;j<v;j++{
           newStr+=string(i+'a')
		}
	}
	return newStr
}