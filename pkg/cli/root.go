package cli

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type RootFlags struct {
	LogLevel string
}

var Flags *RootFlags

func RunRootCommand() {
	command := SetupRootCommand()
	if err := command.Execute(); err != nil {
		log.Errorf("Ocorreu um erro: %v", err)
		os.Exit(1)
	}
}

func SetupRootCommand() *cobra.Command {
	Flags = &RootFlags{}
	command := &cobra.Command{
		Use:   "govirt",
		Short: "govirt - bla bla",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return SetupLogs(Flags.LogLevel)
		},
	}

	command.PersistentFlags().StringVarP(&Flags.LogLevel, "verbosity", "v", log.WarnLevel.String(), "Log level: debug, info, warn, error, fatal, panic")

	command.AddCommand(SetupVirtCmd())
	command.AddCommand(SetupVersionCmd())

	return command
}

func SetupLogs(level string) error {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.Errorf("%s", err)
	}
	log.SetLevel(lvl)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if lvl == log.DebugLevel {
		log.SetReportCaller(true)
	}
	return nil
}
