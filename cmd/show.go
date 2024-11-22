package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"keykube/internal/store"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all secrets according to the given pattern or the secret with the given name",
	Run: func(cmd *cobra.Command, args []string) {
		pattern := ""
		if len(args) > 0 {
			pattern = args[0]
		}

		if secret, ok := store.Get(pattern); ok {
			secret.Print()
			return
		}

		entities := []store.Entity{}
		for _, entity := range store.Data {
			if strings.HasPrefix(entity.Name, pattern) {
				entities = append(entities, entity)
			}
		}

		if len(entities) == 0 {
			fmt.Printf("No secrets found\n")
		}

		for _, entity := range entities {
			fmt.Printf("%v\n", entity.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
