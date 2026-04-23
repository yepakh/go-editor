# go-editor

A terminal text editor build from scrant using GO.

The idea behind is to learn GO and also to try solving some problems that I did not encounter.

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
