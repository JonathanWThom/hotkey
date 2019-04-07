package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestTest(t *testing.T) {
	hotkeys = []hotkey{
		hotkey{
			description: "Atom: Insert new line after current line",
			command:     "command enter",
		},
	}

	tests := []struct {
		input    string
		expected string
	}{
		{"command enter", successMsg},
		{"wrong", failureMsg},
		{"solution", "command enter"},
	}

	for _, test := range tests {
		input := fmt.Sprintf("%s\n", test.input)
		stdin = strings.NewReader(input)
		stdout := bytes.NewBufferString("")
		Test(hotkeys[0], stdout)

		if !strings.Contains(stdout.String(), test.expected) {
			t.Errorf("For input %s: got %s, expected %s", input, stdout.String(), test.expected)
		}
	}
}
