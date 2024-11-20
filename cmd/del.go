package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a todo by it's index",
	Long:  "Delete a todo by it's index, ex: todo del <idx>",
	Run: func(cmd *cobra.Command, args []string) {
		todos.Delete(args[0])
		todos.PrintAll()
		fmt.Println("Deleted task at index", args[0])
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
