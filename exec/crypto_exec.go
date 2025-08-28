package exec

import (
	"errors"

	"github.com/Carlo451/vb-password-local-client/storefuncs"
)

func ChangeKeyPairOfEncryptionId(passphrase, encryptionId string) error {
	keyPair, error := storefuncs.GetDecryptedKeyPair(encryptionId, passphrase)
	if error != nil {
		return error
	}
	validKeyPair, validationError := keyPair.CheckIfKeyPairIsValid()
	if validationError != nil {
		return validationError
	}
	if !validKeyPair {
		return errors.New("invalid key pair -> wrong passphrase")
	}
	storefuncs.CreateTempKeyPair(encryptionId, passphrase)
	handler := storefuncs.CreateHandler()
	baseStore := handler.ReadPasswordStore("")
	stores := baseStore.GetStoreDirectories()
	for _, store := range stores {
		storeEncryptionId := storefuncs.ReadConfigEncryptionId(store.GetDirName(), handler)
		if storeEncryptionId == encryptionId {
			errReencrypt := storefuncs.ReencryptStoreWithNewKeyPair(store.GetDirName(), passphrase)
			if errReencrypt != nil {
				return errReencrypt
			}
		}
	}
	storefuncs.OverwriteKeyPairWithTempKeyPair(encryptionId)
	return nil
}
