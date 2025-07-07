# 📝 Terminal TODO App (Built with Go)

A simple command-line TODO list manager written in Go.  
Features persistent storage, task completion toggling, deletion, and SHA-256 based change detection.

> **Built completely from scratch — no external libraries, no AI assistance.**

---

## 📦 Features

- ✅ Add, complete, and delete TODOs
- 💾 Persists tasks in a hidden file (`~/.TODO.txt`)
- 🔐 Uses SHA-256 hashing to detect changes before saving
- 🧹 Clean terminal UI
- 🏗️ Modular Go package structure (`todo` package)

---

## 🚀 Installation & Run

### Clone the Repo
```bash
git clone https://github.com/yourusername/terminal.todo.git
cd terminal.todo
```

### Initialize Go Module (if not already)

```bash
go mod tidy
```
### Build & Run

```bash
go run main.go
```

---

## 🖥️ Screenshot

> **Will be coming soon**

## 📁 File Format (`~/.TODO.txt`)

Each TODO item is saved in a `|`-separated format:

```bash
Learn Go concurrency|false
Finish README|true
Refactor the write logic|false
```

* First column: TODO title
* Second column: Completion status (`true` or `false`)

---

## 🔐 How Hashing Works

Before saving, the app checks if the in-memory TODO list has changed by hashing it using SHA-256:

```go
data, _ := json.Marshal(T.TODOS)
hash := sha256.Sum256(data)
```

If the hash differs from the previously loaded state, it writes the updated list to the file.
This prevents unnecessary writes and provides a simple "dirty state" check.

---

## 🛠️ TODO (Future Features)

* [ ] Add CLI flags for quick actions (`todo add "Task"`)
* [ ] Export to JSON/CSV
* [ ] Color-coded output using `fatih/color`
* [ ] Cobra integration for subcommands

---

## 🤝 Contribution

This was built as a personal project, but feel free to fork and extend it.
Suggestions and issues are welcome!

---

## 📍 Author

**Mani Sankar Chintagunti**
Backend Developer | OSS Contributor
Portfolio: [@mindslayer001.in](https://mindslayer001.in)
GitHub: [@mindslayer001](https://github.com/mindslayer001)
X: [@ismycodeshit](https://x.com/ismycodeshit)

---

## 📄 License

MIT License – see file for details. 
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
 
