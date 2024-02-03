package directives

import (
	"fmt"
	"strings"
	"time"
)

type Voter struct {
	Name     string
	Lastname string
	Dni      uint64
	HasVoted bool
}

type Party string

type ElectoralRoll []Voter

var BlankVote = "blank"

type VotingTimeImpl struct {
	votingStart    time.Time
	votingEnd      time.Time
	votingDuration time.Duration
}

type VotingImpl struct {
	Title      string
	TimeWindow VotingTimeImpl
	VoteCount  map[string]uint64
}

func (v VotingImpl) StartTime() *time.Time {
	return v.TimeWindow.Start()
}

func (v VotingImpl) EndTime() *time.Time {
	return v.TimeWindow.End()
}

func (v VotingTimeImpl) String() string {
	s := v.votingStart
	return fmt.Sprintf("%d%d%d%d%d%d+%d", s.Year(), s.Month(), s.Day(), s.Hour(), s.Minute(), s.Day(), v.votingDuration)
}

func (v VotingTimeImpl) Start() *time.Time {
	return &v.votingStart
}

func (v VotingTimeImpl) End() *time.Time {
	return &v.votingEnd
}

func VotingTimeFromString(s string) VotingTimeImpl {
	timeStrings := make([]string, 2)
	for i, t := range strings.Split(s, " ") {
		timeStrings[i] = fmt.Sprintf("%s-%s-%s %s:%s:%s", t[:4], t[4:6], t[6:8], t[8:10], t[10:12], t[12:14])
	}
	start, _ := time.Parse(time.DateTime, timeStrings[0])
	end, _ := time.Parse(time.DateTime, timeStrings[1])

	return VotingTimeImpl{
		votingStart:    start.UTC(),
		votingEnd:      end.UTC(),
		votingDuration: end.Sub(start),
	}
}

func NewVoting(title string, timeWindow VotingTimeImpl, roll *ElectoralRoll, parties ...Party) *VotingImpl {
	partyList := make(map[string]uint64, len(parties))
	partyList["blank"] = 0

	for _, partyName := range parties {
		partyList[string(partyName)] = 0
	}
	return &VotingImpl{
		Title:      title,
		TimeWindow: timeWindow,
		VoteCount:  partyList,
	}
}

func Parse(title string, tw string, vc *map[string]uint64) *VotingImpl {
	res := &VotingImpl{}
	res.Title = title
	res.TimeWindow = VotingTimeFromString(tw)
	res.VoteCount = *vc
	return res
}

func (v VotingImpl) StillActive() bool {
	now := time.Now()
	end := v.EndTime()
	return !end.Before(now)
}
