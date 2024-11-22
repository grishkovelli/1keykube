package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"keykube/internal/password"
)

type SecretParams struct {
	Password string
	Lifetime int
	Spec     password.Spec
}

var params = SecretParams{}

var examples = []string{
	"  keykube new email/yourname@gmail.com --length 20 --lifetime 90",
	"  keykube show email",
	"  keykube rename my-name@mail.com email/my-name@mail.com",
}

var rootCmd = &cobra.Command{
	Use:     "keykube",
	Short:   "keykube 1.0.0 ( https://github.com/grishkovelli/keykube )",
	Example: strings.Join(examples, "\n"),
}

func initPasswordCommonFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&params.Password, "password", "p", "", "Secret password")
	cmd.Flags().IntVarP(&params.Lifetime, "time", "t", 0, "Password lifetime (days)")

	cmd.Flags().IntVarP(&params.Spec.Length, "length", "l", 16, "Password length")
	cmd.Flags().IntVarP(&params.Spec.Uppercase, "uppercase", "u", 4, "Number of uppercase letters")
	cmd.Flags().IntVarP(&params.Spec.Numbers, "numbers", "n", 4, "Number of numbers")
	cmd.Flags().IntVarP(&params.Spec.Symbols, "symbols", "s", 4, "Number of symbols")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
