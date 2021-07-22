// Package localstore contains a single function to store a payload to a temporary store.
// Use in case of network errors.
package localstore

import (
	"fmt"

	"github.com/google/uuid"
	"git.mills.io/prologic/bitcask"
)

// SaveToLocalStore function to store payload to a temporary store
func SaveToLocalStore(payload []byte) error {
	db, err := bitcask.Open("/tmp/db")
	if err != nil {
		return fmt.Errorf("Error setting up local store: %v", err)
	}

	defer db.Close()

	id := uuid.New()
	db.Put(id.NodeID(), payload)

	return nil
}
