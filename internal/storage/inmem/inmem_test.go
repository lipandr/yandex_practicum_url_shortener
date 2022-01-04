package inmem

import (
	"testing"
)

func TestNewStorage(t *testing.T) {
	store := NewStorage()
	if len(store.data) != 0 {
		t.Errorf("Invalid non 0 store")
	}
	if store == nil {
		t.Errorf("Store not initialized")
	}
}
