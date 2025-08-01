package exec

import "github.com/Carlo451/vb-password-local-client/storefuncs"

func AddContentToStore(storeName, contentPath, identifier, content string) error {
	handler := storefuncs.CreateHandler()
	encryptionId := storefuncs.ReadStoreEncryptionId(storeName, handler)
	encryptedContent, err := storefuncs.EncryptContentWithEncrypptionId(encryptionId, content)
	if err != nil {
		return err
	}
	return storefuncs.WriteNewContentToStore(storeName, contentPath, identifier, encryptedContent)
}

func AddPasswordToStore(storeName, contentPath, content string) error {
	return AddContentToStore(storeName, contentPath, "password", content)
}

func AddUsernameToStore(storeName, contentPath, content string) error {
	return AddContentToStore(storeName, contentPath, "username", content)
}

func AddEmailToStore(storeName, contentPath, content string) error {
	return AddContentToStore(storeName, contentPath, "email", content)

}
