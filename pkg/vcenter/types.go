package vcenter

import (
	"context"

	"github.com/vmware/govmomi/session/cache"
	vmwarecli "github.com/vmware/govmomi/vim25"
)

type Config struct {
	VMClient *vmwarecli.Client
	Session  *cache.Session
	Context  context.Context
}
