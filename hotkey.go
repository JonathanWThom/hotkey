package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	_ "github.com/lib/pq"
)

const (
	successMsg = "Correct!\n"
	failureMsg = "Not quite, try again. Type 'solution' to see the answer.\n"
	solution   = "solution"
	hint       = "hint"
)

var stdin io.Reader
var db *sql.DB
var questions []question
var add = flag.String("add", "", "add a new question with the following format: \"my prompt:my answer\"")

func main() {
	var err error
	db, err = sql.Open("postgres", "dbname=hotkey sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	flag.Parse()
	if *add != "" {
		err := NewQuestion(*add)
		if err != nil {
			fmt.Printf("Unable add to add question. Error: %v\n", err)
			return
		}

		fmt.Println("Question successfully added.")
		return
	}

	stdin = os.Stdin
	questions, err := allQuestions()
	if err != nil {
		panic(err)
	}

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

	if input == hint {
		fmt.Fprintf(stdout, "%s\n\n", q.hint())
		Test(q, stdout)
		return
	}

	if input != q.answer {
		fmt.Fprintln(stdout, failureMsg)
		Test(q, stdout)
	} else {
		fmt.Fprintln(stdout, successMsg)
	}
}

func (q *question) hint() string {
	split := strings.Split(q.answer, " ")
	var newSplit []string

	for i, s := range split {
		var newStr string
		remainder := strings.Repeat("*", utf8.RuneCountInString(s)-1)
		if len(s) == 1 {
			newStr = "*"
		} else if i%2 == 0 {
			first := s[0]
			newStr = string(first) + remainder
		} else {
			last := s[len(s)-1]
			newStr = remainder + string(last)
		}
		newSplit = append(newSplit, newStr)
	}
	return strings.Join(newSplit, " ")
}
