leaderboard
==============================================
a leaderboard implementation in Golang

## what is a leaderboard?

> The concept of a leaderboard—a scoreboard showing the ranked names and current scores (or other data points) of the leading competitors—is essential to the world of computer gaming, but leaderboards are now about more than just games. They are about gamification, a broader implementation that can include any group of people with a common goal (coworkers, students, sales groups, fitness groups, volunteers, and so on).

if you want to look into the details, follow [this link](https://redis.com/solutions/use-cases/leaderboards/), please.

## usage
* [a basic usage of leaderboard package in Golang](https://github.com/nurettintopal/leaderboard/tree/main/examples/basic)
* [a basic usage of leaderboard package with Fiber in Golang](https://github.com/nurettintopal/leaderboard/tree/main/examples/api)

## features
* Multiple leaderboard by name or any key
* Adding a new user and score
* Finding out any user's ranking
* You can get any part of the score list you want

## dependencies
* Go
* Redis client (github.com/go-redis/redis/v8)

## contributing
* if you want to add anything, contributions are welcome.
* Open a pull request that has your explanations

## license
leaderboard is open-sourced software licensed under the [MIT license](LICENSE).
