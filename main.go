package main

import "fmt"
import "math/rand"
import "math"
import "time"

const (
	height = 15
	width  = 80
)

type Universe [][]bool

func NewUniverse() Universe {
	Universe := make([][]bool,height)
	//EinU := make([]bool, width)
	for i := range Universe{      
		Universe[i] = make([]bool,width)
	}
	return Universe
}

func (u Universe) show() {
	//fmt.Println("-------")
	for _, row := range u {
		for _, column := range row {
			if column {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 随机激活25%的细胞
func (u Universe) Seed() {
	num := int(math.Round(height * width * 0.25))
	for count := 0; count < num; {
		wrand := rand.Intn(width)
		hrand := rand.Intn(height)
		if u[hrand][wrand] {
			continue
		}
		u[hrand][wrand] = true
		count++
	}
}

func(u Universe) Alive(x,y int)bool{
    for x<0{
		x+=width
	}
	for y<0{
		y+=height
	}
	x = x%width
	y = y%height
	return u[y][x]
}

func (u Universe) Neighbors(x,y int)int{
	num:=0
	xslice:=[]int{x-1,x,x+1,x-1,x+1,x-1,x,x+1}
	yslice:=[]int{y-1,y-1,y-1,y,y,y+1,y+1,y+1}
	for i:=0;i<8;i++{
		if u.Alive(xslice[i],yslice[i]){
		num+=1
		}
	}
	return num
}

func(u Universe) Next(x,y int)bool{
	if u.Alive(x,y)&&u.Neighbors(x,y)<2{
		return false
	}else if u.Alive(x,y)&&(u.Neighbors(x,y)==2||u.Neighbors(x,y)==3){
		return u[y][x]
	}else if u.Alive(x,y)&&u.Neighbors(x,y)>3{
		return false
	}else if !u.Alive(x,y)&&u.Neighbors(x,y)==3{
		return true
	}
	return u.Alive(x,y)
}

func Step(a,b Universe){
	for x:=0;x<width;x++{
		for y:=0;y<height;y++{
			b[y][x]=a.Next(x,y)
		}
	}
}

func main() {
	u := NewUniverse()
    u.Seed()
	//u.Seed()
	//u.Seed()
	//u.Seed()
	//u.show()
	//fmt.Println(u.Alive(0,-1))
    //fmt.Println(u.Neighbors(33,655))
	//fmt.Println(u.Next(1,1))
	tempu:=NewUniverse()
	for{
		time.Sleep(time.Second)
		u.show()
		Step(u,tempu)
		u,tempu = tempu,u
		time.Sleep(time.Second)
		fmt.Print("\x0c")
	}
}
