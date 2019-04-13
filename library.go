package main

import (
	_ "github.com/lib/pq"
)

type question struct {
	prompt string
	answer string
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
