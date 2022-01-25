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
		fmt.Printf("\r  \033[36mcomputing\033[m %s ", s.Next())
		time.Sleep(delay)
	}
}
