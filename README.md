# Hotkey

A simple flashcard CLI. I originally built this to teach myself more keyboard
shortcuts (hence the name), but it can be used for anything.

### Setup

Requires Go and PostgreSQL.

1. `go get -u github.com/JonathanWThom/hotkey`
2. In Postgres, run the contents of `database.sql`.
3. Add questions and answers via CLI (see [Features](#Features)).
4. Run `./hotkey`

### Features

#### Adding, Listing, & Updating Questions

Add
```
./hotkey -add "my prompt:my answer"
Question successfully added.
./hotkey -add "another prompt:another answer"
```

List
```
./hotkey -all
1. my prompt: my answer
2. another prompt: another answer
```

Update
```
./hotkey -edit "2:modified prompt:modified answer"
Question successfully updated.
./hotkey -all
1. my prompt: my answer
2. modified prompt: modified answer
```

Delete
```
./hotkey -delete=2
Question successfully deleted.
./hotkey -all
1. my prompt: my answer
```

#### Answering Questions

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

### LICENSE

MIT
