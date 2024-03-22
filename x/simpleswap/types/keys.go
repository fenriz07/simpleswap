package types

const (
	// ModuleName defines the module name
	ModuleName = "simpleswap"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_simpleswap"
)

var (
	ParamsKey = []byte("p_simpleswap")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey = "SystemInfo/value/"
)
