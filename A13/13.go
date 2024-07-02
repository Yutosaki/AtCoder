package main

import "fmt"
func main(){
	n,diff :=0,0
	fmt.Scan(&n,&diff)
	array:=make([]int,n)

	counter:=0
	for i:=0;i<n;i++{
		fmt.Scan(&array[i])
	}
	
	for i:=1;i<n;i++{
		left_p,right_p:=0,i-1
		for left_p<=right_p{
			mid := (left_p+right_p)/2
			if array[i]-array[mid] <=diff{
				right_p=mid-1
			}else{
				left_p=mid+1
			}
		}
		counter+=i - left_p
	}
	fmt.Println(counter)
}