package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
)

var ctx = context.Background()

func main() {
	client := createRedisClient()
	defer client.Close()

	leaderboardService := NewLeaderboardService(client)

	var leaderboardName string
	for i := 1; i < 10; i++ {
		leaderboardName = fmt.Sprintf("leadershipboard-%d", i)

		addSamplePlayers(leaderboardService, leaderboardName, 5)

		displayLeaderboard(leaderboardService, leaderboardName)
		displayRank(leaderboardService, leaderboardName, "player3")

		fmt.Println("")
		fmt.Println("--------")
		fmt.Println("")
	}
}

func addSamplePlayers(service *LeaderboardService, leaderboardName string, count int) {
	for i := 1; i <= count; i++ {
		player := Player{
			ID:       fmt.Sprintf("player%d", i),
			Score:    rand.Intn(1000),
			Username: fmt.Sprintf("User%d", i),
		}

		if err := service.AddPlayer(leaderboardName, player); err != nil {
			log.Printf("Failed to add player %s to leaderboard: %v", player.ID, err)
		}
	}
}

func displayLeaderboard(service *LeaderboardService, leaderboardName string) {
	leaderboard, err := service.GetLeaderboard(leaderboardName)
	if err != nil {
		log.Fatalf("Failed to retrieve leaderboard: %v", err)
	}

	fmt.Printf("Board for %s:", leaderboardName)
	fmt.Println()
	for rank, member := range leaderboard {
		playerID := member.Member.(string)
		playerScore := int(member.Score)
		fmt.Printf("%d. %s - Score: %d\n", rank+1, playerID, playerScore)
	}
}

func displayRank(service *LeaderboardService, leaderboardName string, targetPlayerID string) {
	rank, err := service.GetPlayerRank(leaderboardName, targetPlayerID)
	if err != nil {
		log.Fatalf("Failed to retrieve rank for player %s: %v", targetPlayerID, err)
	}

	fmt.Printf("%s is ranked #%d on the leaderboard.\n", targetPlayerID, rank+1)
}
