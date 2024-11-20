package cmd

import (
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a todo as done",
	Long:  "Mark a todo as done by it's index, ex: todo done <idx>",
	Run: func(cmd *cobra.Command, args []string) {
		for _, idxStr := range args {
			todos.Done(idxStr)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
