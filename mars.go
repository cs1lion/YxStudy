package main

import "fmt"
import "math/rand"

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(19) + 1
	fmt.Println(num)
    //homework
    fmt.Printf("Malacandra's speed is %v\n",56000000/(28*24))
}
