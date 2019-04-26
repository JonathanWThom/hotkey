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
var add = flag.String("add", "", "Add a new question with the following format: \"my prompt:my answer\"")
var all = flag.Bool("all", false, "List all questions and answers")
var edit = flag.String("edit", "", "Edit question with the following format: \"1:my new prompt:my new answer\"")
var delete = flag.Int("delete", 0, "Delete question by number, corresponding with how it appears in the -all list.")

func main() {
	// Connect to database
	var err error
	db, err = sql.Open("postgres", "dbname=hotkey sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Flags
	flag.Parse()

	// refactor similar flag methods?
	if *add != "" {
		err := NewQuestion(*add)
		if err != nil {
			fmt.Printf("Unable to add question. Error: %v\n", err)
			return
		}

		fmt.Println("Question successfully added.")
		return
	}

	if *all {
		list, _ := listAllQuestions()
		fmt.Println(list)
		return
	}

	if *edit != "" {
		err := EditQuestion(*edit)
		if err != nil {
			fmt.Printf("Unable to edit question. Error: %v\n", err)
			return
		}

		fmt.Println("Question successfully updated.")
		return
	}

	if *delete != 0 {
		err := DeleteQuestion(*delete)
		if err != nil {
			fmt.Printf("Unable able to delete question. Error: %v\n", err)
			return
		}

		fmt.Println("Question successfully deleted.")
		return
	}

	// Process Q & A
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

// Test reads input while in the repl, and responds to a correct answer, an
// incorrect answer, a request for a hint, or a request for the solution.
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
