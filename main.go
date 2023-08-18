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
	client := createRedisClient()
	defer client.Close()

	clearLeaderboard(client)
	addSamplePlayers(client, 100)

	displayLeaderboard(client)
	displayRank(client, "player2")
	displayRank(client, "player4")
	displayRank(client, "player8")
	displayRank(client, "player16")
	displayRank(client, "player32")
	displayRank(client, "player64")
}

func createRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func clearLeaderboard(client *redis.Client) {
	err := client.Del(ctx, "leaderboard").Err()
	if err != nil {
		log.Fatalf("Failed to clear leaderboard: %v", err)
	}
}

func addSamplePlayers(client *redis.Client, count int) {
	for i := 1; i <= count; i++ {
		player := Player{
			ID:       fmt.Sprintf("player%d", i),
			Score:    rand.Intn(1000),
			Username: fmt.Sprintf("User%d", i),
		}

		addPlayerToLeaderboard(client, player)
	}
}

func addPlayerToLeaderboard(client *redis.Client, player Player) {
	err := client.ZAdd(ctx, "leaderboard", &redis.Z{
		Score:  float64(player.Score),
		Member: player.ID,
	}).Err()
	if err != nil {
		log.Printf("Failed to add player %s to leaderboard: %v", player.ID, err)
	}
}

func displayLeaderboard(client *redis.Client) {
	leaderboard, err := client.ZRevRangeWithScores(ctx, "leaderboard", 0, -1).Result()
	if err != nil {
		log.Fatalf("Failed to retrieve leaderboard: %v", err)
	}

	fmt.Println("Leaderboard:")
	for rank, member := range leaderboard {
		playerID := member.Member.(string)
		playerScore := int(member.Score)
		fmt.Printf("%d. %s - Score: %d\n", rank+1, playerID, playerScore)
	}
}

func displayRank(client *redis.Client, targetPlayerID string) {
	rank, err := client.ZRevRank(ctx, "leaderboard", targetPlayerID).Result()
	if err != nil {
		log.Fatalf("Failed to retrieve rank for player %s: %v", targetPlayerID, err)
	}

	fmt.Printf("%s is ranked #%d on the leaderboard.\n", targetPlayerID, rank+1)
}
