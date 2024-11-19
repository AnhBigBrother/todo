package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandFlag struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCommandFlag() *CommandFlag {
	cf := CommandFlag{}
	flag.StringVar(&cf.Add, "add", "", "Add your new todo")
	flag.StringVar(&cf.Edit, "edit", "", "Edit todo by it's index, ex: id:new_title")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Mark/unmark a todo as completed")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by it's index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.Parse()
	return &cf
}

func (cf *CommandFlag) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.Split(cf.Edit, ":")
		if len(parts) != 2 {
			fmt.Println("Error: invalid format for edit, use id:new_title format")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.Edit(index, parts[1])

	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)

	case cf.Del != -1:
		todos.Delete(cf.Del)

	default:
		fmt.Println("Invalid command")
	}
}
