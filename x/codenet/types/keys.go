package types

const (
	// ModuleName defines the module name
	ModuleName = "codenet"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_codenet"

	// Version defines the current version the IBC module supports
	Version = "codenet-1"

	// PortID is the default port id that module binds to
	PortID = "codenet"
)

var (
	ParamsKey = []byte("p_codenet")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("codenet-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
