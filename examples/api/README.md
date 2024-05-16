docker build -t leaderboard-api -f Dockerfile .

docker run -p 1919:1919 --env-file=.env leaderboard-api:latest