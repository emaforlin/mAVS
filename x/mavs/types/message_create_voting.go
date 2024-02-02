package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateVoting{}

func NewMsgCreateVoting(creator string, title string, timewindow string, candidates string) *MsgCreateVoting {
	return &MsgCreateVoting{
		Creator:    creator,
		Title:      title,
		Timewindow: timewindow,
		Candidates: candidates,
	}
}

func (msg *MsgCreateVoting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
