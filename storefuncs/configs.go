package storefuncs

import (
	"github.com/Carlo451/vb-password-base-package/api"
	"github.com/Carlo451/vb-password-base-package/passwordstore/passwordstoreFilesystem"
	"path/filepath"
)

const configDirName = "configs"
const OwnerIdent = "owner"
const EncryptionIdent = "enryptionId"

func CreateConfig(owner, encryptionId string) []passwordstoreFilesystem.PasswordStoreContentFile {
	ownerFile := passwordstoreFilesystem.NewUnrelatedPasswordStoreContentFile(owner, OwnerIdent)
	enryptionIdFile := passwordstoreFilesystem.NewUnrelatedPasswordStoreContentFile(EncryptionIdent, encryptionId)
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
	return retriveContentOutOfContentDir(contentDir, configPart)
}
