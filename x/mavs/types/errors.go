package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/mavs module sentinel errors
var (
	ErrInvalidSigner        = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample               = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrVotationEnded        = sdkerrors.Register(ModuleName, 1102, "time window for voting has ended")
	ErrCandidateNotFound    = sdkerrors.Register(ModuleName, 1103, "the candidate does not exists")
	ErrBadStoredVoting      = sdkerrors.Register(ModuleName, 1104, "error when voting was stored")
	ErrVotingTimeTooShort   = sdkerrors.Register(ModuleName, 1105, "minimum voting is one(1) hour")
	ErrVotingCannotBeParsed = sdkerrors.Register(ModuleName, 1106, "voting cannot be parsed")
	ErrNotEnoughCandidates  = sdkerrors.Register(ModuleName, 1107, "need more candidates")
)
