package twitter

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kotaoue/go-tput"
	"github.com/kotaoue/twtw/pkg/config"
	"github.com/kotaoue/twtw/pkg/spinner"
)

func Tweet() {
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

func HomeTimeline() error {
	searchResult, err := getHomeTimeline()
	if err != nil {
		return err
	}

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
	go spinner.Spin(100 * time.Millisecond)

	cfg, err := config.NewConfig().Load()
	if err != nil {
		return nil, err
	}

	api := anaconda.NewTwitterApiWithCredentials(cfg.AccessToken, cfg.AccessTokenSecret, cfg.ConsumerKey, cfg.ConsumerKeySecret)
	v := url.Values{}
	return api.GetHomeTimeline(v)
}
