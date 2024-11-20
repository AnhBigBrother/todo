package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a todo by it's index",
	Long:  "Edit a todo by it's index, ex: todo edit <id>:<new_title>",
	Run: func(cmd *cobra.Command, args []string) {
		for _, str := range args {
			parts := strings.Split(str, ":")
			if len(parts) != 2 {
				fmt.Println("Error: invalid format for edit, use <id>:<new_title> format")
				os.Exit(1)
			}

			todos.Edit(parts[0], parts[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
