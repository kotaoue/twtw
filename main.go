package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/kotaoue/twtw/pkg/config"
	"github.com/kotaoue/twtw/pkg/twitter"
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
		if err := initializeConfig(); err != nil {
			return err
		}
	}

	if err := twitter.HomeTimeline(); err != nil {
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
