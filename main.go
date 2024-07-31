package main


import "fmt"
import "math"

type location struct{
	lat,long float64
}
//coordinate结构使用度分秒格式的坐标表示东西南北半球
type coordinate struct{
	d,m,s float64
	h rune
}

func newLocation(lat,long coordinate) location{
	return location{lat.decimal(),long.decimal()}
}

func(c coordinate) decimal() float64{
	sign:=1.0
	switch c.h{
		case 'S','W','s','w':
		sign=-1
	}
	return sign*(c.d+c.m/60+c.s/3600)
}

//world以千米为单位记录了行星的测定半径
type world struct{
	radius float64
}
//distance使用余弦球面定律计算两个位置之间的距离
func(w world) distance(p1,p2 location) float64{
	s1,c1:=math.Sincos(rad(p1.lat))
	s2,c2:=math.Sincos(rad(p2.lat))
	clong:= math.Cos(rad(p1.long-p2.long))
	return w.radius*math.Acos(s1*s2+c1*c2*clong)
}

//rad函数将角度转换为弧度
func rad(deg float64) float64{
	return deg*math.Pi/180
}

var(
	mars = world{radius:3389.5}
	earth = world{radius:6371}
)


func main(){
	spirit:= newLocation(coordinate{14,34,6.2,'S'},coordinate{175,28,21.5,'E'})
	fmt.Println("Spirit",spirit)
	opportunity := newLocation(coordinate{1,56,46.3,'S'},coordinate{354,28,24.2,'E'})
	fmt.Println("Opportunity",opportunity)
	fmt.Printf("Spirit to Opportunity %.2f km\n",mars.distance(spirit,opportunity))
}