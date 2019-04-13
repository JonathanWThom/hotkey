package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestTest(t *testing.T) {
	questions = []question{
		question{
			prompt: "Atom: Insert new line after current line",
			answer: "command enter",
		},
	}

	tests := []struct {
		input    string
		expected string
	}{
		{"command enter", successMsg},
		{"wrong", failureMsg},
		{"solution", "command enter"},
		{"hint", "c****** ****r"},
	}

	for _, test := range tests {
		input := fmt.Sprintf("%s\n", test.input)
		stdin = strings.NewReader(input)
		stdout := bytes.NewBufferString("")
		Test(questions[0], stdout)

		if !strings.Contains(stdout.String(), test.expected) {
			t.Errorf("For input %s: got %s, expected %s", input, stdout.String(), test.expected)
		}
	}
}
