package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

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

	tweet()

	return nil
}

func initializeConfig() {
	fmt.Println("please input your Bearer Token")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() != "" {
			break
		}
	}

	config := config.NewConfig()
	config.Save(s.Text())
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
