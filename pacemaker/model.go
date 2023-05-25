package pacemaker

import (
	"github.com/xm0onh/FHS_PoSpace_IoT/blockchain"
	"github.com/xm0onh/FHS_PoSpace_IoT/crypto"
	"github.com/xm0onh/FHS_PoSpace_IoT/identity"
	"github.com/xm0onh/FHS_PoSpace_IoT/types"
)

type TMO struct {
	View   types.View
	NodeID identity.NodeID
	HighQC *blockchain.QC
}

type TC struct {
	types.View
	crypto.AggSig
	crypto.Signature
}

func NewTC(view types.View, requesters map[identity.NodeID]*TMO) *TC {
	// TODO: add crypto
	return &TC{View: view}
}
