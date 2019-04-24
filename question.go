package main

import (
	"errors"
	_ "github.com/lib/pq"
	"strings"
)

var invalidFormat = errors.New("Invalid question format.")

type question struct {
	prompt string
	answer string
}

func invalidQuestionArgs(args []string) bool {
	return len(args) != 2 || len(string(args[0])) == 0 || len(string(args[1])) == 0
}

func NewQuestion(params string) error {
	args := strings.Split(params, ":")
	if invalidQuestionArgs(args) {
		return invalidFormat
	}
	question := &question{args[0], args[1]}

	return addQuestion(question)
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
		SELECT prompt, answer 
		FROM questions 
	`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var questions = []question{}
	for rows.Next() {
		question := question{}
		err := rows.Scan(&question.prompt, &question.answer)
		if err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	return questions, nil
}
