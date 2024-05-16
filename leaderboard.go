package leaderboard

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

type Player struct {
	ID    string `json:"id"`
	Score int    `json:"score"`
}

type RedisSettings struct {
	Host     string
	Password string
}

type LeaderboardService struct {
	redisClient *redis.Client
}

func New(rs RedisSettings) *LeaderboardService {
	return &LeaderboardService{redisClient: createRedisClient(rs.Host, rs.Password)}
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

func createRedisClient(host string, password string) *redis.Client {

	//TODO: connection pooling???
	if redisClient == nil {
		fmt.Println("Redis connection...")
		redisClient = redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       0,
		})

	}

	//TODO: close the connection
	//defer redisClient.Close()

	return redisClient
}
