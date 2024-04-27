package leaderboard

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"nurettintopal/leaderboard/config"
)

var ctx = context.Background()

type Player struct {
	ID    string `json:"id"`
	Score int    `json:"score"`
}

type LeaderboardService struct {
	redisClient *redis.Client
}

func New() *LeaderboardService {
	return &LeaderboardService{redisClient: createRedisClient()}
}

func (ls *LeaderboardService) Delete(leaderboardName string) error {
	return ls.redisClient.Del(ctx, leaderboardName).Err()
}

func (ls *LeaderboardService) Create(leaderboardName string, player Player) error {
	return ls.redisClient.ZAdd(ctx, leaderboardName, &redis.Z{
		Score:  float64(player.Score),
		Member: player.ID,
	}).Err()
}

func (ls *LeaderboardService) Show(leaderboardName string) ([]redis.Z, error) {
	return ls.redisClient.ZRevRangeWithScores(ctx, leaderboardName, 0, -1).Result()
}

func (ls *LeaderboardService) Players(leaderboardName string, startOffset int64, endOffset int64) ([]redis.Z, error) {
	return ls.redisClient.ZRevRangeWithScores(ctx, leaderboardName, int64(startOffset), int64(endOffset)).Result()
}

func (ls *LeaderboardService) Player(leaderboardName string, playerID string) (int64, error) {
	return ls.redisClient.ZRevRank(ctx, leaderboardName, playerID).Result()
}

func createRedisClient() *redis.Client {
	fmt.Println("Redis connection...")
	//TODO: connection pooling???
	return redis.NewClient(&redis.Options{
		Addr:     config.Config("REDIS_ADDR"),
		Password: config.Config("REDIS_PASS"),
		DB:       0,
	})
}
