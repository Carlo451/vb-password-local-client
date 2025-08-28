package storefuncs

import (
	"github.com/Carlo451/vb-password-base-package/passwordstoreFilesystem"
)

func retrieveContentOutOfContentDir(dir passwordstoreFilesystem.PasswordStoreContentDir, fileName string) string {
	for _, file := range dir.ReturnFiles() {
		if file.GetFileName() == fileName {
			return file.GetContent()
		}
	}
	return ""
}

func goDownTreeAndReencryptContent(dir passwordstoreFilesystem.PasswordStoreDir, encryptionId, passphrase string) error {
	contentDirs := dir.GetContentDirectories()
	for _, contentDir := range contentDirs {
		if contentDir.GetDirName() == configDirName {
			continue
		}
		contentFiles := contentDir.ReturnFiles()
		for _, contentFile := range contentFiles {
			encryptedContent := contentFile.GetContent()
			decryptedContent, decryptErr := DecryptContentWithEncryptionIdAndPassword(encryptionId, encryptedContent, passphrase)
			if decryptErr != nil {
				return decryptErr
			}
			newlyEncryptedContent, encryptError := EncryptContentWithTempEncryptionId(encryptionId, decryptedContent)
			if encryptError != nil {
				return encryptError
			}
			contentFile.SetContent(newlyEncryptedContent)
		}
		contentDir.WriteFiles()
	}
	subDirs := dir.GetStoreDirectories()
	for _, nextDir := range subDirs {
		return goDownTreeAndReencryptContent(nextDir, encryptionId, passphrase)
	}
	return nil
}
