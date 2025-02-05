/*
Copyright © 2024 big_bro <anh.bigbrother@gmail.com>
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/AnhBigBrother/todo/data"
)

var (
	todos   data.Todos
	store   *data.Storage[data.Todos]
	rootCmd *cobra.Command
)

func Execute() {
	todos = data.Todos{}

	homeDir, _ := os.UserHomeDir()
	store = data.NewStorage[data.Todos](fmt.Sprintf("%s/todo_data.json", homeDir))

	rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "todo-list cli",
		Long:  "simple todo-list cli app",
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	err = store.Save(todos)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	store.Load(&todos)
}
