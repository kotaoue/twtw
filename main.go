package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kotaoue/twtw/pkg/config"
	"github.com/kotaoue/twtw/pkg/tput"
)

var (
	initialize = flag.Bool("init", false, "initialize config file")
	excludeRT  = flag.Bool("ex", false, "exclude RT")
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
		initializeConfig()
	}

	config := config.NewConfig()
	if err := config.Load(); err != nil {
		return err
	}

	if err := getHomeTimeline(); err != nil {
		return err
	}

	// tweet()

	return nil
}

func initializeConfig() {
	cf := config.NewConfig()
	cf.Save(config.ConfigJson{
		ConsumerKey:       scanText("please input your Consumer Key"),
		ConsumerKeySecret: scanText("please input your Consumer Key Secret"),
		AccessToken:       scanText("please input your Access Token"),
		AccessTokenSecret: scanText("please input your Access Token Secret"),
	})
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
	cf := config.NewConfig()
	cf.Load()

	api := anaconda.NewTwitterApiWithCredentials(cf.AccessToken, cf.AccessTokenSecret, cf.ConsumerKey, cf.ConsumerKeySecret)

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
		fmt.Printf("\x1b[36m%s\x1b[0m\n", tweet.User.Name)
		fmt.Printf("%s\n", tweet.FullText)
		fmt.Printf("\x1b[4m\x1b[34mhttps://twitter.com/%s/status/%s\x1b[0m\n", tweet.User.IdStr, tweet.IdStr)
	}

	return nil
}
