package main

import (
	"fmt"
	"github.com/nurettintopal/leaderboard"
)

var rs = leaderboard.RedisSettings{
	Host:     "localhost:6379",
	Password: "",
}

func main() {
	score1, _ := leaderboard.New(rs).Show("board-1")
	fmt.Println("")
	fmt.Println("Board board-1: ", score1)

	newPlayer := leaderboard.Player{
		ID:    "player-1",
		Score: 200,
	}

	leaderboard.New(rs).Create("board-1", newPlayer)

	score2, _ := leaderboard.New(rs).Show("board-1")
	fmt.Println("")
	fmt.Println("Board board-1: ", score2)

}
