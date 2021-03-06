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
	initialize  = flag.Bool("init", false, "initialize config file")
	commitTweet = flag.Bool("c", false, "commit tweet. if don't input a message after this args, external editor will run")
	tweetCount  = flag.Int("n", 10, "number of fetch tweets")

	tweetMessage *string
)

func init() {
	flag.Parse()
	s := flag.Arg(0)
	tweetMessage = &s
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	tput.Clear()
	switch {
	case *initialize:
		return initializeConfig()
	default:
		t, err := twitter.NewTwitter()
		if err != nil {
			return err
		}

		if *commitTweet {
			return t.Tweet(*tweetMessage)
		}

		return t.HomeTimeline(*tweetCount)
	}
}

func initializeConfig() error {
	var opts []*tput.Option
	opts = append(opts, &tput.Option{Attribute: tput.TextColor, Color: tput.Magenta})
	opts = append(opts, &tput.Option{Attribute: tput.UnderLine})
	opts = append(opts, &tput.Option{Attribute: tput.BoldText})

	cfg := config.NewConfig()
	cfg.AccessToken = scanner.Scan(">>> please input your Access Token", opts)
	cfg.AccessTokenSecret = scanner.Scan(">>> please input your Access Token Secret", opts)
	cfg.ConsumerKey = scanner.Scan(">>> please input your Consumer Key", opts)
	cfg.ConsumerKeySecret = scanner.Scan(">>> please input your Consumer Key Secret", opts)
	return cfg.Save()
}
