package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredVotingKeyPrefix is the prefix to retrieve all StoredVoting
	StoredVotingKeyPrefix = "StoredVoting/value/"
)

// StoredVotingKey returns the store key to retrieve a StoredVoting from the index fields
func StoredVotingKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
