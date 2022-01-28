package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kotaoue/go-tput"
	"github.com/kotaoue/twtw/pkg/config"
	"github.com/kotaoue/twtw/pkg/scanner"
	"github.com/kotaoue/twtw/pkg/twitter"
)

var (
	initialize   = flag.Bool("init", false, "initialize config file")
	commitTweet  = flag.Bool("c", false, "commit tweet")
	tweetMessage = flag.String("m", "", "message of wish to tweet")
)

func init() {
	flag.Parse()
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	tput.Clear()
	if *initialize {
		return initializeConfig()
	}

	if *commitTweet {
		return twitter.Tweet(*tweetMessage)
	}

	return twitter.HomeTimeline()
}

func initializeConfig() error {
	cfg := config.NewConfig()
	cfg.AccessToken = scanner.Scan(">>> please input your Access Token")
	cfg.AccessTokenSecret = scanner.Scan(">>> please input your Access Token Secret")
	cfg.ConsumerKey = scanner.Scan(">>> please input your Consumer Key")
	cfg.ConsumerKeySecret = scanner.Scan(">>> please input your Consumer Key Secret")
	return cfg.Save()
}
