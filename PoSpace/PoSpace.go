package PoSpace

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/gitferry/bamboo/identity"
)

type PoSpace struct {
	id   identity.NodeID
	Test chan int
}

func ToBinary(n int, digits int) string {
	// return strconv.FormatInt(int64(n), 2)
	return fmt.Sprintf("%0"+strconv.Itoa(digits)+"b", n)
}

func HashSHA256(s string) string {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

// func (P *PoSpace) New(id identity.NodeID) {
// 	K := int(math.Pow(2, 2))
// 	db, err := leveldb.OpenFile(strconv.Itoa(id.Node()), nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()
// 	nodeIDBinary := ToBinary((id.Node()), K)
// 	for i := 0; i < K; i++ {
// 		key := nodeIDBinary + ToBinary(i, K)
// 		hash := HashSHA256(key)
// 		err = db.Put([]byte(key), []byte(hash), nil)
// 		fmt.Println("---", hash, "---", key, "---")
// 		P.Test <- i
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// 	fmt.Println("-----")
// 	// data, _ := db.Get([]byte("14ec116a770ea76d699c7467b2c3082725d369a7813493fc25b60d2695ec30c5"), nil)
// 	// fmt.Println("---->", string(data))

// }
func (P *PoSpace) New(id identity.NodeID) {
	K := int(math.Pow(2, 10))
	nodeIDBinary := ToBinary((id.Node()), K)
	for i := 0; i < K; i++ {
		key := nodeIDBinary + ToBinary(i, K)
		hash := HashSHA256(key)

		// open the file using the append flag
		file, err := os.OpenFile(fmt.Sprintf("%d.txt", id.Node()), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("error opening file:", err)
			return
		}

		// write the key-hash pair to the file
		if _, err := file.WriteString(fmt.Sprintf("%s---%s\n", hash, key)); err != nil {
			fmt.Println("error writing to file:", err)
			file.Close()
			return
		}

		// close the file after writing
		if err := file.Close(); err != nil {
			fmt.Println("error closing file:", err)
			return
		}

		P.Test <- i
	}
	fmt.Println("-----")
}
