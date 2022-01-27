package scanner

import (
	"bufio"
	"fmt"
	"os"
)

func Scan(msg string) string {
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
