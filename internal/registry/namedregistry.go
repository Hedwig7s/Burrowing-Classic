package registry

import (
	"reflect"

	"github.com/Hedwig7s/Burrowing-Classic/internal/cerror"
)

type Named[K comparable] interface {
	Name() K
}

type NamedRegistry[K comparable, V Named[K]] struct {
	entries map[K]V
}

func (registry *NamedRegistry[K, V]) Register(entry V) error {
	key := entry.Name()
	if _, ok := registry.entries[key]; ok {
		return cerror.NewErrorf(REGISTRY_ENTRY_EXISTS, "Entry %v already exists in registry", key)
	}
	registry.entries[key] = entry
	return nil
}

func (registry *NamedRegistry[K, V]) Unregister(key K) error {
	if _, ok := registry.entries[key]; ok {
		delete(registry.entries, key)
		return nil
	}
	return cerror.NewErrorf(REGISTRY_ENTRY_NOT_EXISTS, notExistsError, key)

}

func (registry *NamedRegistry[K, V]) UnregisterByValue(entry V) error {
	key := entry.Name()
	var ok bool
	var existing V
	if existing, ok = registry.entries[key]; !ok {
		return cerror.NewErrorf(REGISTRY_ENTRY_NOT_EXISTS, notExistsError, key)
	}
	if reflect.ValueOf(existing).Pointer() != reflect.ValueOf(entry).Pointer() { // Who tf made this language
		return cerror.NewErrorf(REGISTRY_ENTRY_MISMATCH, mismatchError, key)
	}
	delete(registry.entries, key)
	return nil
}

func NewNamedRegistry[K comparable, V Named[K]]() *NamedRegistry[K, V] {
	return &NamedRegistry[K, V]{}
}
