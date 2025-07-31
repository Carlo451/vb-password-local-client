package storefuncs

import (
	"github.com/Carlo451/vb-password-base-package/passwordstore/passwordstoreFilesystem"
	"strings"
)

func retriveContentOutOfContentDir(dir passwordstoreFilesystem.PasswordStoreContentDir, fileName string) string {
	for _, file := range dir.ReturnFiles() {
		if file.GetFileName() == fileName {
			return strings.Trim(file.GetContent(), "\n")
		}
	}
	return ""
}
