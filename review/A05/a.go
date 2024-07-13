package main

import "fmt"

func main(){
	n,k:=0,0
	fmt.Scan(&n,&k)

	counter:=0
	for i:=1;i<=n;i++{
		for j:=1;j<=n;j++{
			if 1 <= k - i - j  && k - i - j <=n{
				counter++
			}
		}
	}
	fmt.Println(counter)
}
