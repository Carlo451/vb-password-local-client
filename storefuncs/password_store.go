package storefuncs

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Carlo451/vb-password-base-package/api"
	"github.com/Carlo451/vb-password-base-package/pathparser"
)

func WriteNewContentToStore(storeName, pathOfNewContentDir, identifier, content string) error {
	handler := CreateHandler()
	parts := strings.Split(pathOfNewContentDir, "/")
	contentDirName := parts[len(parts)-1]

	absoluteContentDirPath := filepath.Join(handler.GetPath(), storeName, pathOfNewContentDir)

	pathParser := pathparser.ParsePathWithContentDirectory(filepath.Join(handler.GetPath(), storeName), absoluteContentDirPath)
	absolutePathOfLastSubfolder := filepath.Join(handler.GetPath(), storeName, pathParser.BuildPathWithoutContentDir())

	exists, _ := api.CheckIfContentDirectoryExists(absoluteContentDirPath)
	if exists {
		handler.UpdateContentInContentDirectory(absoluteContentDirPath, storeName, content, identifier)
	} else {
		handler.AddContentDirectoryToStore(absolutePathOfLastSubfolder, storeName, contentDirName, content, identifier)
	}
	_, err := os.Stat(filepath.Join(absoluteContentDirPath, identifier))
	if err != nil {
		return err
	}
	return nil
}

func ReadContentFromStore(storeName, contentPath, identifier string) (string, error) {
	handler := CreateHandler()
	absolutePath := filepath.Join(handler.GetPath(), storeName, contentPath)
	content, err := handler.ReadContentDir(absolutePath, storeName)
	if err != nil {
		return "", err
	}
	return retrieveContentOutOfContentDir(*content, identifier), nil
}
