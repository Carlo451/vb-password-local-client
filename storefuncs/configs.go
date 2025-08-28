package storefuncs

import (
	"path/filepath"

	"github.com/Carlo451/vb-password-base-package/api"
	"github.com/Carlo451/vb-password-base-package/passwordstoreFilesystem"
)

const configDirName = "configs"
const OwnerIdent = "owner"
const EncryptionIdent = "encryptionId"

func CreateConfig(owner, encryptionId string) []passwordstoreFilesystem.PasswordStoreContentFile {
	ownerFile := passwordstoreFilesystem.NewUnrelatedPasswordStoreContentFile(owner, OwnerIdent)
	enryptionIdFile := passwordstoreFilesystem.NewUnrelatedPasswordStoreContentFile(encryptionId, EncryptionIdent)
	return []passwordstoreFilesystem.PasswordStoreContentFile{
		*ownerFile, *enryptionIdFile,
	}
}

func ReadConfigDirectory(storeName string, handler api.PasswordStoreHandler) passwordstoreFilesystem.PasswordStoreContentDir {
	configPath := filepath.Join(handler.GetPath(), storeName+"/"+configDirName)
	content, err := handler.ReadContentDir(configPath, storeName)
	if err != nil {
		panic(err)
	}
	return *content
}

func ReadConfigEntry(configPart, storeName string, handler api.PasswordStoreHandler) string {
	contentDir := ReadConfigDirectory(storeName, handler)
	return retrieveContentOutOfContentDir(contentDir, configPart)
}

func ReadStoreEncryptionId(storeName string, handler api.PasswordStoreHandler) string {
	return ReadConfigEntry("encryptionId", storeName, handler)
}

func ReadConfigOwner(storeName string, handler api.PasswordStoreHandler) string {
	return ReadConfigEntry(OwnerIdent, storeName, handler)
}

func ReadConfigEncryptionId(storeName string, handler api.PasswordStoreHandler) string {
	return ReadConfigEntry(EncryptionIdent, storeName, handler)
}
