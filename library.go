package main

type hotkey struct {
	description string
	command     string
}

var hotkeys = []hotkey{
	hotkey{
		description: "Atom: Insert new line after current line",
		command:     "command enter",
	},
	hotkey{
		description: "Atom: Insert new line before current line",
		command:     "command shift enter",
	},
	hotkey{
		description: "Atom: Delete current line",
		command:     "control shift k",
	},
	hotkey{
		description: "Atom: Move current line up",
		command:     "control command up",
	},
	hotkey{
		description: "Atom: Move current line down",
		command:     "control command down",
	},
	hotkey{
		description: "Atom: Duplicate current line",
		command:     "shift command d",
	},
	hotkey{
		description: "Atom: Join current and next lines",
		command:     "command j",
	},
}
