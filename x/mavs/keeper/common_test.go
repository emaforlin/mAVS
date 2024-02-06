package keeper_test

import (
	"time"

	"github.com/emaforlin/mAVS/testutil"
	"github.com/emaforlin/mAVS/testutil/sample"
	"github.com/emaforlin/mAVS/x/mavs/directives"
	"github.com/emaforlin/mAVS/x/mavs/types"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
	carol = testutil.Carol
)

var DefaultTestingMsgVoting = types.MsgCreateVoting{
	Creator:    sample.AccAddress(),
	Title:      "Default voting for testing purpouses only",
	Timewindow: directives.NewVotingTime(time.Now(), time.Now().Add(1*time.Hour)).String(),
	Candidates: "Cadidate_1 Candidate_2 Candidate_3 Candidate_4",
}
