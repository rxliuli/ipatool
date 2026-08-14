package keychain

import (
	"fmt"

	"github.com/byteness/keyring"
)

func (k *keychain) Set(key string, data []byte) error {
	err := k.keyring.Set(keyring.Item{
		Key:  key,
		Data: data,
		// Set a human-readable label so that macOS Keychain access
		// prompts show a meaningful item name instead of an empty string.
		Label: "ipatool-auth.service",
	})
	if err != nil {
		return fmt.Errorf("failed to set item: %w", err)
	}

	return nil
}
