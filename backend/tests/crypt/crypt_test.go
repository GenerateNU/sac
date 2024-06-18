package crypt

import (
	"testing"

	"github.com/GenerateNU/sac/backend/integrations/oauth/crypt"
)

func TestCryptEncryptDecrypt(t *testing.T) {
	t.Parallel()

	data := "test data"
	passphrase := "test passphrase"

	encrypted, err := crypt.Encrypt(data, passphrase)
	if err != nil {
		t.Fatal(err)
	}

	decrypted, err := crypt.Decrypt(encrypted, passphrase)
	if err != nil {
		t.Fatal(err)
	}

	if decrypted != data {
		t.Fatalf("expected %s, got %s", data, decrypted)
	}
}
