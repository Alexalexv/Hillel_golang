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
}

type Tiger struct {
	animal     Animal
	weight     int
	tailLength int
}

type Fish struct {
	animal       Animal
	countOfFins  int
	haveBigTeeth bool
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

func (r Race) BoardChecker(name string) bool {
	for _, v := range r.board {
		if name == v {
			return true
		}
	}
	return false
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
		animal:     Animal{runSpeed: userTurtleSpeed, name: "Боб", voice: "-", ifWinSay: "Кавабанга!", ifLooseSay: "Ой, я знову програв :("},
		canSwim:    false,
		haveShield: true,
	}

	myTiger := Tiger{
		animal:     Animal{runSpeed: userTigerSpeed, name: "Шерхан", voice: "рррррр", ifWinSay: "АРГХХХХ!!", ifLooseSay: "аргххххх"},
		weight:     333,
		tailLength: 100,
	}

	myFish := Fish{
		animal:       Animal{runSpeed: userFishSpeed, name: "Дорі", voice: "блуп", ifWinSay: "Де я?", ifLooseSay: "Де я?"},
		countOfFins:  4,
		haveBigTeeth: false,
	}

	myRace := Race{
		turtle:   myTurtle,
		tiger:    myTiger,
		fish:     myFish,
		distanse: userDistanse,
		board:    make(map[int]string),
	}

	for {
		if myRace.BoardChecker(myTurtle.animal.name) == false {
			myRace.IncreaseTurtleDistanse()
			if myRace.turtleDistanse >= myRace.distanse {
				myRace.board[myRace.lap] = myTurtle.animal.name
				switch len(myRace.board) {
				case 1:
					fmt.Printf("Черепаха перемогла - час забігу становить %d ітерацій\n**%s**\n", myRace.lap, myTurtle.animal.ifWinSay)
				case 2:
					fmt.Printf("Другою прийшла черепаха - час забігу становить %d ітерацій\n", myRace.lap)
				case 3:
					fmt.Printf("Третя черепаха - час забігу становить %d ітерацій\n**%s**\n", myRace.lap, myTurtle.animal.ifLooseSay)
				}
			}

		}
		if myRace.BoardChecker(myTiger.animal.name) == false {
			myRace.IncreaseTigerDistanse()
			if myRace.tigerDistanse >= myRace.distanse {
				myRace.board[myRace.lap] = myTiger.animal.name
				switch len(myRace.board) {
				case 1:
					fmt.Printf("Тигр переміг - час забігу становить %d ітерацій\n**%s**\n", myRace.lap, myTiger.animal.ifWinSay)
				case 2:
					fmt.Printf("Другим прийшов тигр - час забігу становить %d ітерацій\n", myRace.lap)
				case 3:
					fmt.Printf("Третій тигр - час забігу становить %d ітерацій\n**%s**\n", myRace.lap, myTiger.animal.ifLooseSay)
				}
			}

		}
		if myRace.BoardChecker(myFish.animal.name) == false {
			myRace.IncreaseFishDistanse()
			if myRace.fishDistanse >= myRace.distanse {
				myRace.board[myRace.lap] = myFish.animal.name
				switch len(myRace.board) {
				case 1:
					fmt.Printf("Риба перемогла - час забігу становить %d ітерацій\n**%s**\n", myRace.lap, myFish.animal.ifWinSay)
				case 2:
					fmt.Printf("Другою прийшла риба - час забігу становить %d ітерацій\n", myRace.lap)
				case 3:
					fmt.Printf("Третя риба - час забігу становить %d ітерацій\n**%s**\n", myRace.lap, myFish.animal.ifLooseSay)
				}
			}

		}
		myRace.LapCounter()
		if len(myRace.board) == 3 {
			break
		}
	}

}
