package types

import (
	"strconv"

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

	// voting id checks
	msgId, _ := strconv.ParseUint(msg.VotingId, 10, 64)
	if msgId < DefaultIndex {
		return errorsmod.Wrapf(ErrInvalidVotingId, "number to low (%d)", msgId)
	}

	if msg.VotingId == "" {
		return errorsmod.Wrapf(ErrInvalidVotingId, "voting id cannot be empty")
	}

	// dni checks
	dni, err := strconv.ParseUint(msg.Dni, 10, 64)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidVoter, "dni must to be a number")
	}

	if dni < 1e6 {
		return errorsmod.Wrapf(ErrInvalidVoter, "dni number too low")
	}

	// proof id checks
	if msg.ProofId == "" {
		return errorsmod.Wrapf(ErrInvalidVoterIdProof, "proof id cannot be empty")
	}

	return nil
}
