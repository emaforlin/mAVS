package types

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/emaforlin/mAVS/x/mavs/directives"
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
	if !directives.VotingTimeFromString(msg.Timewindow).StillActive() {
		return ErrVotationEnded
	}
	if len(strings.Split(msg.Candidates, " ")) <= 1 {
		return ErrNotEnoughCandidates
	}
	if msg.Title == "" {
		return ErrBadStoredVoting
	}
	return nil
}
