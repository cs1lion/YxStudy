package main

import "fmt"
import "os"
import "errors"

const(
	rows=9
	columns=9
	empty=0
)

type Cell struct{
	num int
	modify bool
}

type Grid [rows][columns]Cell

var(
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invilid digit")
	ErrInRow  = errors.New("already in row")
	ErrInColumn = errors.New("already in column")
	ErrInRegion = errors.New("already in region")
	ErrModifyDigit = errors.New("initial digits cannot be modified")
)

func NewSudoKu(digits [rows][columns]int) *Grid{
	var grid Grid
	for i:=0;i<rows;i++{
		for j:=0;j<columns;j++{
			grid[i][j].num=digits[i][j]
			if digits[i][j]==empty{
				grid[i][j].modify=true	
			}
			
		}
		return &grid
	}
}

func(g *Grid)Set(row,column int,digit int)error{
	//switch case验证错误
	g[row][column].num=digit
	return nil
}

func(g *Grid)CLear(r,c int)error{
	//switch case验证错误
	g[r][c].num=empty
	return nil
}

func inBounds(r,c int)bool{
	if (r<0||r>=rows)||(c<0||c>=columns){
		return false
	}
	return true
}
func invilidDigit(digit int)bool{
	if digit>0&&digit<=9{
		return false
	}
	return true
}
func (g *Grid)modifyDigit(r,c int)bool{
	return g[r][c].modify
}
func (g *Grid)digitInRow(r,digit int)bool{
	for i:=0;i<columns;i++{
		if g[r][i].num==digit{
			return true
		}
	}
	return false
}
func (g *Grid)digitInColumn(c,digit int)bool{
	for i:=0;i<rows;i++{
		if g[i][c].num==digit{
			return true
		}
	}
	return false
}
func(g *Grid)digitInRegion(row,column,digit int)bool{
	//分成9个区域
	startRow,startColumn:=row/3*3,column/3*3
	for r:=startRow;r<startRow+3;r++{
		for c:=startColumn;c<startColumn+3;c++{
			if g[r][c]=digit{
				return true
			}
		}
	}
	return false
}
