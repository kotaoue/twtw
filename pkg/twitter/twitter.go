package twitter

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/briandowns/spinner"
	"github.com/davecgh/go-spew/spew"
	"github.com/kotaoue/go-eeditor"
	"github.com/kotaoue/go-tput"
	"github.com/kotaoue/twtw/pkg/config"
	"github.com/kotaoue/twtw/pkg/wd"
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
		opt := eeditor.Path(fmt.Sprintf("%s/tmp", wd.Get()))
		editor := eeditor.NewEditor(opt)
		b, err := editor.Open()
		if err != nil {
			return err
		}
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

func (t *Twitter) HomeTimeline(cnt int) error {
	searchResult, err := t.getHomeTimeline(cnt)
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

func (t *Twitter) getHomeTimeline(cnt int) ([]anaconda.Tweet, error) {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Color("fgGreen")

	s.Start()
	defer s.Stop()

	v := url.Values{}
	v.Set("count", strconv.Itoa(cnt))
	return t.api.GetHomeTimeline(v)
}
