package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"keykube/internal/password"
	"keykube/internal/store"
	"keykube/internal/utils"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new secret",
	Run: func(cmd *cobra.Command, args []string) {
		if store.Exists(args[0]) {
			fmt.Printf("The secret %v already exists\n", args[0])
			return
		}

		if params.Password == "" {
			params.Password = password.Generate(&params.Spec)
		}

		store.Add(args[0], params.Password, utils.Lifetime(params.Lifetime)).Print()
	},
}

func init() {
	initPasswordCommonFlags(newCmd)
	rootCmd.AddCommand(newCmd)
}
