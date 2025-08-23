package storefuncs

import (
	"errors"
	"os"

	"github.com/Carlo451/vb-password-base-package/cryptography/cryptographyoperations"
	"github.com/Carlo451/vb-password-base-package/cryptography/keys"
	"github.com/Carlo451/vb-password-base-package/passwordstore/passwordstoreFilesystem"
)

const dirName = "keystore"

func CreateCryptoStore(baseDirectory passwordstoreFilesystem.PasswordStoreDir, masterPassword string) (bool, error) {
	contentDir := passwordstoreFilesystem.NewEmptyContentDirecotry(baseDirectory, dirName)
	infoContent := passwordstoreFilesystem.NewPasswordstoreContentFile("THis directory holds the encrypted passPhrases for the stores", "INFO", contentDir)
	contentDir.AppendFile(infoContent)
	contentDir.WriteDirectory()
	WriteNewKeyPairs("main", masterPassword)
	_, err := os.Stat(contentDir.GetAbsoluteDirectoryPath())
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetKeyStore() (*passwordstoreFilesystem.PasswordStoreContentDir, error) {
	handler := CreateHandler()
	store := handler.ReadPasswordStore("")
	contentDirs := store.GetContentDirectories()
	for _, contentDir := range contentDirs {
		if contentDir.GetDirName() == dirName {
			return &contentDir, nil
		}
	}
	return nil, errors.New("keystore not found")
}

func WriteNewKeyPairs(encryptionId, passhphrase string) (bool, error) {
	keyStoreDir, err := GetKeyStore()
	if err != nil {
		return false, err
	}
	keyPair := keys.GenerateAsymmetricKey()
	publicKey := keyPair.PublicKey
	encryptedPrivKey, encryptionErr := cryptographyoperations.EncryptStringSymmetric(keyPair.PrivateKey, passhphrase)
	if encryptionErr != nil {
		return false, encryptionErr
	}
	pubKeyFile := passwordstoreFilesystem.NewPasswordstoreContentFile(publicKey, encryptionId+".pub", *keyStoreDir)
	privateKeyFile := passwordstoreFilesystem.NewPasswordstoreContentFile(encryptedPrivKey, encryptionId+".priv.age", *keyStoreDir)
	keyStoreDir.AppendFiles(pubKeyFile, privateKeyFile)
	keyStoreDir.WriteFiles()
	return true, nil
}

func EncryptContentWithEncrypptionId(encryptionId, clearContent string) (string, error) {
	keyStoreDir, err := GetKeyStore()
	if err != nil {
		return "", err
	}
	pubKey := retrieveContentOutOfContentDir(*keyStoreDir, encryptionId+".pub")
	return cryptographyoperations.EncryptStringAsymmetric(clearContent, pubKey)
}

func DecryptContentWithEncryptionIdAndPassword(encryptionId, encryptedContent, passphrase string) (string, error) {
	keyStoreDir, err := GetKeyStore()
	if err != nil {
		return "", err
	}
	privKeyEncrypted := retrieveContentOutOfContentDir(*keyStoreDir, encryptionId+".priv.age")
	privClear, errDecrypt := cryptographyoperations.DecryptStringSymmetric(privKeyEncrypted, passphrase)
	if errDecrypt != nil {
		return "", errDecrypt
	}
	return cryptographyoperations.DecryptStringAsymmetric(encryptedContent, privClear)
}
