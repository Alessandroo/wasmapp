package types

const (
	// ModuleName defines the module name
	ModuleName = "interquery"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_interquery"

	// Version defines the current version the IBC module supports
	Version = "interquery-1"

	// PortID is the default port id that module binds to
	PortID = "interquery"
)

var (
	ParamsKey = []byte("p_interquery")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("interquery-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
