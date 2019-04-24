package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"testing"
)

func TestNewQuestion(t *testing.T) {
	setup()

	tests := []struct {
		input    string
		expected error
	}{
		{"prompt:answer", nil},
		{"prompt with space:answer with space", nil},
		{":missing prompt has colon", invalidFormat},
		{"missing answer has colon:", invalidFormat},
		{"one argument", invalidFormat},
		{"too:many:arguments", invalidFormat},
	}

	for _, test := range tests {
		err := NewQuestion(test.input)
		if err != test.expected {
			t.Errorf("For input %s: got %s, expected %s", test.input, err, test.expected)
		}
	}

	teardown()
}

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
