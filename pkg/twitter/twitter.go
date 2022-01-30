package twitter

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/davecgh/go-spew/spew"
	"github.com/kotaoue/go-eeditor"
	"github.com/kotaoue/go-tput"
	"github.com/kotaoue/twtw/pkg/config"
	"github.com/kotaoue/twtw/pkg/spinner"
)

type Twitter struct {
	api *anaconda.TwitterApi
}

func NewTwitter() (*Twitter, error) {
	t := &Twitter{}

	if err := t.init(); err != nil {
		return t, err
	}

	return t, nil
}

func (t *Twitter) init() error {
	api, err := t.apiWithCredentials()
	if err != nil {
		return err
	}
	t.api = api
	return nil
}

func (t *Twitter) Tweet(msg string) error {
	if msg == "" {
		editor := eeditor.NewEditor()
		b, _ := editor.Open()
		msg = string(b)
	}

	if msg != "" {
		tweet, err := t.api.PostTweet(msg, url.Values{})
		if err != nil {
			return err
		}

		spew.Dump(tweet)
		return nil
	}

	return errors.New("message is nil")
}

func (*Twitter) apiWithCredentials() (*anaconda.TwitterApi, error) {
	cfg, err := config.NewConfig().Load()
	if err != nil {
		return nil, err
	}

	return anaconda.NewTwitterApiWithCredentials(cfg.AccessToken, cfg.AccessTokenSecret, cfg.ConsumerKey, cfg.ConsumerKeySecret), nil
}

func (t *Twitter) HomeTimeline() error {
	searchResult, err := t.getHomeTimeline()
	if err != nil {
		return err
	}

	time.Sleep(1 * time.Microsecond)
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

func (t *Twitter) getHomeTimeline() ([]anaconda.Tweet, error) {
	tput.Setaf(tput.Green)
	defer tput.Sgr0()

	go spinner.Spin(100 * time.Millisecond)

	return t.api.GetHomeTimeline(url.Values{})
}
