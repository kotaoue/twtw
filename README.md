# twtw
CLI Twitter Client for bash

# Preparation
1. Access [Developer Portal](https://developer.twitter.com/en/portal/projects-and-apps) And Get Bearer Token.
  * cf. https://developer.twitter.com/en/docs/twitter-api/getting-started/getting-access-to-the-twitter-api 

## Usage
### Initialize
```
go run main.go -init=true
```
### Get HomeTimeLine
```
go run main.go
```
### Post Tweet
```
go run main.go -c
go run main.go -c -m "Hello World"
```

# Links
* Twitter
  * [Twitter API](https://developer.twitter.com/en/docs/twitter-api)
  * [Getting started](https://developer.twitter.com/en/docs/twitter-api/getting-started/getting-access-to-the-twitter-api)
  * [Developer Portal](https://developer.twitter.com/en/portal/projects-and-apps)
  * [Twitter API](https://developer.twitter.com/en/docs/twitter-api)
    * [Manage Tweets](https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/introduction)
* [ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda)
* [bash:tip_colors_and_formatting](https://misc.flogisoft.com/bash/tip_colors_and_formatting)