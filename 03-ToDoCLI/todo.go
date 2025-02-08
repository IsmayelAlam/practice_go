package main

import (
	"errors"
	"os"
	"os/exec"
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

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Completed

	if !isCompleted {
		completedTime := time.Now()
		t[index].CompletedAt = &completedTime
	}
	t[index].Completed = true

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	clearConsole()
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""
		if t.Completed {
			completed = "✅"
			completedAt = t.CompletedAt.Format(time.ANSIC)
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.ANSIC), completedAt)
	}
	table.Render()
}

func clearConsole() {
	// fmt.Print("\033[H\033[2J")
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
