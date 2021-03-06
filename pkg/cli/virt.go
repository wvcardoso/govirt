package cli

import (
	"fmt"

	vc "github.com/wvcardoso/govirt/pkg/vcenter"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type VirtArgs struct {

	// Confs Virtualizacao
	IP, User, Password string
	SkipInsecure       bool
}

func SetupVirtCmd() *cobra.Command {

	args := &VirtArgs{}

	command := &cobra.Command{
		Use:           "virt",
		Short:         "virt - bla bla",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, as []string) error {
			return RunVirt(args)
		},
	}

	command.Flags().StringVar(&args.IP, "ip", "", "vcenter endpoint")
	command.Flags().StringVar(&args.User, "user", "", "vcenter user")
	command.Flags().StringVar(&args.Password, "pass", "", "vcenter password")
	command.Flags().BoolVar(&args.SkipInsecure, "skip-insecure", true, "true to ignore")

	command.MarkPersistentFlagRequired("ip")
	command.MarkPersistentFlagRequired("user")
	command.MarkPersistentFlagRequired("pass")

	return command
}

func RunVirt(args *VirtArgs) error {

	log.Debugf("Iniciando goware")

	vCenterURL := fmt.Sprintf("https://%s:443/sdk", args.IP)

	con, err := vc.Conn(vCenterURL, args.User, args.Password, args.SkipInsecure)
	if err != nil {
		log.Errorf("Ocorreu um erro: %v", err)
	}

	vmName := "vm0001"
	con.GetAllVMs(vmName)

	return nil
}
