package types

import (
	"fmt"
	"strings"
	"time"
)

func (s StoredVoting) CountVotes() (votes uint64) {
	return s.Votes
}

func (s StoredVoting) StillActive() (err error) {
	tw := strings.Split(s.Timewindow, " ")
	for i, j := range tw {
		year := j[:4]
		month := j[4:6]
		day := j[6:8]
		hours := j[8:10]
		minutes := j[10:12]
		seconds := j[12:]
		var dateSep string = "-"
		var timeSep string = ":"
		tw[i] = year + dateSep + month + dateSep + day + " " + hours + timeSep + minutes + timeSep + seconds
	}

	start, err := time.Parse(time.DateTime, tw[0])
	if err != nil {
		fmt.Printf("Raw Time: %v\nStart time %v\n", tw[0], start)
		return err
	}
	end, err := time.Parse(time.DateTime, tw[1])
	if err != nil {
		return err
	}
	if end.Sub(start) < time.Hour {
		return ErrVotingTimeTooShort
	}
	if end.Before(time.Now()) {
		return ErrVotationEnded
	}
	return nil
}

func (s StoredVoting) Validate() (err error) {
	if s.Index == "" {
		return ErrBadStoredVoting
	}
	if s.Title == "" {
		return ErrInvalidVotingTitle
	}

	return nil
}
