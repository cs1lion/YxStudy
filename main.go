package main

import "fmt"
import "time"
import "math/rand"

func main() {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")

	//homework
	var target = 43
	var num int
	for {
		num = rand.Intn(100)
		if num < target {
			fmt.Println("lower")
		} else if num > target {
			fmt.Println("higher")
		} else {
			fmt.Println("bingo!")
			break
		}
		time.Sleep(time.Second)
	}
}
