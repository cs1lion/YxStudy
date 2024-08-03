package main

import (
	"fmt"
)

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v picks up a %v\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v has nothing to give\n", c.name)
	}
	if to.leftHand != nil {
		fmt.Printf("%v 's hand is full\n", to.name)
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v gives %v to %v\n", c.name, to.leftHand.name, to.name)
}

func (c *character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v is carring nothing", c.name)
	}
	return fmt.Sprintf("%v is carring %v", c.name, c.leftHand.name)
}
func main() {
	arthur := &character{name: "Arthur"}
	boom := &item{name: "boom!"}
	arthur.pickup(boom)

	knight := &character{name: "Knight1"}
	arthur.give(knight)

	fmt.Println(arthur)
	fmt.Println(knight)
}
