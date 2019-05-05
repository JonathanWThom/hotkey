package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var errInvalidFormat = errors.New("invalid question format")
var errNotFound = errors.New("question not found")

type question struct {
	id     int
	prompt string
	answer string
}

func addQuestion(question *question) error {
	question.cleanup()

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

// DeleteQuestion receives an int and deletes the corresponding question, based on what its number is in the
// allQuestions list.
func DeleteQuestion(selected int) error {
	questions, err := allQuestions()
	if err != nil {
		return err
	}

	err = questionInRange(selected, questions)
	if err != nil {
		return err
	}

	question := findQuestion(selected, questions)

	return question.destroy()
}

// EditQuestion receives a string and parses it, ultimately updating a question
// in the database
func EditQuestion(params string) error {
	args := strings.Split(params, ":")
	if invalidEditQuestionArgs(args) {
		return errInvalidFormat
	}

	questions, err := allQuestions()
	if err != nil {
		return err
	}

	selected, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	err = questionInRange(selected, questions)
	if err != nil {
		return err
	}

	question := findQuestion(selected, questions)
	question.prompt = args[1]
	question.answer = args[2]

	return question.update()
}

func findQuestion(selected int, questions []question) question {
	index := selected - 1
	return questions[index]
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

// NewQuestion receives a string, and then builds a new question that is saved to
// the database.
func NewQuestion(params string) error {
	args := strings.Split(params, ":")
	if invalidQuestionArgs(args) {
		return errInvalidFormat
	}
	question := &question{prompt: args[0], answer: args[1]}

	return addQuestion(question)
}

func questionInRange(selected int, questions []question) error {
	if selected > len(questions) || selected < 1 {
		return errNotFound
	}

	return nil
}

// Question methods
func (q *question) cleanup() {
	q.prompt = strings.TrimSpace(q.prompt)
	q.answer = strings.TrimSpace(q.answer)
}

func (q *question) destroy() error {
	sql := `
		DELETE FROM questions
		WHERE id = $1
	`
	_, err := db.Exec(sql, q.id)

	return err
}

func (q *question) String() string {
	return fmt.Sprintf("%s: %s", q.prompt, q.answer)
}

func (q *question) update() error {
	q.cleanup()

	sql := `
		UPDATE questions
		SET prompt = $1, answer = $2
		WHERE id = $3
	`
	_, err := db.Exec(sql, q.prompt, q.answer, q.id)

	return err
}
