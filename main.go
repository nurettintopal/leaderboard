package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Player struct {
	ID       string
	Score    int
	Username string
}

func main() {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	// Clear any previous data
	client.Del(ctx, "leaderboard")

	// Simulate adding players with random scores
	for i := 1; i <= 100; i++ {
		player := Player{
			ID:       fmt.Sprintf("player%d", i),
			Score:    rand.Intn(1000),
			Username: fmt.Sprintf("User%d", i),
		}

		// Add player to the leaderboard
		client.ZAdd(ctx, "leaderboard", &redis.Z{
			Score:  float64(player.Score),
			Member: player.ID,
		})
	}

	// Retrieve and display the leaderboard
	leaderboard, err := client.ZRevRangeWithScores(ctx, "leaderboard", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Leaderboard:")
	for rank, member := range leaderboard {
		playerID := member.Member.(string)
		playerScore := int(member.Score)
		fmt.Printf("%d. %s - Score: %d\n", rank+1, playerID, playerScore)
	}

	// Get the rank of a specific player
	targetPlayerID := "player5"
	rank, err := client.ZRevRank(ctx, "leaderboard", targetPlayerID).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s is ranked #%d on the leaderboard.\n", targetPlayerID, rank+1)
}
