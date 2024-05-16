leaderboard
==============================================
a leaderboard implementation in Golang

## what is a leaderboard?

> The concept of a leaderboard—a scoreboard showing the ranked names and current scores (or other data points) of the leading competitors—is essential to the world of computer gaming, but leaderboards are now about more than just games. They are about gamification, a broader implementation that can include any group of people with a common goal (coworkers, students, sales groups, fitness groups, volunteers, and so on).

if you want to look into the details, follow [this link](https://redis.com/solutions/use-cases/leaderboards/), please.

## basic usage
```go
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


```


## usage with API

#### home:
```json
GET http://127.0.0.1:1919
        
response:
{
  "message": "Your request has been completed successfully.",
  "status": "success"
}
```

#### leaderboard detail:
```json
GET http://127.0.0.1:1919/api/leaderboards/{leaderboard-id}
        
response:
{
    "data":[
        {
            "Score":1000,
            "Member":"lb-1-user-4"
        },
        {
            "Score":300,
            "Member":"lb-1-user-1"
        },
        {
            "Score":100,
            "Member":"lb-1-user-2"
        },
        {
            "Score":10,
            "Member":"lb-1-user-3"
        }
    ],
    "message":"Your request has been completed successfully.",
    "status":"success"
}
```

#### delete leaderboard:
```json
DELETE http://127.0.0.1:1919/api/leaderboards/{leaderboard-id}
        
response:
204 No Content
```

#### players with range in leaderboard:
```json
GET http://127.0.0.1:1919/api/leaderboards/{leaderboard-id}/players?start={start}&end={end}
        
response:
{
    "data": [
        {
            "Score": 300,
            "Member": "lb-1-user-3"
        },
        {
            "Score": 200,
            "Member": "lb-1-user-2"
        },
        {
            "Score": 100,
            "Member": "lb-1-user-1"
        }
    ],
    "message": "Your request has been completed successfully.",
    "status": "success"
}
```

#### player rank in leaderboard:
```json
GET http://127.0.0.1:1919/api/leaderboards/{leaderboard-id}/players/{player-id}
        
response:
{
    "data": {
      "rank": 2
    },
    "message": "Your request has been completed successfully.",
    "status": "success"
}
```

#### add player score to leaderboard:
```json
POST http://127.0.0.1:1919/api/leaderboards/{leaderboard-id}/players/{player-id}

request:
{
  "score": 1000
}
response:
{
    "data": null,
    "message": "Your request has been completed successfully.",
    "status": "success"
}
```



## features
* Multiple leaderboard by name or any key
* Adding a new user and score
* Finding out any user's ranking
* You can get any part of the score list you want

## dependencies
* Go
* Fiber Framework
* Redis client (github.com/go-redis/redis/v8)

## screenshot
![leaderboard sample](https://github.com/nurettintopal/leaderboard/blob/main/docs/sample.png?raw=true)

## contributing
* if you want to add anything, contributions are welcome.
* Open a pull request that has your explanations

## license
leaderboard is open-sourced software licensed under the [MIT license](LICENSE).
