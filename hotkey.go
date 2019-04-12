package main

import (
	"bufio"
	"database/sql"
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
var db *sql.DB

func main() {
	db, err := sql.Open("postgres", "dbname=hotkey sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	stdin = os.Stdin
	for {
		rand.Seed(time.Now().Unix())
		question := questions[rand.Intn(len(questions))]
		Test(question, os.Stdout)
	}
}

func Test(q question, stdout io.Writer) {
	fmt.Fprintln(stdout, q.prompt)

	reader := bufio.NewReader(stdin)
	fmt.Fprint(stdout, "-> ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(stdout, err)
		return
	}

	input = strings.TrimSuffix(input, "\n")
	if input == solution {
		fmt.Fprintf(stdout, "%s\n\n", q.answer)
		return
	}

	if input != q.answer {
		fmt.Fprintln(stdout, failureMsg)
		Test(q, stdout)
	} else {
		fmt.Fprintln(stdout, successMsg)
	}
}
