package kvstatemachine

import (
	// "fmt"
	// "reflect"
	"testing"
)

func TestInitialize(t *testing.T) {
	kv := new(KVStateMachine)

	// Test initializing kv store
	_, err := kv.init()

	if err != nil {
		t.Error(err)
	}

	// Test reinitializing kv store
	_, err = kv.init()
	if err == nil {
		t.Error("Expected error when reinitializing hash machine, did not get any")
	}
}

// func TestInitializeByApplyCommand(t *testing.T) {
// 	kv := new(KVStateMachine)
// 	// _, err := kv.init()
//
// 	// _, err := kv.ApplyCommand(KV_STATE_INIT, []int{})
//
// }

func TestAddAndGet(t *testing.T) {

	kv := new(KVStateMachine)
	_, err := kv.init()
	if err != nil {
		t.Error(err)
	}

	// test add one value
	err = kv.add("hey", "what's up")

	if err != nil {
		t.Error(err)
	}

	// test get the first value
	val, err := kv.get("hey")
	if err != nil {
		t.Error(err)
	}
	if val != "what's up" {
		t.Errorf("Expected %v, got %v", "what's up", val)
	}

	// replace the value
	err = kv.add("hey", "yo!")
	if err != nil {
		t.Error(err)
	}
	val, _ = kv.get("hey")
	if val == "what's up" {
		t.Error("value did not update on second add()")
	}
	if val != "yo!" {
		t.Errorf("Expected %v for new value, but got %v", "yo!", val)
	}
}
func TestRemove(t *testing.T) {
	kv := new(KVStateMachine)
	kv.init()
	kv.add("avicii", "a legend")

	// check that it exists in store first
	val, _ := kv.get("avicii")
	if val != "a legend" {
		t.Errorf("Expected %v, got %v", "a legend", val)
	}

	// now remove that value
	kv.remove("avicii")
	val, _ = kv.get("avicii")
	if val != "" {
		t.Errorf("Expected %v, got %v", "", val)
	}
}

func TestGetNoKey(t *testing.T) {
	// test getting a value for a key that doesn't exist in the kv-store
	kv := new(KVStateMachine)
	kv.init()
	val, _ := kv.get("nonexistent")
	if val != "" {
		t.Error("Should have received an empty string for a nonexistent key")
	}
}

func TestGetState(t *testing.T) {
	kv := new(KVStateMachine)
	kv.init()
	kv.GetState()
}

func TestString(t *testing.T) {
	kv := new(KVStateMachine)
	kv.init()
	// put some stuff in the kv store
	kv.add("above", "beyond")
	kv.add("1901", "phoenix")
	kv_string := kv.String()
	if kv_string == "" {
		t.Errorf("Got unexpected state %v for string", kv_string)
	}
}

func TestAddUninitialized(t *testing.T) {
	kv := new(KVStateMachine)
	// don't initialize but call add:
	err := kv.add("above", "beyond")
	if err == nil {
		t.Error("Should have received an error when calling add without init")
	}
}

func TestGetUninitialized(t *testing.T) {
	kv := new(KVStateMachine)
	// don't initialize but call add:
	_, err := kv.get("above")
	if err == nil {
		t.Error("Should have received an error when calling get without init")
	}
}

func TestRemoveUninitialized(t *testing.T) {
	kv := new(KVStateMachine)
	// don't initialize but call add:
	err := kv.remove("above")
	if err == nil {
		t.Error("Should have received an error when calling remove without init")
	}
}

func TestFormatCommand(t *testing.T) {
	kv := new(KVStateMachine)

	if str := kv.FormatCommand(KV_STATE_INIT); str != "KV_STATE_INIT" {
		t.Errorf("Expected KV_STATE_INIT, got %v\n", str)
	}
	if str := kv.FormatCommand(KV_STATE_ADD); str != "KV_STATE_ADD" {
		t.Errorf("Expected KV_STATE_ADD, got %v\n", str)
	}
	if str := kv.FormatCommand(KV_STATE_GET); str != "KV_STATE_GET" {
		t.Errorf("Expected KV_STATE_GET, got %v\n", str)
	}
	if str := kv.FormatCommand(KV_STATE_REMOVE); str != "KV_STATE_REMOVE" {
		t.Errorf("Expected KV_STATE_REMOVE, got %v\n", str)
	}
	if str := kv.FormatCommand(99); str != "UNKNOWN_COMMAND" {
		t.Errorf("Expected UNKNOWN_COMMAND, got %v\n", str)
	}
}

func TestApplyCommand(t *testing.T) {
	kv := new(KVStateMachine)

	// Test initializing hash machine
	initialValue := []byte("")
	_, err := kv.ApplyCommand(KV_STATE_INIT, initialValue)

	if err != nil {
		t.Error(err)
	}

	// add key value pair
	keyvalpair := []byte("hello.goodbye")
	_, err = kv.ApplyCommand(KV_STATE_ADD, keyvalpair)
	if err != nil {
		t.Error(err)
	}

	// now test getting it
	keybyte := []byte("hello")
	val, err := kv.ApplyCommand(KV_STATE_GET, keybyte)
	if string(val) != "goodbye" {
		t.Errorf("Expected %v, got %v", "goodbye", string(val))
	}
	if err != nil {
		t.Error(err)
	}

	// now test remove
	_, err = kv.ApplyCommand(KV_STATE_REMOVE, keybyte)
	if err != nil {
		t.Error(err)
	}

	// check if it actually removed lol
	val, err = kv.ApplyCommand(KV_STATE_GET, keybyte)
	if string(val) != "" {
		t.Errorf("Expected %v, got %v", "", string(val))
	}
	if err != nil {
		t.Error(err)
	}

	// test unkown command
	_, err = kv.ApplyCommand(99, nil)

	if err == nil {
		t.Error("Expected error when applying unknown command, did not get any")
	}
}
