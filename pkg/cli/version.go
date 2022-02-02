package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func SetupVersionCmd() *cobra.Command {

	command := &cobra.Command{
		Use:   "version",
		Short: "Versão do Go-Ware",
		Run: func(cmd *cobra.Command, as []string) {
			RunVersion()
		},
	}

	return command
}

func RunVersion() {
	fmt.Printf("v.01\n")
}
