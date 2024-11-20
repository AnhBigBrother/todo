package cmd

import (
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark todos as done",
	Long:  "Mark todos as done by their index (allow multiple), ex: todo done <idx1> <idx2>",
	Run: func(cmd *cobra.Command, args []string) {
		for _, idxStr := range args {
			todos.Done(idxStr)
		}
		todos.PrintAll()
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
