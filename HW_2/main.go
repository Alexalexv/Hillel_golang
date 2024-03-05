package main

import "fmt"

func main() {
	//3
	chess := TableGame{
		Location: "home",
	}
	chess.game.Category = "Board games"
	chess.game.Finished = true
	chess.game.TotalPlayers = 2
	chess.game.Name = "Chess"
	chess.game.Rounds = 1

	//4 (6)
	sl := make([]TableGame, 3)
	sl[0] = chess

	//5
	soccer := TableGame{
		Location: "soccer field",
		game: Game{
			TotalPlayers: 12,
			Name:         "Soccer",
			Category:     "by foot",
			Rounds:       3,
			Finished:     false,
		},
	}
	sl[1] = soccer

	tenis := TableGame{
		Location: "tenis field",
		game: Game{
			TotalPlayers: 2,
			Name:         "Tenis",
			Category:     "by arms",
			Rounds:       10,
			Finished:     true,
		},
	}
	sl[2] = tenis

	//7
	var mapGame = map[int]TableGame{}
	for i, v := range sl {
		mapGame[i] = v
	}

	//8
	for _, v := range mapGame {
		fmt.Println(v)
	}

	//9
	var sumTotalPlayers int
	for _, v := range mapGame {
		sumTotalPlayers += v.game.TotalPlayers
	}

	//10
	fmt.Printf("Total players count: %d\n", sumTotalPlayers)

}

// 1
type Game struct {
	TotalPlayers int
	Name         string
	Category     string
	Rounds       int
	Finished     bool
}

// 2
type TableGame struct {
	Location string
	game     Game
}
