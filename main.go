package main

import "fmt"
type Planets []string

func (p Planets)terraform() {
    for i,v:=range p{
        p[i]=fmt.Sprintf("New %v",v)
    }
}

func main(){
    planets:=[]string{"a","b","Mars","c","Uranus","Neptune"}
    Planets(planets[2:3]).terraform()
    Planets(planets[4:]).terraform()
    fmt.Print(planets)
}