package PoSpace

// import (
// 	"fmt"

// 	"github.com/syndtr/goleveldb/leveldb"
// )

// func Challenge(id int, hash string) string {

// 	db, err := leveldb.OpenFile("1.db", nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()
// 	data, err := db.Get([]byte(hash), nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return string(data)
// }

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Challenge(id int, hash string) string {
	key := findKey(hash, fmt.Sprintf("%d.txt", id))
	return key
}

// Function to find the key corresponding to the input hash in a file.
func findKey(hash string, filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "---")
		if parts[0] == hash {
			// fmt.Println("Found key:", parts[1])
			return parts[1]
		}
	}

	if scanner.Err() != nil {
		return ""
	}

	// Return an error if the hash was not found in the file.
	return ""
}
