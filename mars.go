package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	fmt.Println("The year is 2100,should you leap?")
	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}
	//homework展示随机日期
	for count := 10; count > 0; count-- {
		year := rand.Intn(2500) + 1
		mouth := rand.Intn(12) + 1
		daysInMouth := 31
		switch mouth {
		case 2:
			if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
				daysInMouth = 29
			} else {
				daysInMouth = 28
			}
		case 4, 6, 9, 11:
			daysInMouth = 30
		}
		day := rand.Intn(daysInMouth)
		fmt.Println(era, year, mouth, day)
	}

}
