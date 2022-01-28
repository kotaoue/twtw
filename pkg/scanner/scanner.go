package scanner

import (
	"bufio"
	"os"

	"github.com/kotaoue/go-tput"
)

func Scan(msg string, opts ...[]*tput.Option) string {
	to := []*tput.Option{}
	for _, opt := range opts {
		for _, v := range opt {
			to = append(to, v)
		}
	}

	if msg != "" {
		tput.Printf(to, "%s\n", msg)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() != "" {
			break
		}
	}

	return s.Text()
}
