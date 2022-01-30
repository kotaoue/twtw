# twtw
CLI Twitter Client

## Preparation
* Access [Developer Portal](https://developer.twitter.com/en/portal/projects-and-apps) And Get Bearer Token.
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
# message specified by argument
go run main.go -c -m "Hello World"

# message inputs by vim
go run main.go -c
```
## Links
* Twitter
  * [Twitter API](https://developer.twitter.com/en/docs/twitter-api)
  * [Getting started](https://developer.twitter.com/en/docs/twitter-api/getting-started/getting-access-to-the-twitter-api)
  * [Developer Portal](https://developer.twitter.com/en/portal/projects-and-apps)
  * [Twitter API](https://developer.twitter.com/en/docs/twitter-api)
* [ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda)

## License
See the [LICENSE](LICENSE) file for license rights and limitations (MIT).
