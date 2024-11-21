package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"keykube/pkg/password"
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate a password",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(password.Generate(&params.Spec))
	},
}

func init() {
	initPasswordCommonFlags(passwordCmd)
	rootCmd.AddCommand(passwordCmd)
}
