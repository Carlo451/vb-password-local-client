package keyring

import (
	"log"

	"github.com/zalando/go-keyring"
)

func savePassword(passwordIdentifier, user, password string) error {
	err := keyring.Set(passwordIdentifier, user, password)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func loadPassword(passwordIdentifier, user string) (string, error) {
	secret, err := keyring.Get(passwordIdentifier, user)
	if err != nil {
		log.Fatal(err)
	}
	return secret, nil
}

func deletePassword(passwordIdentifier, user string) error {
	err := keyring.Delete(passwordIdentifier, user)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func buildIdentWithEncryptionId(encryptionId string) string {
	return keyRingIdent + "-" + encryptionId
}
