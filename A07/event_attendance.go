package main

import "fmt"
func main(){
	d, n:=0, 0
	fmt.Scan(&d, &n)
	attendant := make([]int, d+2)
	for i:= 1; i <= n; i++{
		l,r:= 0, 0
		fmt.Scan(&l,&r)
		attendant[l]++
		attendant[r+1]--
	}
	for i:=1;i<=d;i++{
		attendant[i]+= attendant[i-1]
		fmt.Println(attendant[i])
	}
}