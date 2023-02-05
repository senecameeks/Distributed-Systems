/*
 *  Brown University, CS138, Spring 2018
 *
 *  Purpose: contains the hash function to be utilized by LiteMiner miners.
 */

package liteminer

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// Concatenates msg with nonce and generates a hash value. Only miners should
// ever need to call this method.
func Hash(msg string, nonce uint64) uint64 {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s %d", msg, nonce)))
	return binary.BigEndian.Uint64(hasher.Sum(nil))
}
