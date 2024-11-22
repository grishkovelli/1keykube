package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"keykube/internal/store"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy the password to the clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		if clipboard.Unsupported {
			fmt.Println("Clipboard not supported")
			return
		}

		entity, ok := store.Get(args[0])
		if !ok {
			fmt.Printf("Secret %v not found\n", args[0])
			return
		}
		clipboard.WriteAll(entity.Password())

		fmt.Println("Password copied to the clipboard")
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
