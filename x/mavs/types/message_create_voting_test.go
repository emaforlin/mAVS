package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/emaforlin/mAVS/testutil/sample"
	"github.com/emaforlin/mAVS/x/mavs/types"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateVoting_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateVoting
		err  error
	}{
		{
			name: "invalid creator address",
			msg: types.MsgCreateVoting{
				Creator:    "invalid_address",
				Title:      "valid_title",
				Candidates: "candidate list separated by a space like this",
				Timewindow: "3099.11.10.8.0.0 3099.11.10.14.0.0.0", //time window valid for a while
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "votation ended",
			msg: types.MsgCreateVoting{
				Creator:    sample.AccAddress(),
				Title:      "This voting has already ended",
				Candidates: "candidate1 candidate2 candidate3",
				Timewindow: "2013.8.21.12.0.0.0 2013.8.21.23.30.0", //time window caducated in 2013
			},
			err: types.ErrVotationEnded,
		}, {
			name: "invalid title",
			msg: types.MsgCreateVoting{
				Creator:    sample.AccAddress(),
				Title:      "",
				Candidates: "candidate1 candidate2 candidate3",
				Timewindow: "3099.11.10.8.0.0 3099.11.10.14.0.0.0",
			},
			err: types.ErrBadStoredVoting,
		}, {
			name: "empty candidate list",
			msg: types.MsgCreateVoting{
				Creator:    sample.AccAddress(),
				Title:      "candidate list empty",
				Candidates: "",
				Timewindow: "3099.11.10.8.0.0 3099.11.10.14.0.0.0",
			},
			err: types.ErrNotEnoughCandidates,
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
