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
		BearerToken:       scanText("please input your Bearer Token"),
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
	api := anaconda.NewTwitterApiWithCredentials("your-access-token", "your-access-token-secret", "your-consumer-key", "your-consumer-secret")
	searchResult, err := api.GetSearch("golang", nil)
	if err != nil {
		return err
	}

	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}

	return nil
}
