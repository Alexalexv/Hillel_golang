package main

import "fmt"

type Animal struct {
	runSpeed   int
	name       string
	voice      string
	ifWinSay   string
	ifLooseSay string
}

func (a Animal) WinnerSay() {
	fmt.Println(a.ifWinSay)
}

func (a Animal) LoserSay() {
	fmt.Println(a.ifLooseSay)
}

type Turtle struct {
	animal     Animal
	canSwim    bool
	haveShield bool
}

func main() {
	my_turtle := Turtle{
		animal:     Animal{runSpeed: 12, name: "Bob", voice: "-", ifWinSay: "Kawabanga!", ifLooseSay: "Oh, I`m loose again"},
		canSwim:    false,
		haveShield: true,
	}

	board := map[int]bool{1: false, 2: false, 3: false}
	fmt.Printf("%+v\n", my_turtle)
	my_turtle.animal.LoserSay()
	my_turtle.animal.WinnerSay()
}
