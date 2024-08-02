package main


import (
    "fmt"
    //"encoding/json"
    "math/rand"
    "time"
)
type HoneyBee struct{
    Name string 
}
func (hb HoneyBee)String()string{
    return fmt.Sprintf("this is a %v",hb.Name)
}
func(hb HoneyBee)Move() string{
    switch rand.Intn(3){
        case 0:
            return "buzzes about"
        default:
            return "files to infinity and beyond"
    }
}
func(hb HoneyBee)Eat()string{
    switch rand.Intn(2){
        case 0:
            return "pollen"
        default:
            return "nectar"    
    }
}


type Gopher struct{
    Name string 
}
func (g Gopher)String()string{
    return fmt.Sprintf("this is a %v",g.Name)
}

func (g Gopher)Move() string{
    switch rand.Intn(2){
        case 0:
            return "scurries along the ground"
        default:
            return "burrows in the sand"    
    }
}

func(g Gopher) Eat() string{
    switch rand.Intn(5){
        case 0:
            return "carrot"
        case 1:
            return "lettuce"
        case 2:
            return "radish"
        case 3:
            return "corn"
        case 4:
            return "root"
        default:
            return ""                
    }
}

type animal interface{
    Move() string
    Eat() string
}

func step(a animal){
    switch rand.Intn(2){
        case 0:
            fmt.Printf("%v %v.\n",a,a.Move())
        case 1:
            fmt.Printf("%v eats %v.\n",a,a.Eat())    
    }
}

const sunrise,sunset = 8,18

func main(){
    rand.Seed(time.Now().UnixNano())

    animals:=[]animal{
        HoneyBee{Name:"New Bee"},
        Gopher{Name:"Go gopher"},
    }

    var sol,hour int
    loop:
    for{
        fmt.Printf("%2d:00 ",hour)
        if hour<sunrise||hour>sunset{
            fmt.Println("the animals are sleeping.")
        }else{
            step(animals[rand.Intn(len(animals))])
        }
        time.Sleep(500*time.Millisecond)
        hour++
        if hour>=24{
            hour=0
            sol++
            if sol>=3{
                break loop
            }
        }
    }
}