package main

import (
	"github.com/go-redis/redis/v8"
)

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

func (ls *LeaderboardService) ClearLeaderboard(leaderboardName string) error {
	return ls.redisClient.Del(ctx, leaderboardName).Err()
}

func (ls *LeaderboardService) AddPlayer(leaderboardName string, player Player) error {
	return ls.redisClient.ZAdd(ctx, leaderboardName, &redis.Z{
		Score:  float64(player.Score),
		Member: player.ID,
	}).Err()
}

func (ls *LeaderboardService) GetLeaderboard(leaderboardName string) ([]redis.Z, error) {
	return ls.redisClient.ZRevRangeWithScores(ctx, leaderboardName, 0, -1).Result()
}

func (ls *LeaderboardService) GetLeaderboardWithRange(leaderboardName string, startOffset int64, endOffset int64) ([]redis.Z, error) {
	return ls.redisClient.ZRevRangeWithScores(ctx, leaderboardName, int64(startOffset), int64(endOffset)).Result()
}

func (ls *LeaderboardService) GetPlayerRank(leaderboardName string, playerID string) (int64, error) {
	return ls.redisClient.ZRevRank(ctx, leaderboardName, playerID).Result()
}
