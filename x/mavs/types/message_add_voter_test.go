package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/emaforlin/mAVS/testutil/sample"
	"github.com/emaforlin/mAVS/x/mavs/types"
	"github.com/stretchr/testify/require"
)

func TestMsgAddVoter_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgAddVoter
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgAddVoter{
				Creator:  "invalid_address",
				VotingId: "1",
				Dni:      "12345678",
				ProofId:  "1234",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgAddVoter{
				Creator:  sample.AccAddress(),
				VotingId: "1",
				Dni:      "12345678",
				ProofId:  "1234",
			},
		}, {
			name: "invalid voting id (too low)",
			msg: types.MsgAddVoter{
				Creator:  sample.AccAddress(),
				VotingId: "-10",
				Dni:      "12345678",
				ProofId:  "proofId",
			},
			err: types.ErrInvalidVotingId,
		}, {
			name: "invalid voting id (empty)",
			msg: types.MsgAddVoter{
				Creator:  sample.AccAddress(),
				VotingId: "",
				Dni:      "12345678",
				ProofId:  "proofId",
			},
			err: types.ErrInvalidVotingId,
		}, {
			name: "invalid voter dni",
			msg: types.MsgAddVoter{
				Creator:  sample.AccAddress(),
				VotingId: "1",
				Dni:      "this is not my dni",
				ProofId:  "proofId",
			},
			err: types.ErrInvalidVoter,
		}, {
			name: "invalid voter dni (too low)",
			msg: types.MsgAddVoter{
				Creator:  sample.AccAddress(),
				VotingId: "1",
				Dni:      "1234",
				ProofId:  "proofId",
			},
			err: types.ErrInvalidVoter,
		}, {
			name: "invalid proof id",
			msg: types.MsgAddVoter{
				Creator:  sample.AccAddress(),
				VotingId: "1",
				Dni:      "12345678",
				ProofId:  "",
			},
			err: types.ErrInvalidVoterIdProof,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
