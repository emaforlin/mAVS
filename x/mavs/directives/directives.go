package directives

import (
	"fmt"
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
	Title      string
	TimeWindow VotingTime
	Roll       *ElectoralRoll
	Candidates []Party
	VoteCount  map[Party]uint64
}

func ConvertTimeToString(t time.Time) string {
	return fmt.Sprintf("%d%d%d%d%d%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func NewVoting(title string, startAt time.Time, endAt time.Time, roll *ElectoralRoll, parties ...Party) *Voting {
	timeWindow := VotingTime{startAt, endAt}
	if timeWindow.End.Sub(timeWindow.Start) <= 0 {
		return nil
	}
	if len(parties) == 0 {
		return nil
	}

	if title == "" {
		return nil
	}

	parties = append(parties, BlankVote)

	return &Voting{
		Title:      title,
		TimeWindow: timeWindow,
		Roll:       roll,
		Candidates: parties,
		VoteCount:  make(map[Party]uint64),
	}
}

func (v *Voting) IsActive() bool {
	return v.TimeWindow.active()
}
