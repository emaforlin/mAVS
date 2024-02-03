package directives_test

import (
	"testing"
	"time"

	"github.com/emaforlin/mAVS/x/mavs/directives"
)

func TestVotingTimeFromString(t *testing.T) {
	timeString := "30991110080000 30991110140000"
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
