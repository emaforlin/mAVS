package types_test

import (
	"testing"

	"github.com/emaforlin/mAVS/x/mavs/types"
	"github.com/stretchr/testify/require"
)

func GetStoredVoting1() types.StoredVoting {
	return types.StoredVoting{
		Index:      "1",
		Title:      "Test voting",
		Timewindow: "30991110080000 30991110140000", // 3099/11/11 08:00:00 +06h
		Counting: map[string]uint64{ // make sure the test dates are valid for a while
			"partido 1": 24,
			"partido 2": 15,
			"partido 3": 22,
		},
	}
}

func TestCanGetStoredVoting(t *testing.T) {
	voting := GetStoredVoting1()

	err := voting.Validate()
	require.Nil(t, err)
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
	err := voting.Validate()
	require.Nil(t, err)
}

func TestCheckWrongVotingStillActive(t *testing.T) {
	voting := GetStoredVoting1()
	voting.Timewindow = "20131110080000 20131110140000"
	require.EqualError(t, voting.Validate(), types.ErrVotationEnded.Error())
}
