package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddVoter{}

func NewMsgAddVoter(creator string, votingId string, dni string, proofId string) *MsgAddVoter {
	return &MsgAddVoter{
		Creator:  creator,
		VotingId: votingId,
		Dni:      dni,
		ProofId:  proofId,
	}
}

func (msg *MsgAddVoter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Dni == "" {
		return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "dni cannot be empty")
	}

	if msg.VotingId == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "voting id cannot be empty")
	}

	if msg.ProofId == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "proof id cannot be empty")
	}

	return nil
}
