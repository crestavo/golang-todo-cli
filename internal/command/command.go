package command

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crestavo/golang-todo-cli/internal/todo"
)

type CmdFlags struct {
	Add, Edit           string
	Del, CompleteToggle int
	List                bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add new todo")
	flag.StringVar(&cf.Edit, "edit", "", "Edit todo by index. id:new_title")
	flag.IntVar(&cf.Del, "delete", 0, "Delete todo by index")
	flag.IntVar(&cf.CompleteToggle, "completeToggle", 0, "Toggle complete for task")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *todo.Todos) {
	switch {
	case cf.List:
		todos.Print()

	case cf.Add != "":
		todos.Add(cf.Add)

	case cf.Edit != "":
		string := strings.SplitN(cf.Edit, ":", 2)
		index, err := strconv.Atoi(string[0])

		if len(string) != 2 {
			fmt.Println("Error invalid format. Use id:new_title to edit task!")
			os.Exit(1)
		}

		if err != nil {
			fmt.Println("Error invalid index!")
			os.Exit(1)
		}

		todos.Edit(index, string[1])

	case cf.Del != 0:
		todos.Delete(cf.Del)

	case cf.CompleteToggle != 0:
		todos.CompleteToggle(cf.CompleteToggle)

	default:
		fmt.Println("Invalid Command")
	}
}
