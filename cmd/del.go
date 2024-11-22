package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"keykube/internal/store"
	"keykube/internal/utils"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a secret",
	Run: func(cmd *cobra.Command, args []string) {
		if !store.Exists(args[0]) {
			fmt.Printf("The secret %v doesn't exist\n", args[0])
			return
		}

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Printf("Are you going to delete %v [yes/%v] ? ", args[0], utils.Bold("no"))
			text, _ := reader.ReadString('\n')
			if text == "yes\n" {
				store.Delete(args[0])
				fmt.Printf("The secret %v has been deleted\n", args[0])
				break
			} else if text == "no\n" || text == "\n" {
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
