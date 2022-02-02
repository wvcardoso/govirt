package vcenter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/session/cache"
	"github.com/vmware/govmomi/view"
	vmwarecli "github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

// NewVirtualizacaoClient implementa um client de virtualizacao a ser utilizado
func Conn(VC_URL, User, Password string, insecure bool) (*Config, error) {

	if VC_URL == "" {
		return &Config{}, fmt.Errorf("configurar url para acesso ao vcenter")
	}

	urlConnect, err := url.Parse(VC_URL)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	urlConnect.User = url.UserPassword(User, Password)

	sessionCache := &cache.Session{
		URL:      urlConnect,
		Insecure: insecure,
	}

	ctx := context.Background()
	client := new(vmwarecli.Client)
	err = sessionCache.Login(ctx, client, nil)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var VMWareConfig Config
	VMWareConfig.VMClient = client
	VMWareConfig.Session = sessionCache
	VMWareConfig.Context = ctx

	return &VMWareConfig, nil
}

func (conn *Config) GetAllVMs(vmName string) {

	manager := view.NewManager(conn.VMClient)
	kind := []string{"VirtualMachine"}

	viewer, err := manager.CreateContainerView(conn.Context, conn.VMClient.ServiceContent.RootFolder, kind, true)
	if err != nil {
		fmt.Println(err.Error())
	}

	var vms []mo.VirtualMachine

	err = viewer.RetrieveWithFilter(conn.Context, kind, []string{"name"}, &vms, property.Filter{"name": vmName})
	if err != nil {
		fmt.Println(err.Error())
	}

	//todo: verificando se há apenas 1 VMs
	for _, vm := range vms {
		fmt.Println(vm.Name)
		//fmt.Println(vm.Network)
	}

	defer viewer.Destroy(conn.Context)

}
