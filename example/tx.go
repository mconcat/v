package example

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mconcat/v"
)

var _ sdk.Msg = &MsgExample{}

func (msg MsgExample) Route() string {
	return ""
}

func (msg MsgExample) Type() string {
	return ""
}

func (msg MsgExample) ValidateBasic() error {
	return nil
}

func (msg MsgExample) GetSignBytes() []byte {
	return nil
}

func (msg MsgExample) GetSigners() []sdk.AccAddress {
	return nil
}

func (msg MsgExample) State() v.State {
	return
}

type PointerThing struct{ v.Pointer }

type State struct {
}
