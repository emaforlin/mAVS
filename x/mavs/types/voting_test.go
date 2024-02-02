package types_test

import (
	"testing"

	"github.com/emaforlin/mAVS/x/mavs/types"
	"github.com/stretchr/testify/require"
)

func GetStoredVoting1() types.StoredVoting {
	return types.StoredVoting{
		Index:      "1",
		Title:      "Stored Voting 1",
		Timewindow: "30991110080000 30991110090000", // make sure the test dates are valid for a while
		Votes:      22,
	}
}

func TestCanGetStoredVoting(t *testing.T) {
	voting := GetStoredVoting1()

	err1 := voting.Validate()
	require.Nil(t, err1)
}

func TestWrongGetStoredVoting(t *testing.T) {
	voting := GetStoredVoting1()

	voting.Index = "" // introduce bad data
	voting.Title = ""

	err := voting.Validate()
	require.EqualError(t, err, types.ErrBadStoredVoting.Error())
}

func TestCheckVotingStillActive(t *testing.T) {
	voting := GetStoredVoting1()
	err := voting.StillActive()
	require.Nil(t, err)
}

func TestCheckWrongVotingStillActive(t *testing.T) {
	voting := GetStoredVoting1()
	voting.Timewindow = "20131110080000 20131110090000"
	require.EqualError(t, voting.StillActive(), types.ErrVotationEnded.Error())
}
