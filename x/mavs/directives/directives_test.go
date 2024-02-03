package directives_test

import (
	"testing"
	"time"

	"github.com/emaforlin/mAVS/x/mavs/directives"
	"github.com/stretchr/testify/require"
)

func TestVotingTimeFromString(t *testing.T) {
	timeString := "3099.11.10.08.00.00 3099.11.10.14.00.00"
	timeT := directives.VotingTimeFromString(timeString)

	expected, err := time.Parse(time.DateTime, "3099-11-10 08:00:00")

	if *timeT.Start() != expected || err != nil {
		t.Fatalf("error parsing time, got: %v, expected: %v. Err: %v", timeT.Start(), expected, err)
	}

	expected, err = time.Parse(time.DateTime, "3099-11-10 14:00:00")
	if *timeT.End() != expected || err != nil {
		t.Fatalf("error parsing time, got: %v, expected: %v. Err: %v", timeT.Start(), expected, err)
	}

}

func TestNewVotingTime(t *testing.T) {
	start := time.Now()
	end := time.Now().Add(2 * time.Hour)

	time := directives.NewVotingTime(start, end)
	timeStr := time.String()
	timeFromStr := directives.VotingTimeFromString(timeStr)

	require.Equal(t, time.String(), timeFromStr.String())
}
