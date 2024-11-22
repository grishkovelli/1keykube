package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"keykube/internal/password"
	"keykube/internal/store"
	"keykube/internal/utils"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh a secret's password",
	Run: func(cmd *cobra.Command, args []string) {
		entity, ok := store.Get(args[0])

		if !ok {
			fmt.Printf("The secret %v not found\n", args[0])
			return
		}

		data := store.Entity{
			Vault: store.Vault{
				Password:    password.Generate(&params.Spec),
				OldPassword: entity.Password(),
			},
			ExpireAt: utils.Lifetime(params.Lifetime),
		}

		entity.Update(data).Print()
	},
}

func init() {
	initPasswordCommonFlags(refreshCmd)
	rootCmd.AddCommand(refreshCmd)
}
