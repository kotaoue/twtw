package spinner

import (
	"fmt"
	"time"

	"github.com/tj/go-spin"
)

func Spin(delay time.Duration) {
	s := spin.New()
	s.Set(spin.Spin3)

	for {
		fmt.Printf("\r fetching %s ", s.Next())
		time.Sleep(delay)
	}
}
