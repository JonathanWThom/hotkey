package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
	"testing"
	"time"
)

func TestNewQuestion(t *testing.T) {
	setup()

	tests := []struct {
		input    string
		expected error
	}{
		{"prompt:answer", nil},
		{"prompt with space:answer with space", nil},
		{":missing prompt has colon", errInvalidFormat},
		{"missing answer has colon:", errInvalidFormat},
		{"one argument", errInvalidFormat},
		{"too:many:arguments", errInvalidFormat},
	}

	for _, test := range tests {
		err := NewQuestion(test.input)
		if err != test.expected {
			testError(t, test.input, err, test.expected)
		}
	}

	teardown()
}

func TestDeleteQuestion(t *testing.T) {
	setup()
	buildQuestion()

	tests := []struct {
		input    int
		expected error
	}{
		{1, nil},
		{0, errNotFound},
		{2, errNotFound},
	}

	for _, test := range tests {
		err := DeleteQuestion(test.input)
		if err != test.expected {
			t.Errorf("For input %d: got %s, expected %s", test.input, err, test.expected)
		}
	}

	teardown()
}

func TestEditQuestion(t *testing.T) {
	setup()

	for i := 0; i < 2; i++ {
		buildQuestion()
	}

	tests := []struct {
		input    string
		expected error
	}{
		{"1:new prompt:new answer", nil},
		{"new prompt:new answer", errInvalidFormat},
		{"3:new prompt:new answer", errNotFound},
	}

	for _, test := range tests {
		err := EditQuestion(test.input)
		if err != test.expected {
			testError(t, test.input, err, test.expected)
		}
	}

	teardown()
}

// Helpers
func setup() {
	var err error
	db, err = sql.Open("postgres", "dbname=hotkey_test sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func teardown() {
	query := `
		TRUNCATE TABLE questions
	`
	db.Exec(query)
	db = &sql.DB{}
}

func testError(t *testing.T, input string, err error, expected error) {
	t.Errorf("For input %s: got %s, expected %s", input, err, expected)
}

// Factories
func buildQuestion() error {
	sourc := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(sourc)
	num := rand.Intn(100)
	prompt := fmt.Sprintf("prompt-%d", num)
	answer := fmt.Sprintf("answer-%d", num)
	question := question{prompt: prompt, answer: answer}

	return addQuestion(&question)
}
