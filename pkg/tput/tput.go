package tput

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func HR() {
	i, _ := terminalColumnLength()
	fmt.Println(strings.Repeat("-", i))
}

func terminalColumnLength() (int, error) {
	s, err := exec.Command("tput", "cols").Output()
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(strings.Trim(string(s), "\n"))
	if err != nil {
		return 0, err
	}

	return i, nil
}
