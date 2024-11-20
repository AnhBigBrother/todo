package data

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

func (todos *Todos) validateIdx(idxStr string) (int, error) {
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	if idx < 0 || idx >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return -1, err
	}
	return idx, nil
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

func (todos *Todos) Delete(idxStr string) error {
	idx, err := (*todos).validateIdx(idxStr)

	if err != nil {
		return err
	}

	*todos = append((*todos)[:idx], (*todos)[idx+1:]...)

	return nil
}

func (todos *Todos) Done(idxStr string) error {
	t := *todos

	idx, err := t.validateIdx(idxStr)

	if err != nil {
		return err
	}
	complete_time := time.Now()
	t[idx].CompletedAt = &complete_time
	t[idx].Completed = true
	return nil
}

func (todos *Todos) Undone(idxStr string) error {
	t := *todos

	idx, err := t.validateIdx(idxStr)

	if err != nil {
		return err
	}
	t[idx].CompletedAt = nil
	t[idx].Completed = false
	return nil
}

func (todos *Todos) Edit(idxStr string, title string) error {
	t := *todos
	idx, err := t.validateIdx(idxStr)

	if err != nil {
		return err
	}

	t[idx].Title = title
	return nil
}

func (todo *Todos) PrintAll() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "title", "completed", "created at", "completed at"})
	count := 0
	for i, t := range *todo {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✔️"
			completedAt = t.CompletedAt.Format(time.RFC1123)
		}
		row := []string{strconv.Itoa(i), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt}
		table.Append(row)
		count++
	}
	table.SetFooter([]string{"", "", "", "Count", strconv.Itoa(count)})
	table.SetBorder(false)
	table.Render()
}

func (todo *Todos) PrintCompleted() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "title", "completed", "created at", "completed at"})
	count := 0
	for i, t := range *todo {
		if t.Completed {
			completed := "✔️"
			completedAt := t.CompletedAt.Format(time.RFC1123)
			row := []string{strconv.Itoa(i), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt}
			table.Append(row)
			count++
		}
	}
	table.SetFooter([]string{"", "", "", "Count", strconv.Itoa(count)})
	table.SetBorder(false)
	table.Render()
}

func (todo *Todos) PrintUnCompleted() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "title", "completed", "created at", "completed at"})
	count := 0
	for i, t := range *todo {
		if !t.Completed {
			completed := "❌"
			row := []string{strconv.Itoa(i), t.Title, completed, t.CreatedAt.Format(time.RFC1123), ""}
			table.Append(row)
			count++
		}
	}
	table.SetFooter([]string{"", "", "", "Count", strconv.Itoa(count)})
	table.SetBorder(false)
	table.Render()
}
