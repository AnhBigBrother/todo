package cmd

import (
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a todo by it's index",
	Long:  "Delete a todo by it's index, ex: todo del <idx>",
	Run: func(cmd *cobra.Command, args []string) {
		for _, idx := range args {
			todos.Delete(idx)
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
