package exec

import "github.com/Carlo451/vb-password-local-client/storefuncs"

func CreatePasswordStoreWithExistingEncryptionId(storeName, encryptionId string) {
	storefuncs.CreatePassStore(storeName, encryptionId)
}

func CreatePasswordStoreWithNewEncryptionId(storeName, encryptionId, passphrase string) {
	storefuncs.WriteNewKeyPairs(encryptionId, passphrase)
}

func CreatePasswordStoreWithMasterPassword(storeName string) {
	storefuncs.CreatePassStore(storeName, storefuncs.GetMasterEncryptionId())
}
