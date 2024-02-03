package types

import (
	"github.com/emaforlin/mAVS/x/mavs/directives"
)

func (s StoredVoting) CountVotes() (votes uint64) {
	var voteCount uint64
	for _, v := range s.Counting {
		voteCount += v
	}
	return uint64(voteCount)
}

func (s StoredVoting) GetVotesFrom(partyName string) (res uint64, err error) {
	res, found := s.Counting[partyName]
	if !found {
		return 0, ErrCandidateNotFound
	}
	return res, nil
}

func (s StoredVoting) ParseVoting() (voting *directives.VotingImpl, err error) {
	return directives.Parse(s.Title, s.Timewindow, &s.Counting), nil
}

func (s StoredVoting) Validate() error {
	v, err := s.ParseVoting()
	if v.Title == "" {
		return ErrBadStoredVoting
	}
	if !v.StillActive() {
		return ErrVotationEnded
	}
	if len(v.VoteCount) <= 1 {
		return ErrNotEnoughCandidates
	}

	if err != nil {
		return err
	}
	return nil
}
