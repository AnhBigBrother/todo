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

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo-list cli",
	Long:  "simple todo-list cli app",
}

var todos = data.Todos{}
var homeDir, _ = os.UserHomeDir()
var store = data.NewStorage[data.Todos](fmt.Sprintf("%s/todo_data.json", homeDir))

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
		fmt.Println(err.Error())
	}
	err = store.Save(todos)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	store.Load(&todos)
}
