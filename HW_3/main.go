package main

import (
	"flag"
	"fmt"
)

type Animal struct {
	runSpeed   int
	name       string
	voice      string
	ifWinSay   string
	ifLooseSay string
}

type Turtle struct {
	animal     Animal
	canSwim    bool
	haveShield bool
	entityName string
}

type Tiger struct {
	animal     Animal
	weight     int
	tailLength int
	entityName string
}

type Fish struct {
	animal       Animal
	countOfFins  int
	haveBigTeeth bool
	entityName   string
}

type Race struct {
	distanse       int
	turtle         Turtle
	tiger          Tiger
	fish           Fish
	turtleDistanse int
	tigerDistanse  int
	fishDistanse   int
	board          map[int]string
	lap            int
}

func (r *Race) CreateTeam(turtle Turtle, tiger Tiger, fish Fish) {
	r.turtle = turtle
	r.tiger = tiger
	r.fish = fish
	r.lap = 1
}

func (r *Race) IncreaseTurtleDistanse() {
	r.turtleDistanse += r.turtle.animal.runSpeed
}

func (r *Race) IncreaseTigerDistanse() {
	r.tigerDistanse += r.tiger.animal.runSpeed
}

func (r *Race) IncreaseFishDistanse() {
	r.fishDistanse += r.fish.animal.runSpeed
}

func (r Race) CheckResults(name string) bool {
	for _, v := range r.board {
		if name == v {
			return true
		}
	}
	return false
}

func (r Race) FinishPrinter(entity string, ifLooseSay string, ifWinSay string) {
	switch len(r.board) {
	case 1:
		fmt.Printf("%s won - race time is %d iterations\n**%s**\n", entity, r.lap, ifWinSay)
	case 2:
		fmt.Printf("%s came second - race time is %d iterations\n", entity, r.lap)
	case 3:
		fmt.Printf("%s came third -  race time is %d iterations\n**%s**\n", entity, r.lap, ifLooseSay)
	}

}

func (r *Race) LapCounter() {
	r.lap++
}

func main() {
	var (
		userTurtleSpeed int
		userTigerSpeed  int
		userFishSpeed   int
		userDistanse    int
	)

	userTurtleSpeed = *flag.Int("turtle", 12, "type speed to start race")
	userTigerSpeed = *flag.Int("tiger", 100, "type tiger speed to start race")
	userFishSpeed = *flag.Int("fish", 50, "type fish speed to start race")
	userDistanse = *flag.Int("d", 1000, "Distanse of race")
	flag.Parse()

	if userTurtleSpeed == userTigerSpeed || userTurtleSpeed == userFishSpeed || userTigerSpeed == userFishSpeed {
		fmt.Println("Arguments cant be equal")
		return
	}

	myTurtle := Turtle{
		animal: Animal{runSpeed: userTurtleSpeed,
			name:       "Bob",
			voice:      "-",
			ifWinSay:   "kawabanga!",
			ifLooseSay: "I lost again :("},
		canSwim:    false,
		haveShield: true,
		entityName: "Turtle",
	}

	myTiger := Tiger{
		animal: Animal{runSpeed: userTigerSpeed,
			name:       "Shere Khan",
			voice:      "rrrr",
			ifWinSay:   "AHHRGG!!!",
			ifLooseSay: "arhrrr"},
		weight:     333,
		tailLength: 100,
		entityName: "Tiger",
	}

	myFish := Fish{
		animal: Animal{runSpeed: userFishSpeed,
			name:       "Dory",
			voice:      "bloop",
			ifWinSay:   "where am I?",
			ifLooseSay: "where am I?"},
		countOfFins:  4,
		haveBigTeeth: false,
		entityName:   "Fish",
	}

	myRace := Race{
		turtle:   myTurtle,
		tiger:    myTiger,
		fish:     myFish,
		distanse: userDistanse,
		board:    make(map[int]string),
	}

	for {
		if !myRace.CheckResults(myTurtle.animal.name) {
			myRace.IncreaseTurtleDistanse()
			if myRace.turtleDistanse >= myRace.distanse {
				myRace.board[myRace.lap] = myTurtle.animal.name
				myRace.FinishPrinter(myTurtle.entityName, myTurtle.animal.ifLooseSay, myTurtle.animal.ifLooseSay)
			}

		}
		if !myRace.CheckResults(myTiger.animal.name) {
			myRace.IncreaseTigerDistanse()
			if myRace.tigerDistanse >= myRace.distanse {
				myRace.board[myRace.lap] = myTiger.animal.name
				if myRace.tigerDistanse >= myRace.distanse {
					myRace.board[myRace.lap] = myTiger.animal.name
					myRace.FinishPrinter(myTiger.entityName, myTiger.animal.ifLooseSay, myTiger.animal.ifLooseSay)
				}
			}

		}
		if !myRace.CheckResults(myFish.animal.name) {
			myRace.IncreaseFishDistanse()
			if myRace.fishDistanse >= myRace.distanse {
				myRace.board[myRace.lap] = myFish.animal.name
				if myRace.fishDistanse >= myRace.distanse {
					myRace.board[myRace.lap] = myFish.animal.name
					myRace.FinishPrinter(myFish.entityName, myFish.animal.ifLooseSay, myFish.animal.ifLooseSay)
				}
			}

		}
		myRace.LapCounter()
		if len(myRace.board) == 3 {
			break
		}
	}

}
