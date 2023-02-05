/*
 *  Brown University, CS138, Spring 2019
 *
 *  Purpose: Defines IDs for tapestry and provides various utility functions
 *  for manipulating and creating them. Provides functions to compare IDs
 *  for insertion into routing tables, and for implementing the routing
 *  algorithm.
 */

package tapestry

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

// An ID is a digit array
type ID [DIGITS]Digit

// A digit is just a typedef'ed uint8
type Digit uint8

// Random number generator for generating random node ID
var random = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

// Returnes a random ID.
func RandomID() ID {
	var id ID
	for i := range id {
		id[i] = Digit(random.Intn(BASE))
	}
	return id
}

// Hashes the string to an ID
func Hash(key string) (id ID) {
	// Sha-hash the key
	sha := sha1.New()
	sha.Write([]byte(key))
	hash := sha.Sum([]byte{})

	// Store in an ID
	for i := range id {
		id[i] = Digit(hash[(i/2)%len(hash)])
		if i%2 == 0 {
			id[i] >>= 4
		}
		id[i] %= BASE
	}

	return id
}

// Returns the length of the prefix that is shared by the two IDs.
func SharedPrefixLength(a ID, b ID) (i int) {
	// TODO: students should implement this function
	count := 0
	//Debug.Printf("a: %v and b: %v ", a, b)
	for i := 0; i < DIGITS; i++ {
		if a[i] == b[i] {
			count ++
		} else {
			break
		}
	}
	//Debug.Printf("count: %v", count)
	return count
}

// Used by Tapestry's surrogate routing.  Given IDs first and second, which is the better choice?
//
// The "better choice" is the ID that:
//  - has the longest shared prefix with id
//  - if both have prefix of length n, which id has a better (n+1)th digit?
//  - if both have the same (n+1)th digit, consider (n+2)th digit, etc.

// Returns true if the first ID is the better choice.
// Returns false if second ID is closer or if first == second.
func (id ID) BetterChoice(first ID, second ID) bool {
	// TODO: students should implement this
	if first == second {
		return false
	}

	firstPrefix := SharedPrefixLength(id, first)
	secondPrefix := SharedPrefixLength(id, second)

	if firstPrefix != secondPrefix {
		return firstPrefix > secondPrefix
	}

	for i := firstPrefix; i < DIGITS; i++ {
		if first[i] != second[i] {
			a := id[i]
			for {
				a += Digit(1)
				a %= BASE

				if a == first[i] {
					return true
				} else if a == second[i] {
					return false
				}
			}
		}
	}

	return false
}

// Used when inserting nodes into Tapestry's routing table.  If the routing
// table has multiple candidate nodes for a slot, then it chooses the node that
// is closer to the local node.
//
// In a production Tapestry implementation, closeness is determined by looking
// at the round-trip-times (RTTs) between (a, id) and (b, id); the node with the
// shorter RTT is closer.
//
// In this implementation, we have decided to define closeness as the absolute
// value of the difference between a and b. This is NOT the same as your
// implementation of BetterChoice.
//
// Return true if a is closer than b.
// Return false if b is closer than a, or if a == b.
func (id ID) Closer(first ID, second ID) bool {
	// TODO: students should implement this
	//idToCompare := id.big()
	idA := first.big()
	idB := second.big()
	id_copy := id.big()
	//Debug.Printf("idA: %v", idA)
	//Debug.Printf("idB: %v", idB)
	idA.Abs(idA.Sub(id_copy,idA))
	idB.Abs(idB.Sub(id_copy,idB))
	//Debug.Printf("subt A: %v", idA)
	//Debug.Printf("subt B: %v", idB)
	if idA.Cmp(idB) < 0 {
		return true
	}
	return false
}

// Helper function: convert an ID to a big int.
func (id ID) big() (b *big.Int) {
	b = big.NewInt(0)
	base := big.NewInt(BASE)
	for _, digit := range id {
		b.Mul(b, base)
		b.Add(b, big.NewInt(int64(digit)))
	}
	return b
}

// String representation of an ID is hexstring of each digit.
func (id ID) String() string {
	var buf bytes.Buffer
	for _, d := range id {
		buf.WriteString(d.String())
	}
	return buf.String()
}

// String representation of a digit is its hex value
func (digit Digit) String() string {
	return fmt.Sprintf("%X", byte(digit))
}

func (i ID) bytes() []byte {
	b := make([]byte, len(i))
	for idx, d := range i {
		b[idx] = byte(d)
	}
	return b
}

func idFromBytes(b []byte) (i ID) {
	if len(b) < DIGITS {
		return
	}
	for idx, d := range b[:DIGITS] {
		i[idx] = Digit(d)
	}
	return
}

// Parse an ID from String
func ParseID(stringID string) (ID, error) {
	var id ID

	if len(stringID) != DIGITS {
		return id, fmt.Errorf("Cannot parse %v as ID, requires length %v, actual length %v", stringID, DIGITS, len(stringID))
	}

	for i := 0; i < DIGITS; i++ {
		d, err := strconv.ParseInt(stringID[i:i+1], 16, 0)
		if err != nil {
			return id, err
		}
		id[i] = Digit(d)
	}

	return id, nil
}
