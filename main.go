package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	tput "github.com/kotaoue/go-tput"
	"github.com/kotaoue/twtw/pkg/config"
)

var (
	initialize = flag.Bool("init", false, "initialize config file")
	excludeRT  = flag.Bool("ex", false, "exclude replies")
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
	if *initialize {
		if err := initializeConfig(); err != nil {
			return err
		}
	}

	if err := getHomeTimeline(); err != nil {
		return err
	}

	// tweet()

	return nil
}

func initializeConfig() error {
	cfg := config.NewConfig()
	cfg.AccessToken = scanText("please input your Access Token")
	cfg.AccessTokenSecret = scanText("please input your Access Token Secret")
	cfg.ConsumerKey = scanText("please input your Consumer Key")
	cfg.ConsumerKeySecret = scanText("please input your Consumer Key Secret")
	return cfg.Save()
}

func scanText(msg string) string {
	if msg != "" {
		fmt.Println(msg)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() != "" {
			break
		}
	}

	return s.Text()
}

func tweet() {
	fmt.Println("What's happening?")

	tweet := ""

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if isTrigger(s.Text()) {
			break
		}
		tweet = fmt.Sprintf("%s\n%s", tweet, s.Text())
	}
	fmt.Println(tweet)
}

func isTrigger(s string) bool {
	switch s {
	case ":w", ":send", ":post", ":tweet":
		return true
	}
	return false
}

func getHomeTimeline() error {
	cfg, err := config.NewConfig().Load()
	if err != nil {
		return err
	}

	api := anaconda.NewTwitterApiWithCredentials(cfg.AccessToken, cfg.AccessTokenSecret, cfg.ConsumerKey, cfg.ConsumerKeySecret)
	v := url.Values{}
	if *excludeRT {
		v.Set("exclude_replies", "true")
	}

	searchResult, err := api.GetHomeTimeline(v)
	if err != nil {
		return err
	}

	for _, tweet := range searchResult {
		tput.HR()
		{
			var opt []*tput.Option
			opt = append(opt, &tput.Option{Attribute: tput.TextColor, Color: tput.Cyan})
			tput.Printf(opt, "%s\n", "tput.Printf")
		}

		fmt.Printf("%s\n", tweet.FullText)

		{
			var opt []*tput.Option
			opt = append(opt, &tput.Option{Attribute: tput.TextColor, Color: tput.Blue})
			opt = append(opt, &tput.Option{Attribute: tput.UnderLine})
			tput.Printf(opt, "https://twitter.com/%s/status/%s\n", tweet.User.IdStr, tweet.IdStr)
		}
	}

	return nil
}
