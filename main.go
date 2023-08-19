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
	boardCount := 1
	playerCount := 100
	targetPlayerID := "player-5"
	for i := 1; i <= boardCount; i++ {
		leaderboardName = fmt.Sprintf("board-%d", i)

		clearLeaderboard(leaderboardService, leaderboardName)
		addSamplePlayers(leaderboardService, leaderboardName, playerCount)

		rank := displayRank(leaderboardService, leaderboardName, targetPlayerID)

		//displayLeaderboard(leaderboardService, leaderboardName)
		fmt.Println()
		displayLeaderboardWithRange(leaderboardService, leaderboardName, 0, 2)
		fmt.Println("...")
		displayLeaderboardWithRange(leaderboardService, leaderboardName, rank-2, rank+2)
		fmt.Println("...")
		displayLeaderboardWithRange(leaderboardService, leaderboardName, int64(int64(playerCount)-3), int64(playerCount))

		fmt.Println("")
		fmt.Println("--------")
		fmt.Println("")
	}
}

func addSamplePlayers(service *LeaderboardService, leaderboardName string, count int) {
	for i := 1; i <= count; i++ {
		player := Player{
			ID:       fmt.Sprintf("player-%d", i),
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

	for rank, member := range leaderboard {
		playerID := member.Member.(string)
		playerScore := int(member.Score)
		fmt.Printf("%d. %s - Score: %d\n", rank+1, playerID, playerScore)
	}
}

func displayLeaderboardWithRange(service *LeaderboardService, leaderboardName string, startOffset int64, endOffset int64) {
	leaderboard, err := service.GetLeaderboardWithRange(leaderboardName, startOffset, endOffset)

	if err != nil {
		log.Fatalf("Failed to retrieve leaderboard: %v", err)
	}

	for rank, member := range leaderboard {
		playerID := member.Member.(string)
		playerScore := int(member.Score)
		fmt.Printf("%d. %s - Score: %d\n", startOffset+int64(rank)+1, playerID, playerScore)
	}
}

func displayRank(service *LeaderboardService, leaderboardName string, targetPlayerID string) int64 {
	rank, err := service.GetPlayerRank(leaderboardName, targetPlayerID)
	if err != nil {
		log.Fatalf("Failed to retrieve rank for player %s: %v", targetPlayerID, err)
	}

	fmt.Printf("%s is ranked #%d on the leaderboard(%s).\n", targetPlayerID, rank+1, leaderboardName)

	return rank
}

func clearLeaderboard(service *LeaderboardService, leaderBoardName string) {
	err := service.ClearLeaderboard(leaderBoardName)
	if err != nil {
		log.Fatalf("Failed to clear the board %s: %v", leaderBoardName, err)
	}
}
