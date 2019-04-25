package main

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
)

var invalidFormat = errors.New("Invalid question format.")
var notFound = errors.New("Question not found.")

type question struct {
	id     int
	prompt string
	answer string
}

func addQuestion(question *question) error {
	sql := `
		INSERT INTO questions (prompt, answer)
		VALUES($1, $2)
		RETURNING prompt, answer
	`

	_, err := db.Exec(sql, question.prompt, question.answer)
	return err
}

func allQuestions() ([]question, error) {
	sql := `
		SELECT id, prompt, answer
		FROM questions 
		ORDER BY id ASC
	`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var questions = []question{}
	for rows.Next() {
		question := question{}
		err := rows.Scan(&question.id, &question.prompt, &question.answer)
		if err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func EditQuestion(params string) error {
	args := strings.Split(params, ":")
	if invalidEditQuestionArgs(args) {
		return invalidFormat
	}

	questions, err := allQuestions()
	if err != nil {
		return err
	}

	selected, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	if selected > len(questions) {
		return notFound
	}

	index := selected - 1
	question := questions[index]
	question.prompt = args[1]
	question.answer = args[2]

	return question.update()
}

func invalidEditQuestionArgs(args []string) bool {
	return len(args) != 3 && invalidQuestionArgs(args[1:])
}

func invalidQuestionArgs(args []string) bool {
	return len(args) != 2 || len(string(args[0])) == 0 || len(string(args[1])) == 0
}

func listAllQuestions() (string, error) {
	questions, err := allQuestions()
	if err != nil {
		return "", err
	}

	var stringified []string
	for i, question := range questions {
		num := i + 1
		str := fmt.Sprintf("%d. %s", num, question.String())
		stringified = append(stringified, str)
	}

	joined := strings.Join(stringified, "\n")

	return joined, nil
}

func NewQuestion(params string) error {
	args := strings.Split(params, ":")
	if invalidQuestionArgs(args) {
		return invalidFormat
	}
	question := &question{prompt: args[0], answer: args[1]}

	return addQuestion(question)
}

// Question methods
func (q *question) String() string {
	return fmt.Sprintf("%s: %s", q.prompt, q.answer)
}

func (q *question) update() error {
	sql := `
		UPDATE questions
		SET prompt = $1, answer = $2
		WHERE id = $3
	`
	_, err := db.Exec(sql, q.prompt, q.answer, q.id)

	return err
}
