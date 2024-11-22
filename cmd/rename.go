package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"keykube/internal/store"
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a secret",
	Run: func(cmd *cobra.Command, args []string) {
		secret, ok := store.Get(args[0])

		if !ok {
			fmt.Printf("The secret %v not found\n", args[0])
			return
		}

		if store.Exists(args[1]) {
			fmt.Printf("The secret %v already exists\n", args[1])
			return
		}

		data := store.Entity{
			Name: args[1],
		}

		secret.Update(data)

		fmt.Printf("The secret '%v' has been renamed to '%v'\n", args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
