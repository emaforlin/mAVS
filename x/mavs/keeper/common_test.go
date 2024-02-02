package keeper_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/emaforlin/mAVS/testutil"
	"github.com/emaforlin/mAVS/x/mavs/directives"
	"github.com/emaforlin/mAVS/x/mavs/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
	carol = testutil.Carol
)

func TestCreateVoting(t *testing.T) {
	_, msgServer, context := setupMsgServer(t)

	createResponse, err := msgServer.CreateVoting(context, &types.MsgCreateVoting{
		Creator:    alice,
		Title:      "test voting",
		Timewindow: fmt.Sprintf("%s %s", directives.ConvertTimeToString(time.Now()), directives.ConvertTimeToString(time.Now().Add(6*time.Hour))),
		Candidates: "candidate1 candidate2 candidate3",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateVotingResponse{
		VotingIndex: "",
	}, *createResponse)

}
