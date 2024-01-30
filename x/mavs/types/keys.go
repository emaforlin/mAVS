package types

const (
	// ModuleName defines the module name
	ModuleName = "mavs"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mavs"
)

var (
	ParamsKey = []byte("p_mavs")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
