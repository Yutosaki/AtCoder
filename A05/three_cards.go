package main

import "fmt"

func main(){
	n, k:= 0,0
	fmt.Scan(&n,&k)
	counter := 0
	for blue:=1;blue<=n;blue++{
		for red := 1; red<=n;red++{
			if k - blue - red <= n && 0 < k - blue - red{
				counter++
			}
		}
	}
	fmt.Println(counter)
}