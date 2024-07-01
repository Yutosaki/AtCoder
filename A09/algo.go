package main

import "fmt"
func main(){
	lows,cols, n:= 0, 0, 0
	fmt.Scan(&lows,&cols, &n)
	matrix:=make([][]int, lows+2)
	for i:=0;i<=lows;i++{
		matrix[i]=make([]int, cols+2)
	}
	for i:=0;i<n;i++{
		x1,y1,x2,y2:=0,0,0,0
		fmt.Scan(&x1,&y1,&x2,&y2)
		matrix[x1][y1]++
		matrix[x1][y2+1]--
		matrix[x2+1][y1]--
		matrix[x2+1][y2+1]++
	}
	
	for i:=1;i<=lows;i++{
		for j:=1;j<=cols;j++{
			matrix[i][j]+=matrix[i-1][j]
		}
	}

	for i:=1;i<=lows;i++{
		for j:=1;j<=cols;j++{
			matrix[i][j]+=matrix[i][j-1]
		}
	}

	for i:=1;i<=lows;i++{
		for j:=1;j<=cols;j++{
			fmt.Printf("%-2d",matrix[i][j])
		}
		fmt.Println()
	}
}
