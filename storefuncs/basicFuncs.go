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
