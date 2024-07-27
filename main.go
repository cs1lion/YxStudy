package main

import "fmt"
import "math"
import "strconv"
func main(){
	v:=42
	if v>=0&&v<=math.MaxUint8{

	
	v8:=uint8(v)
	fmt.Println("converted:",v8)
	}

	intnum:=25649
	fmt.Print("con:",string(intnum))
	fmt.Print("con:",strconv.Itoa(intnum))
}