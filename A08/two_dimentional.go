package main

import "fmt"
func main(){
	lows,cols := 0, 0
	fmt.Scan(&lows,&cols)
	matrix := make([][]int, lows +1)
	matrix[0] = make([]int, cols+1)
	for i:=1;i<=lows;i++{
		matrix[i] = make([]int, cols+1)
		for j:=1;j<=cols;j++{
			fmt.Scan(&matrix[i][j])
			matrix[i][j] += matrix[i][j-1]
		}
	}

	quesion := 0
	fmt.Scan(&quesion)
	answer:=make([]int,quesion)
	for i:=0;i<quesion;i++{
		x1,y1,x2,y2:=0,0,0,0
		fmt.Scan(&x1,&y1,&x2,&y2)
		for low := x1;low<=x2;low++{
			answer[i]+=matrix[low][y2] - matrix[low][y1-1]
		}
	}
	for i:=0;i<quesion;i++{
		fmt.Println(answer[i])
	}
}