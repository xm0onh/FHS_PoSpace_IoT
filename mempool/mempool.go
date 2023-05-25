package mempool

import (
	"time"

	"github.com/xm0onh/FHS_PoSpace_IoT/config"
	"github.com/xm0onh/FHS_PoSpace_IoT/message"
)

type MemPool struct {
	*Backend
}

// NewTransactions creates a new memory pool for transactions.
func NewMemPool() *MemPool {
	mp := &MemPool{
		Backend: NewBackend(config.GetConfig().MemSize),
	}

	return mp
}

func (mp *MemPool) addNew(tx *message.Transaction) {
	tx.Timestamp = time.Now()
	mp.Backend.insertBack(tx)
	mp.Backend.addToBloom(tx.ID)
}

func (mp *MemPool) addOld(tx *message.Transaction) {
	mp.Backend.insertFront(tx)
}
