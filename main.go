package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kotaoue/twtw/pkg/config"
)

var (
	initialize = flag.Bool("init", false, "initialize config file")
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

	if err := getTweet(); err != nil {
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

func getTweet() error {
	cf := config.NewConfig()
	cf.Load()
	api := anaconda.NewTwitterApiWithCredentials(cf.AccessToken, cf.AccessTokenSecret, cf.ConsumerKey, cf.ConsumerKeySecret)
	searchResult, err := api.GetSearch("golang", nil)
	if err != nil {
		return err
	}

	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}

	return nil
}
