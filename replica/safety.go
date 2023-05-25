package replica

import (
	"github.com/xm0onh/FHS_PoSpace_IoT/blockchain"
	"github.com/xm0onh/FHS_PoSpace_IoT/message"
	"github.com/xm0onh/FHS_PoSpace_IoT/pacemaker"
	"github.com/xm0onh/FHS_PoSpace_IoT/types"
)

type Safety interface {
	ProcessBlock(block *blockchain.Block) error
	ProcessVote(vote *blockchain.Vote)
	ProcessRemoteTmo(tmo *pacemaker.TMO)
	ProcessLocalTmo(view types.View)
	MakeProposal(view types.View, payload []*message.Transaction) *blockchain.Block
	GetChainStatus() string
}
