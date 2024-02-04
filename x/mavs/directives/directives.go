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

func NewVotingTime(start, end time.Time) VotingTimeImpl {
	return VotingTimeImpl{
		votingStart:    start,
		votingEnd:      end,
		votingDuration: end.Sub(start),
	}
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
	e := v.votingEnd
	sStr := fmt.Sprintf("%d.%d.%d.%d.%d.%d", s.Year(), s.Month(), s.Day(), s.Hour(), s.Minute(), s.Second())
	eStr := fmt.Sprintf("%d.%d.%d.%d.%d.%d", e.Year(), e.Month(), e.Day(), e.Hour(), e.Minute(), e.Second())
	return fmt.Sprintf("%s %s", sStr, eStr)
}

func (v VotingTimeImpl) Start() *time.Time {
	return &v.votingStart
}

func (v VotingTimeImpl) End() *time.Time {
	return &v.votingEnd
}

func (v VotingTimeImpl) StillActive() bool {
	now := time.Now().UTC()
	end := v.End().UTC()
	return !end.Before(now)
}

func VotingTimeFromString(s string) VotingTimeImpl {
	//time format <year>.<month>.<day>.<hour>.<minute>.<seconds>
	tz := make([]string, 2)
	for i, t := range strings.Split(s, " ") {
		tp := strings.Split(t, ".")
		for k, l := range tp {
			if len(l) == 1 {
				tp[k] = fmt.Sprintf("0%s", l)
			}
		}
		tz[i] = fmt.Sprintf("%s-%s-%s %s:%s:%s", tp[0], tp[1], tp[2], tp[3], tp[4], tp[5])
	}
	start, err := time.Parse(time.DateTime, tz[0])
	if err != nil {
		fmt.Println(err)
	}
	end, err := time.Parse(time.DateTime, tz[1])
	if err != nil {
		fmt.Println(err)
	}

	return NewVotingTime(start.UTC(), end.UTC())
}

func NewVoting(title string, timeWindow VotingTimeImpl, roll *ElectoralRoll, parties ...string) *VotingImpl {
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
