package Todo

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	Path string
	old  TODOList
)

func init() {
	homeDir, _ := os.UserHomeDir()
	Path = filepath.Join(homeDir, ".TODO.txt")
	old.LoadTODO()
}

type TODO struct {
	TODOTitle   string
	IsCompleted bool
}

type TODOList struct {
	TODOS []TODO
}

func (T *TODOList) PrintTODO() {
	var Status string
	fmt.Println("*************************TODO*************************")
	for i := 0; i < len(T.TODOS); i++ {
		if T.TODOS[i].IsCompleted {
			Status = "[✅]"
		} else {
			Status = "[ ]"
		}
		fmt.Printf("%d %s %s %v\n", i+1, Status, T.TODOS[i].TODOTitle, T.TODOS[i].IsCompleted)
	}
}

func (T *TODOList) CreateTODO(Tile string, status bool) {
	todo := TODO{TODOTitle: Tile, IsCompleted: status}
	T.TODOS = append(T.TODOS, todo)
}

func (T *TODOList) LoadTODO() {
	data, err := os.ReadFile(Path)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) != 3 {
			continue
		}

		T.CreateTODO(parts[0], parts[1] == "true")
	}
}

func (T *TODOList) writeTODO() {
	var content strings.Builder
	for _, todo := range T.TODOS {
		var Status string = "false"
		if todo.IsCompleted {
			Status = "true"
		}
		content.WriteString(fmt.Sprintf("\n%s|%s|", todo.TODOTitle, Status))
	}
	err := os.WriteFile(Path, []byte(content.String()), 0o644)
	if err != nil {
		log.Fatal(err)
	}
}

func (T *TODOList) DeleteTodo(index int) {
	T.TODOS = append(T.TODOS[:index-1], T.TODOS[index:]...)
}

func (T *TODOList) CheckTODONodes() {
	if T.Hash() == old.Hash() {
		return
	}
	T.writeTODO()
	return
}

//	func (T *TODOList) Hash() string {
//		data, err := json.Marshal(T.TODOS)
//		if err != nil {
//			log.Fatal(err)
//		}
//		hashByte := sha256.Sum256(data)
//		hashText := hex.EncodeToString(hashByte[:])
//		return hashText
//	}
func (T *TODOList) Hash() [32]byte {
	data, err := json.Marshal(T.TODOS)
	if err != nil {
		log.Fatal(err)
	}
	hashByte := sha256.Sum256(data)
	return hashByte
}

func (T *TODOList) Complete(index int) {
	T.TODOS[(index - 1)].IsCompleted = !T.TODOS[(index - 1)].IsCompleted
}
