package types

const (
	// ModuleName defines the module name
	ModuleName = "flashloanchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_flashloanchain"
)

var (
	ParamsKey = []byte("p_flashloanchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
