package exec

import "github.com/Carlo451/vb-password-local-client/storefuncs"

func ReadPasswordFromStore(storeName, contentPath, passphrase string) (string, error) {
	return ReadContentFromStore(storeName, contentPath, "password", passphrase)
}

func ReadUsernameFromStore(storeName, contentPath, passphrase string) (string, error) {
	return ReadContentFromStore(storeName, contentPath, "username", passphrase)
}

func ReadEmailFromStore(storeName, contentPath, passphrase string) (string, error) {
	return ReadContentFromStore(storeName, contentPath, "email", passphrase)

}

func ReadContentFromStore(storeName, contentPath, identifier, passphrase string) (string, error) {
	handler := storefuncs.CreateHandler()
	encryptionId := storefuncs.ReadStoreEncryptionId(storeName, handler)
	content, readErr := storefuncs.ReadContentFromStore(storeName, contentPath, identifier)
	if readErr != nil {
		return "", readErr
	}
	decryptedContent, err := storefuncs.DecryptContentWithEncryptionIdAndPassword(encryptionId, content, passphrase)
	if err != nil {
		return "", err
	}
	return decryptedContent, nil
}
