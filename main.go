package main

import "fmt"
import "math/rand"

type kelvin float64

type sensor func() kelvin

func realsensor() kelvin{
    return 0
}

func fakesensor() kelvin{
    return kelvin(rand.Intn(50)+50)
}

func calibrate(s sensor, offset kelvin) sensor{
    return func() kelvin{
        //offset+=100
        return s()+offset
    }
}

func main() {
    var offset kelvin = 5
    sensor:=calibrate(fakesensor,offset)

    for count:=0;count<10;count++{
        fmt.Println(sensor())
        //offset+=100
    }
}
