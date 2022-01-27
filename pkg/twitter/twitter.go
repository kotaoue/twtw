package twitter

import (
	"fmt"
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kotaoue/go-tput"
	"github.com/kotaoue/twtw/pkg/config"
	"github.com/kotaoue/twtw/pkg/scanner"
	"github.com/kotaoue/twtw/pkg/spinner"
)

func Tweet(msg string) error {
	if msg == "" {
		msg = scanner.Scan("please input your Access Token")
	}
	fmt.Println(msg)
	return nil
}

func apiWithCredentials() (*anaconda.TwitterApi, error) {
	cfg, err := config.NewConfig().Load()
	if err != nil {
		return nil, err
	}

	return anaconda.NewTwitterApiWithCredentials(cfg.AccessToken, cfg.AccessTokenSecret, cfg.ConsumerKey, cfg.ConsumerKeySecret), nil
}

func isTrigger(s string) bool {
	switch s {
	case ":w", ":send", ":post", ":tweet":
		return true
	}
	return false
}

func HomeTimeline() error {
	searchResult, err := getHomeTimeline()
	if err != nil {
		return err
	}

	tput.Clear()
	for _, tweet := range searchResult {
		tput.HR()
		{
			var opt []*tput.Option
			opt = append(opt, &tput.Option{Attribute: tput.TextColor, Color: tput.Cyan})
			tput.Printf(opt, "%s\n", tweet.User.Name)
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

func getHomeTimeline() ([]anaconda.Tweet, error) {
	tput.Setaf(tput.Green)
	defer tput.Sgr0()

	go spinner.Spin(100 * time.Millisecond)

	api, err := apiWithCredentials()
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	return api.GetHomeTimeline(v)
}
