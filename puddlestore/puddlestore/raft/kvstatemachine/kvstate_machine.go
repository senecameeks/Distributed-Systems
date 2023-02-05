package kvstatemachine

import (
	// "crypto/md5"
	"errors"
	"fmt"
	"strings"
)

const (
	KV_STATE_INIT uint64 = iota
	KV_STATE_ADD
	KV_STATE_GET
	KV_STATE_REMOVE
)

// // StateMachine is a general interface defining the methods a Raft state machine
// // should implement. For this project, we use a Key-Value StateMachine (KVStateMachine)
// as our state machine.
// type StateMachine interface {
// 	GetState() (state interface{}) // Useful for testing once you use type assertions to convert the state
// 	ApplyCommand(command uint64, data []byte) (message string, err error)
// 	FormatCommand(command uint64) (commandString string)
// }

// KVStateMachine implements the raft.StateMachine interface, and represents a
// finite state machine storing a hash value. It stores a map which takes aguids to vguids

type KVStateMachine struct {
	// hash []byte

	kvstore map[string]string
}

func (h *KVStateMachine) init() (hash string, err error) {
	if h.kvstore != nil {
		return fmt.Sprintf("%v", h.kvstore), nil
		//return "", errors.New("KVStore: the kvstore should only be initialized once")
	}

	h.kvstore = make(map[string]string)

	return fmt.Sprintf("%v", h.kvstore), nil
}

// TODO: initialize from log capability (init(mapping)

func (h *KVStateMachine) add(key string, val string) (err error) {
	if h.kvstore == nil {
		return errors.New("KVStore: the key-value store hasn't been initialized yet")
	}
	// check if a value exists for this key
	//if oldval, exists := h.kvstore[key]; exists {
	//fmt.Printf("KVStore: val for %v changing from %v to %v\n", key, oldval, val)
	//}

	// store new value
	h.kvstore[key] = val

	//fmt.Printf("KVStore: val for %v is %v\n", key, h.kvstore[key])
	return nil
}

func (h *KVStateMachine) get(key string) (val string, err error) {
	if h.kvstore == nil {
		return "", errors.New("KVStore: the key-value store hasn't been initialized yet")
	}

	// if doesnt exist, return empty string
	if val, exists := h.kvstore[key]; exists {
		return val, nil
	} else {
		// fmt.Printf("KVStore:")
		return "", nil
	}
}

// removes a value from the KV mapping
func (h *KVStateMachine) remove(key string) (err error) {
	if h.kvstore == nil {
		return errors.New("KVStore: the key-value store hasn't been initialized yet")
	}
	delete(h.kvstore, key)
	return nil
}

// GetState returns the state of the state machine as an interface{}, which can
// be converted to the expected type using type assertions.
func (h *KVStateMachine) GetState() (state interface{}) {
	return h.kvstore
}

// ApplyCommand applies the given state machine command to the HashMachine, and
// returns a message and an error if there is one.
func (h *KVStateMachine) ApplyCommand(command uint64, data []byte) (message string, err error) {
	switch command {
	case KV_STATE_INIT:
		return h.init()
	case KV_STATE_ADD:
		// separate key and val with a period bc only one data argument to ApplyCommand
		keyval := strings.Split(string(data), ".")
		// key = aguid, val = vguid
		key := keyval[0]
		val := keyval[1]
		return "", h.add(key, val)
	case KV_STATE_GET:
		key := string(data)
		return h.get(key)
	case KV_STATE_REMOVE:
		key := string(data)
		return "", h.remove(key)
	default:
		return "", errors.New("unknown command type")
	}
}

// FormatCommand returns the string representation of a HashMachine state machine command.
func (h *KVStateMachine) FormatCommand(command uint64) (commandString string) {
	switch command {
	case KV_STATE_INIT:
		return "KV_STATE_INIT"
	case KV_STATE_ADD:
		return "KV_STATE_ADD"
	case KV_STATE_GET:
		return "KV_STATE_GET"
	case KV_STATE_REMOVE:
		return "KV_STATE_REMOVE"
	default:
		return "UNKNOWN_COMMAND"
	}
}

func (h KVStateMachine) String() string {
	return fmt.Sprintf("KVStateMachine{%v}", h.kvstore)
}
