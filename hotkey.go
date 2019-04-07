package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	successMsg = "Correct!\n"
	failureMsg = "Not quite, try again. Type 'solution' to see the answer.\n"
	solution   = "solution"
)

var stdin io.Reader

func main() {
	stdin = os.Stdin
	for {
		rand.Seed(time.Now().Unix())
		hotkey := hotkeys[rand.Intn(len(hotkeys))]
		Test(hotkey, os.Stdout)
	}
}

func Test(h hotkey, stdout io.Writer) {
	fmt.Fprintln(stdout, h.description)

	reader := bufio.NewReader(stdin)
	fmt.Fprint(stdout, "-> ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(stdout, err)
		return
	}

	input = strings.TrimSuffix(input, "\n")
	if input == solution {
		fmt.Fprintf(stdout, "%s\n\n", h.command)
		return
	}

	if input != h.command {
		fmt.Fprintln(stdout, failureMsg)
		Test(h, stdout)
	} else {
		fmt.Fprintln(stdout, successMsg)
	}
}
