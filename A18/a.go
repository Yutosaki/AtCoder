package main

import "fmt"
func main(){
	n,sum:=0, 0
	fmt.Scan(&n,&sum)
	array:=make([]int, n+1)
	for i:=1;i<=n;i++{
		fmt.Scan(&array[i])
	}

	dp:=make([][]bool, n+1)
	for i:=0;i<=n;i++{
		dp[i]=make([]bool, sum+1)
	}

	dp[0][0]=true
	for i:=1;i<=n;i++{
		for j:=0;j<=sum;j++{
			if j < array[i]{
				if dp[i-1][j]==true {
					dp[i][j]=true
				}else{
					dp[i][j]=false
				}
			}else{
				if dp[i-1][j-array[i]]==true || dp[i-1][j]==true{
					dp[i][j]=true
				}else{
					dp[i][j]=false
				}
			}
		}
	}
	if dp[n][sum]==true{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}