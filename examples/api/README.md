leaderboard
==============================================
a basic usage of leaderboard package with Fiber in Golang

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

## dependencies
* Go
* Fiber Framework
* leaderboard(github.com/nurettintopal/leaderboard)
* Redis client (github.com/go-redis/redis/v8)