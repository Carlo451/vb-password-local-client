package exec

import "github.com/Carlo451/vb-password-local-client/storefuncs"

func AddPasswordToStore(storeName, contentPath, password string) error {
	return AddContentToStore(storeName, contentPath, "password", password)
}

func AddUsernameToStore(storeName, contentPath, username string) error {
	return AddContentToStore(storeName, contentPath, "username", username)
}

func AddEmailToStore(storeName, contentPath, email string) error {
	return AddContentToStore(storeName, contentPath, "email", email)

}

func AddContentToStore(storeName, contentPath, identifier, content string) error {
	handler := storefuncs.CreateHandler()
	encryptionId := storefuncs.ReadStoreEncryptionId(storeName, handler)
	encryptedContent, err := storefuncs.EncryptContentWithEncryptionId(encryptionId, content)
	if err != nil {
		return err
	}
	return storefuncs.WriteNewContentToStore(storeName, contentPath, identifier, encryptedContent)
}
