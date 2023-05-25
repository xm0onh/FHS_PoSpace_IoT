package election

import (
	"github.com/xm0onh/FHS_PoSpace_IoT/identity"
	"github.com/xm0onh/FHS_PoSpace_IoT/types"
)

type Election interface {
	IsLeader(id identity.NodeID, view types.View) bool
	FindLeaderFor(view types.View) identity.NodeID
}
