package types

const (
	// ModuleName defines the module name
	ModuleName = "flashloan"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_flashloan"
)

var (
	ParamsKey = []byte("p_flashloan")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
