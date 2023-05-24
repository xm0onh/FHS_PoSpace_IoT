package election

import (
	"crypto/sha1"
	"encoding/binary"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/gitferry/bamboo/PoSpace"
	"github.com/gitferry/bamboo/identity"
	"github.com/gitferry/bamboo/log"
	"github.com/gitferry/bamboo/types"
)

type Rotation struct {
	peerNo int
}

func NewRotation(peerNo int) *Rotation {
	return &Rotation{
		peerNo: peerNo,
	}
}

func (r *Rotation) IsLeader(id identity.NodeID, view types.View) bool {
	if view <= 3 {
		if id.Node() < r.peerNo {
			return false
		}
		return true
	}
	h := sha1.New()
	h.Write([]byte(strconv.Itoa(int(view) + 1)))
	bs := h.Sum(nil)
	data := binary.BigEndian.Uint64(bs)
	return data%uint64(r.peerNo) == uint64(id.Node()-1)
}

func (r *Rotation) FindLeaderFor(view types.View) identity.NodeID {
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator.

	if view <= 3 {
		return identity.NewNodeID(r.peerNo)
	}
	// h := sha1.New()
	// h.Write([]byte(strconv.Itoa(int(view + 1))))
	// bs := h.Sum(nil)
	// data := binary.BigEndian.Uint64(bs)
	// // id := data%uint64(r.peerNo) + 1
	id := rand.Intn(r.peerNo)

	// PoSpace
	var wg sync.WaitGroup
	k := int(math.Pow(2, 10))
	test := PoSpace.ToBinary(int(id), k)
	key := test + PoSpace.ToBinary(rand.Intn(k), k)
	hash := PoSpace.HashSHA256(key)

	space := ""
	wg.Add(1)
	go func() {
		defer wg.Done()
		space = PoSpace.Challenge(int(id), hash)
	}()
	wg.Wait()
	if space == key {
		log.Debugf("PoSpace is Successfuly passed for node: [%v]", id)
		return identity.NewNodeID(int(id))
	}
	return ""

}
