package storefuncs

import (
	"github.com/Carlo451/vb-password-base-package/passwordstore/passwordstoreFilesystem"
	"os"
)

const dirName = "keystore"

func CreateCryptoStore(baseDirectory passwordstoreFilesystem.PasswordStoreDir) (bool, error) {
	contentDir := passwordstoreFilesystem.NewEmptyContentDirecotry(baseDirectory, dirName)
	contentDir.WriteDirectory()
	_, err := os.Stat(contentDir.GetAbsoluteDirectoryPath())
	if err != nil {
		return false, err
	}
	return true, nil
}
