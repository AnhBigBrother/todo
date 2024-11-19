package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) validateIdx(idx int) error {
	if idx < 0 || idx >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) Add(title string) {
	newTodo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, newTodo)
}

func (todos *Todos) Delete(idx int) error {
	t := *todos

	if err := t.validateIdx(idx); err != nil {
		return err
	}
	t = append(t[:idx], t[idx+1:]...)

	return nil
}

func (todos *Todos) Toggle(idx int) error {
	t := *todos

	if err := t.validateIdx(idx); err != nil {
		return err
	}

	isCompleted := t[idx].Completed
	if !isCompleted {
		complete_time := time.Now()
		t[idx].CompletedAt = &complete_time
	}
	t[idx].Completed = !isCompleted
	return nil
}

func (todos *Todos) Edit(idx int, title string) error {
	t := *todos
	if err := t.validateIdx(idx); err != nil {
		return err
	}
	t[idx].Title = title
	return nil
}

func (todo *Todos) Print() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "title", "completed", "created at", "completed at"})
	for i, t := range *todo {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✔️"
			completedAt = t.CompletedAt.Format(time.RFC1123)
		}
		row := []string{strconv.Itoa(i), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt}
		table.Append(row)
	}
	table.Render()
}
