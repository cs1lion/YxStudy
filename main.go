package main

import "fmt"

type celsius float64

func(c celsius) fahrenhelt() fahrenhelt{
    return fahrenhelt((c*9.0/5.0)+32.0)
}
func(c celsius) kelvin() kelvin{
    return kelvin(c+273.15)
}

type fahrenhelt float64

func (f fahrenhelt) celsius() celsius{
    return celsius((f-32.0)*5.0/9.0)
}

func(f fahrenhelt) kelvin() kelvin{
    return f.celsius().kelvin()
}

type kelvin float64

func(k kelvin) celsius() celsius{
    return celsius(k-273.15)
}

func(k kelvin) fahrenhelt() fahrenhelt{
    return k.celsius().fahrenhelt()
}

func main(){
    var f fahrenhelt = 255.2
    fmt.Println("f:",f,"f to c:",f.celsius())
}