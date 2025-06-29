# terminal.todo — A Beginner's TODO App in Go

A simple terminal-based TODO list manager built in Go.  
This project was created as a fun way to explore **Go structs**, **file operations**, and **basic CLI interaction**.

---

## Features

- Add TODO items
- Mark TODOs as complete/incomplete
- Delete TODOs
- Automatically saves your list to `~/.TODO.txt`
- Uses basic file read/write and SHA-256 hashing to detect changes

---

## What I’m Trying to Learning

This project is part of my journey to learn Go.  
Specifically, I wanted to get comfortable with:

- Defining and using `struct`s
- Writing and reading from files
- Using slices to manage dynamic lists
- Building a terminal menu interface
- Using Go’s `bufio`, `os`, `fmt`, and `crypto/sha256` packages

---

## How to Run

### 1. Clone the repository

```bash
git clone https://github.com/mindslayer001/terminal.todo.git
cd terminal.todo
```

### 2. Run the app

```bash
go run main.go
```

---

## File Structure

```
terminal.todo/
├── TODO/             # Package containing the TODO logic
│   └── todo.go
├── main.go           # Entry point with CLI logic
├── go.mod
└── README.md         # You are here :)
```

---

## Data File

All TODOs are saved in your home directory as:

```
~/.TODO.txt
```

Each item is stored in this format:

```
Title|true
Title 2|false
```

---

## Sample Usage

```
*************************TODO*************************
1 [ ] Learn Go
2 [✅] Build a CLI
Choose from below options
1 Create New todo
2 Change the Status of a certain todo
3 Delete the todo
4 Exit
```

---

## Credits

This is a personal project built by a beginner learning Go.
If you're also learning and want to collaborate or ask questions, feel free to reach out!

---

## Goals for the Future

* [ ] Save TODOs in JSON format
* [ ] Add due dates to tasks
* [ ] Support editing existing TODOs
* [ ] Improve UI with colors using a popular library (Still Not Selected)

---

## Inspiration

Learning by building — nothing fancy, just getting hands-on with Go's core features!

---

## Go Version

Tested with: `go1.22+`
