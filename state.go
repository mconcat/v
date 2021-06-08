package v

import sdk "github.com/cosmos/cosmos-sdk/types"

type Stateful interface {
	// Lock executes fn with the Stateful locked. External calls are called
	// immedietly.
	// Any async transaction including interchain transactions are not allowed.
	Lock(fn func())

	// Atomic executes fn with the Stateful inside the transactional memory.
	// All effects will be reverted if Stateful has been modified during the call.
	// Async transactions are allowed.
	Atomic(fn func())
}

type State interface {
	Stateful

	Substate(key string) State
}

// Collection is set of pointers.
// Array, mapping, tree, things like that
type Collection interface {
	Stateful

	// Pointer returns pointer to a key-value
	// Pointer may not be locked, use explicit Lock() in such case
	Pointer(key string) Pointer // use ocap instead of string
}

// Pointer is pointer to a specific key-value pair
// That could be
type Pointer interface {
	Stateful

	// CN 1
	Read() interface{}
	Write() chan<- interface{}

	// CN 2
	// Update() chan interface{}
}

type pointer struct {
	ctx sdk.Context
	key string
}

type RuntimeTypecheckingPointer struct {
	Pointer
	checker func(interface{}) interface{}
}

func Typecheck(ptr Pointer, checker func(interface{}) interface{}) Pointer {
	return RuntimeTypecheckingPointer{ptr, checker}
}

func (ptr RuntimeTypecheckingPointer) Write(v interface{}) {
	ptr.Pointer.Write(ptr.checker(v))
}

type StateConstructor struct {
	Key string
}
