package storefuncs

import (
	"errors"
	"github.com/Carlo451/vb-password-base-package/api"
	"github.com/Carlo451/vb-password-base-package/environment"
	"github.com/Carlo451/vb-password-base-package/passwordstore/passwordstoreFilesystem"
	"github.com/Carlo451/vb-password-base-package/pathparser"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func CheckIfBaseDirExists() (bool, error) {
	path, exists := environment.LookUpAndGetEnvValue("VB_PASSWORD_STORE_BASE_DIR_PATH")
	if !exists {
		return false, errors.New("the VB_PASSWORD_STORE_BASE_DIR_PATH env var is not set")
	}
	_, err := os.Stat(path)
	if err != nil {
		return false, nil
	}
	return true, nil

}

func CheckIfBaseDirExistsAndPanic() {
	exists, err := CheckIfBaseDirExists()
	if err != nil {
		panic(err)
	}
	if !exists {
		panic("No base directory found- it has to be initialised")
	}
}

func CreateHandler() api.PasswordStoreHandler {
	CheckIfBaseDirExistsAndPanic()
	path, _ := environment.LookUpAndGetEnvValue("VB_PASSWORD_STORE_BASE_DIR_PATH")
	return api.NewPasswordStoreHandler(path)
}

func CreateBaseDir() (*passwordstoreFilesystem.PasswordStoreDir, error) {
	path, exists := environment.LookUpAndGetEnvValue("VB_PASSWORD_STORE_BASE_DIR_PATH")
	if !exists {
		_, err := CheckIfBaseDirExists()
		return nil, err
	}
	os.MkdirAll(path, os.ModePerm)
	dirs := strings.Split(path, "/")
	name := dirs[len(dirs)-1]
	dirs = dirs[:len(dirs)-1]
	parentPath := strings.Join(dirs, "/")
	baseStore := passwordstoreFilesystem.CreateNewEmptyStoreDir(name, parentPath)

	return &baseStore, nil
}

func Init(masterPassword string) (bool, error) {
	exists, err := CheckIfBaseDirExists()
	if err != nil {
		return false, err
	}
	if !exists {
		base, errbase := CreateBaseDir()

		if errbase != nil {
			return false, err
		}
		_, errcrypto := CreateCryptoStore(*base, masterPassword)
		if errcrypto != nil {
			return false, err
		}

		return true, nil
	}
	return true, nil
}

func GetUsername() string {
	user, _ := user.Current()
	return user.Username
}

func CreatePassStore(name, encryptionId string) {
	handler := CreateHandler()
	confFiles := CreateConfig(GetUsername(), encryptionId)
	handler.CreateCustomPasswordStore(name, confFiles)
}

func ReturnPassStore(name string) (*passwordstoreFilesystem.PasswordStoreDir, error) {
	handler := CreateHandler()
	storePath := filepath.Join(handler.GetPath(), name)
	exists := api.CheckIfDirectoryExists(storePath)
	if exists {
		store := handler.ReadPasswordStore(filepath.Join(handler.GetPath(), name))
		return &store, nil
	}
	return nil, errors.New("the password store does not exist")
}

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
	return retriveContentOutOfContentDir(*content, identifier), nil
}
