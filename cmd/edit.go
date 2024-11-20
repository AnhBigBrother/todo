package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit todos by their index",
	Long:  "Edit todos by their index (allow multiple), ex: todo edit <id1>:<new_title1> <id2>:<new_title2>",
	Run: func(cmd *cobra.Command, args []string) {
		for _, str := range args {
			parts := strings.Split(str, ":")
			if len(parts) != 2 {
				fmt.Println("Error: invalid format for edit, use <id>:<new_title> format")
				os.Exit(1)
			}

			todos.Edit(parts[0], parts[1])
		}
		todos.PrintAll()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
