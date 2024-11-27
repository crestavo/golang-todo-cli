package main

import (
	"github.com/crestavo/golang-todo-cli/internal/command"
	"github.com/crestavo/golang-todo-cli/internal/todo"
)

// var Todos todo.Todos

func main() {
	todos := todo.Todos{}
	storage := todo.NewStorage[todo.Todos]("data/todo.json")
	storage.Load(&todos)

	CmdFlags := command.NewCmdFlags()
	CmdFlags.Execute(&todos)

	storage.Save(todos)
}
