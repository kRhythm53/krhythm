package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct{
	row int
	col int
	matrix [][]int
}
func (m Matrix)getRows() int{
	return m.row
}
func (m Matrix)getCols() int{
	return m.col
}
func(m *Matrix)setPosition(i int,j int,val int) {
	m.matrix[i][j] = val
}
func(m Matrix)addMatrix(m2 Matrix) Matrix{

	for i:=0;i<m.row;i++{
		for j:=0;j<m.col;j++{
			m.matrix[i][j] = m.matrix[i][j]+m2.matrix[i][j]
		}
	}
	return m
}
func(m Matrix) printToJson(){

	b,err := json.Marshal(m.matrix);
	if err==nil {fmt.Println(string(b))}
}
func newMatrix() Matrix{
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}

	m1 := Matrix{2,3,nil}
	m1.matrix = append(m1.matrix,row1)
	m1.matrix = append(m1.matrix,row2)

	return m1
}
func main(){
	var m1 = newMatrix()
	var m2 = newMatrix()
	m1.setPosition(0,1,5)
	fmt.Println(m1.getRows())
	fmt.Println(m2.getCols())
	fmt.Println(m1.addMatrix(m2))
	m1.printToJson()
}
