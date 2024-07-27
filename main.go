package main

import "fmt"

//import "time"

// const(
// 	Virgin Galactic = iota +1
// 	SpaceX
// 	Space Adventures
// )

func main() {
	message:="Lfdph,L vdz,L frqtxhuhg"
	for _,c:=range message{
		fmt.Printf("%c",c-3)

		if c>='a'&&c<='z'{
			c+=13
			if c>'z'{
				c-=26
			}

			//print
		}
	}
}