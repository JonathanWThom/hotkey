# Hotkey

A simple flashcard CLI. I originally built this to teach myself more keyboard
shortcuts (hence the name), but it can be used for anything.

### Setup

Requires Go and PostgreSQL.

1. `go get -u github.com/JonathanWThom/hotkey`
2. In Postgres, run the contents of `database.sql`.
3. Add some questions and answers: `INSERT INTO questions (prompt, answer) VALUES('Atom: Move current line down:', 'control command down');`
4. Run `./hotkey`

### Features

```
Atom: Move current line down:
-> hint
c****** ******d d***

Atom: Move current line down
-> control command down
Correct!

Vim (NERDTree): Toggle tree
-> solution
control n
```

### Todo

* It would be nice to read keystrokes, instead of entering the names of the commands. [termbox-go](https://github.com/nsf/termbox-go) might be what I need to use.
* Add/edit/delete records.
* Keep score.
* Dockerize?
* Create a web ui?

### License

MIT