package main

import "fmt"
//import "time"
import "math/rand"

// const(
// 	Virgin Galactic = iota +1
// 	SpaceX
// 	Space Adventures
// )

func main() {
	const distance = 62100000
	fmt.Printf("%-17s%-5v%-11v%v\n","Spaceline","Days","Trip type","Price")
	fmt.Println("=======================================")
	for count:=10;count>0;count--{
		comynum:=rand.Intn(3)+1
		comyname:=""
		switch comynum{
			case 1:
			comyname = "Virgin Galactic"
			case 2:
			comyname = "SpaceX"
			case 3:
			comyname = "Space Adventures"
			default:
			comyname = "Unknown"
		}
		speed:=rand.Intn(15)+16
		Days:=(int)(distance/(speed*3600*24))
		tripnum:=rand.Intn(2)+1
		triptype:=""
		switch tripnum{
			case 1:
			triptype = "One-Way"
			case 2:
			triptype = "Round-trip"
		}
		pricerand:=rand.Intn(15)+36
		price:=pricerand*tripnum
		fmt.Printf("%-17s%4v %-11v$ %3v\n",comyname,Days,triptype,price)
	}
		

	
}
