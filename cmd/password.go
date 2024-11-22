package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"keykube/internal/password"
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate a password",
	Run: func(cmd *cobra.Command, args []string) {
		pass := password.Generate(&params.Spec)
		fmt.Println(pass, password.Entropy(pass))
	},
}

func init() {
	initPasswordCommonFlags(passwordCmd)
	rootCmd.AddCommand(passwordCmd)
}
