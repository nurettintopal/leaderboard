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

type LeaderboardService struct {
	redisClient *redis.Client
}

func NewLeaderboardService(redisClient *redis.Client) *LeaderboardService {
	return &LeaderboardService{redisClient: redisClient}
}

func (ls *LeaderboardService) ClearLeaderboard() error {
	return ls.redisClient.Del(ctx, "leaderboard").Err()
}

func (ls *LeaderboardService) AddPlayer(player Player) error {
	return ls.redisClient.ZAdd(ctx, "leaderboard", &redis.Z{
		Score:  float64(player.Score),
		Member: player.ID,
	}).Err()
}

func (ls *LeaderboardService) GetLeaderboard() ([]redis.Z, error) {
	return ls.redisClient.ZRevRangeWithScores(ctx, "leaderboard", 0, -1).Result()
}

func (ls *LeaderboardService) GetPlayerRank(playerID string) (int64, error) {
	return ls.redisClient.ZRevRank(ctx, "leaderboard", playerID).Result()
}

func main() {
	client := createRedisClient()
	defer client.Close()

	leaderboardService := NewLeaderboardService(client)

	// Clear any previous data
	if err := leaderboardService.ClearLeaderboard(); err != nil {
		log.Fatalf("Failed to clear leaderboard: %v", err)
	}

	addSamplePlayers(leaderboardService, 100)

	displayLeaderboard(leaderboardService)
	displayRank(leaderboardService, "player2")
	displayRank(leaderboardService, "player4")
	displayRank(leaderboardService, "player8")
	displayRank(leaderboardService, "player16")
	displayRank(leaderboardService, "player32")
	displayRank(leaderboardService, "player64")

}

func createRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func addSamplePlayers(service *LeaderboardService, count int) {
	for i := 1; i <= count; i++ {
		player := Player{
			ID:       fmt.Sprintf("player%d", i),
			Score:    rand.Intn(1000),
			Username: fmt.Sprintf("User%d", i),
		}

		if err := service.AddPlayer(player); err != nil {
			log.Printf("Failed to add player %s to leaderboard: %v", player.ID, err)
		}
	}
}

func displayLeaderboard(service *LeaderboardService) {
	leaderboard, err := service.GetLeaderboard()
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

func displayRank(service *LeaderboardService, targetPlayerID string) {
	rank, err := service.GetPlayerRank(targetPlayerID)
	if err != nil {
		log.Fatalf("Failed to retrieve rank for player %s: %v", targetPlayerID, err)
	}

	fmt.Printf("%s is ranked #%d on the leaderboard.\n", targetPlayerID, rank+1)
}
