package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	for {
		rand.Seed(time.Now().Unix())
		hotkey := hotkeys[rand.Intn(len(hotkeys))]
		test(hotkey)
	}
}

func test(h hotkey) {
	fmt.Println(h.description)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	input = strings.TrimSuffix(input, "\n")
	if input == "solution" {
		fmt.Printf("%s\n\n", h.command)
		return
	}

	if input != h.command {
		fmt.Printf("%s\n\n", input)
		fmt.Println("Not quite, try again. Type 'solution' to see the answer.\n")
		test(h)
	} else {
		fmt.Println("Correct!\n")
	}
}
