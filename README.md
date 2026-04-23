# go-editor

A terminal text editor build from scratch using GO.

The idea behind is to learn GO and have some fun. So clanker assist is to minimal, no generated code, mostly just Q/A with Claude.

## Stack

- **Go** — no frameworks, standard library where possible
- **[tcell](https://github.com/gdamore/tcell)** — terminal rendering

## Usage

```bash
go run . <filepath>
```

## TODO

- [ ] Status bar — show open file, cursor position, and basic keybind hints
- [ ] Edit mode — insert and delete characters, rethink the buffer data structure
