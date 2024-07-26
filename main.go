package main

import "fmt"

//import "time"
import "math/big"

// const(
// 	Virgin Galactic = iota +1
// 	SpaceX
// 	Space Adventures
// )

func main() {
	LightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away")

	seconds := new(big.Int)
	seconds.Div(distance, LightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Println("That is", days, "days of travel at light speed")
}
