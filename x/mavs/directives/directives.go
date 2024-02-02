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

type VotingImpl struct {
	Title      string
	TimeWindow VotingTime
	VoteCount  map[Party]uint64
}

func ConvertTimeToString(t time.Time) string {
	return fmt.Sprintf("%d%d%d%d%d%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func NewVoting(title string, startAt time.Time, endAt time.Time, roll *ElectoralRoll, parties ...Party) *VotingImpl {
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
	emptyVotes := make(map[Party]uint64)
	emptyVotes[Party("blank")] = 0

	return &VotingImpl{
		Title:      title,
		TimeWindow: timeWindow,
		VoteCount:  emptyVotes,
	}
}

func (v VotingImpl) IsActive() bool {
	return v.TimeWindow.active()
}
