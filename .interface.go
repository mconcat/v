package v

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Executable func(ctx sdk.Context) ExecuteResult

// CONTRACT: call ExecuteResult.Do() only once.
// CONTRACT: access Result() and Error() only inside
// Do() and only once. Do not store channel.
type ExecuteResult interface {
	Do(func())
	Result() chan interface{}
	Error() chan error
}

type Pointer interface {
	Read() Executable
	Write(value interface{}) Executable
}

type Process interface {
	RegisterPointer(ptr Pointer)
}
