package main

type question struct {
	prompt string
	answer string
}

var questions = []question{
	question{
		prompt: "Atom: Insert new line after current line",
		answer: "command enter",
	},
	question{
		prompt: "Atom: Insert new line before current line",
		answer: "command shift enter",
	},
	question{
		prompt: "Atom: Delete current line",
		answer: "control shift k",
	},
	question{
		prompt: "Atom: Move current line up",
		answer: "control command up",
	},
	question{
		prompt: "Atom: Move current line down",
		answer: "control command down",
	},
	question{
		prompt: "Atom: Duplicate current line",
		answer: "shift command d",
	},
	question{
		prompt: "Atom: Join current and next lines",
		answer: "command j",
	},
}
