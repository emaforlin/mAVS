package directives

import (
	"time"
)

type Voter struct {
	Name     string
	Lastname string
	Dni      uint64
	HasVoted bool
}

type ElectoralRoll []Voter

type Party string

var BlankVote = Party("blank")

type VotingTime struct {
	Start time.Time
	End   time.Time
}

func (t VotingTime) active() bool {
	return t.End.Sub(t.Start) > time.Duration(0)
}

type Voting struct {
	TimeWindow VotingTime
	Roll       *ElectoralRoll
	Candidates []Party
	VoteCount  map[Party]uint64
}

func NewVoting(startAt time.Time, endAt time.Time, roll *ElectoralRoll, parties ...Party) *Voting {
	timeWindow := VotingTime{startAt, endAt}
	if timeWindow.End.Sub(timeWindow.Start) <= 0 {
		return nil
	}
	if len(parties) == 0 {
		return nil
	}

	parties = append(parties, BlankVote)

	return &Voting{
		TimeWindow: timeWindow,
		Roll:       roll,
		Candidates: parties,
		VoteCount:  make(map[Party]uint64),
	}
}

func (v *Voting) IsActive() bool {
	return v.TimeWindow.active()
}
