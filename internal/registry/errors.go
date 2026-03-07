package registry

const (
	REGISTRY_ENTRY_EXISTS = iota
	REGISTRY_ENTRY_NOT_EXISTS
	REGISTRY_ENTRY_MISMATCH
)

const notExistsError = "Entry %v is not present in the registry"
const existsError = "Entry by key %v already exists in registry"
const mismatchError = "Entry %v exists in registry, but is different to one provided"
