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

var todos = data.Todos{}
var store = data.NewStorage[data.Todos]("/home/bigbro/todo_data.json")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo-list cli",
	Long:  "simple, easy to use todo-list cli",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
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
