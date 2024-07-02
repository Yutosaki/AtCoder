package main

import (
	"fmt"
	"sort"
)

type numbers struct{
	number int
	index int
	compressioned int
}

func main(){
	n:=0
	fmt.Scan(&n)
	array:=make([]numbers,n+1)
	for i:=1;i<=n;i++{
		fmt.Scan(&array[i].number)
		array[i].index=i
	}
	
	sort.Slice(array ,func (i,j int)bool  {return array[i].number < array[j].number})

	comp := 0
	for i:=1;i<=n;i++{
		if array[i].number != array[i-1].number{
			comp++
		}
		array[i].compressioned=comp
	}

	sort.Slice(array, func(i, j int)bool{return array[i].index < array[j].index})

	for i:=1;i<=n;i++{
		fmt.Print(array[i].compressioned, " ")
	}
	fmt.Println()
}