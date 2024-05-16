package main

import (
	"fmt"
	"github.com/nurettintopal/leaderboard"
	"nurettintopal/leaderboard/config"
	"testing"

	"launchpad.net/gocheck"
)

func Test(t *testing.T) {
	gocheck.TestingT(t)
}

type S struct{}

var _ = gocheck.Suite(&S{})
var rs = leaderboard.RedisSettings{
	Host:     config.Config("REDIS_ADDR"),
	Password: config.Config("REDIS_PASS"),
}

func (s *S) TearDownSuite(c *gocheck.C) {
	fmt.Println("test suites is shouting down.")
}

func (s *S) TestLeaderboard(c *gocheck.C) {
	fmt.Println("test a leaderboard")
}
