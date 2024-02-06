package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/mavs module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")

	ErrVotationEnded        = sdkerrors.Register(ModuleName, 1102, "time window for voting has ended")
	ErrBadStoredVoting      = sdkerrors.Register(ModuleName, 1103, "error when voting was stored")
	ErrVotingCannotBeParsed = sdkerrors.Register(ModuleName, 1104, "voting cannot be parsed")
	ErrVotingTimeTooShort   = sdkerrors.Register(ModuleName, 1105, "minimum voting is one(1) hour")
	ErrInvalidVotingId      = sdkerrors.Register(ModuleName, 1106, "invalid voting id")

	ErrCandidateNotFound   = sdkerrors.Register(ModuleName, 1107, "the candidate does not exists")
	ErrNotEnoughCandidates = sdkerrors.Register(ModuleName, 1108, "need more candidates")

	ErrInvalidVoter        = sdkerrors.Register(ModuleName, 1109, "invalid voter data")
	ErrInvalidVoterIdProof = sdkerrors.Register(ModuleName, 1110, "invalid proof id")
)
