package todo

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Error invalid index!")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	i := index - 1

	if err := t.validateIndex(i); err != nil {
		return err
	}

	t[i].Title = title

	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos
	i := index - 1

	if err := t.validateIndex(i); err != nil {
		return err
	}

	*todos = append(t[:i], t[i+1:]...)

	return nil
}

func (todos *Todos) CompleteToggle(index int) error {
	t := *todos
	i := index - 1

	if err := t.validateIndex(i); err != nil {
		return err
	}

	isComplete := t[i].Completed

	if !isComplete {
		time := time.Now()
		t[i].CompletedAt = &time
	}

	t[i].Completed = !isComplete

	return nil
}

func (todos *Todos) Print() {
	appTable := table.New(os.Stdout)

	appTable.SetRowLines(false)
	appTable.SetHeaders("#", "Task", "Completed", "CreatedAt", "CompletedAt")
	appTable.SetDividers(table.UnicodeRoundedDividers)

	for index, t := range *todos {
		complete := "❌"
		completedAt := ""

		if t.Completed {
			complete = "✅"

			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		appTable.AddRow(
			strconv.Itoa(index+1),
			t.Title,
			complete,
			t.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}

	appTable.Render()
}
