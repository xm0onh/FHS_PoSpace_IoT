package mempool

import (
	"github.com/xm0onh/FHS_PoSpace_IoT/config"
	"github.com/xm0onh/FHS_PoSpace_IoT/message"
)

type Producer struct {
	mempool *MemPool
}

func NewProducer() *Producer {
	return &Producer{
		mempool: NewMemPool(),
	}
}

func (pd *Producer) GeneratePayload() []*message.Transaction {
	return pd.mempool.some(config.Configuration.BSize)
}

func (pd *Producer) AddTxn(txn *message.Transaction) {
	pd.mempool.addNew(txn)
}

func (pd *Producer) CollectTxn(txn *message.Transaction) {
	pd.mempool.addOld(txn)
}

func (pd *Producer) TotalReceivedTxNo() int64 {
	return pd.mempool.totalReceived
}
