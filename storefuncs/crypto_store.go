package storefuncs

import (
	"errors"
	"os"

	"github.com/Carlo451/vb-password-base-package/cryptography/cryptographyoperations"
	"github.com/Carlo451/vb-password-base-package/cryptography/keys"
	"github.com/Carlo451/vb-password-base-package/passwordstoreFilesystem"
)

const (
	masterEncryptionId = "main"
	dirName            = "keystore"
	pubKeyEnding       = ".pub"
	privKeyEnding      = ".priv.age"
)

func CreateCryptoStore(baseDirectory passwordstoreFilesystem.PasswordStoreDir, masterPassword string) (bool, error) {
	contentDir := passwordstoreFilesystem.NewEmptyContentDirecotry(baseDirectory, dirName)
	infoContent := passwordstoreFilesystem.NewPasswordstoreContentFile("THis directory holds the encrypted passPhrases for the stores", "INFO", contentDir)
	contentDir.AppendFile(infoContent)
	contentDir.WriteDirectory()
	WriteNewKeyPairs(masterEncryptionId, masterPassword)
	_, err := os.Stat(contentDir.GetAbsoluteDirectoryPath())
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetMasterEncryptionId() string {
	return masterEncryptionId
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

func GetEncryptedKeyPair(encryptionId string) (*keys.AsymmetricKeyPair, error) {
	keyStore, _ := GetKeyStore()
	keyExists := keyStore.LookUpFile(encryptionId + pubKeyEnding)
	if !keyExists {
		return nil, errors.New("keyPair not found")
	}
	publicKeyFile, _ := keyStore.ReturnFile(encryptionId + pubKeyEnding)
	privateKeyFile, _ := keyStore.ReturnFile(encryptionId + privKeyEnding)
	keyPair := keys.NewAsymmetricKeyPair(publicKeyFile.GetContent(), privateKeyFile.GetContent())
	return &keyPair, nil
}

func GetDecryptedKeyPair(encryptionId, passphrase string) (*keys.AsymmetricKeyPair, error) {
	keyPair, error := GetEncryptedKeyPair(encryptionId)
	if error != nil {
		return nil, error
	}
	publicKey := keyPair.PublicKey
	encryptedPrivateKey := keyPair.PrivateKey
	decryptedPrivateKey, decryptError := cryptographyoperations.DecryptStringSymmetric(encryptedPrivateKey, passphrase)
	if decryptError != nil {
		return nil, decryptError
	}
	decryptedKeyPair := keys.NewAsymmetricKeyPair(publicKey, decryptedPrivateKey)
	return &decryptedKeyPair, nil
}

func WriteNewKeyPairs(encryptionId, passphrase string) (bool, error) {
	keyStoreDir, err := GetKeyStore()
	if err != nil {
		return false, err
	}
	keyPair := keys.GenerateAsymmetricKey()
	publicKey := keyPair.PublicKey
	encryptedPrivKey, encryptionErr := cryptographyoperations.EncryptStringSymmetric(keyPair.PrivateKey, passphrase)
	if encryptionErr != nil {
		return false, encryptionErr
	}
	pubKeyFile := passwordstoreFilesystem.NewPasswordstoreContentFile(publicKey, encryptionId+pubKeyEnding, *keyStoreDir)
	privateKeyFile := passwordstoreFilesystem.NewPasswordstoreContentFile(encryptedPrivKey, encryptionId+privKeyEnding, *keyStoreDir)
	keyStoreDir.AppendFiles(pubKeyFile, privateKeyFile)
	keyStoreDir.WriteFiles()
	return true, nil
}

func EncryptContentWithEncryptionId(encryptionId, clearContent string) (string, error) {
	keyStoreDir, err := GetKeyStore()
	if err != nil {
		return "", err
	}
	pubKey := retrieveContentOutOfContentDir(*keyStoreDir, encryptionId+pubKeyEnding)
	return cryptographyoperations.EncryptStringAsymmetric(clearContent, pubKey)
}

func DecryptContentWithEncryptionIdAndPassword(encryptionId, encryptedContent, passphrase string) (string, error) {
	keyStoreDir, err := GetKeyStore()
	if err != nil {
		return "", err
	}
	privKeyEncrypted := retrieveContentOutOfContentDir(*keyStoreDir, encryptionId+privKeyEnding)
	privClear, errDecrypt := cryptographyoperations.DecryptStringSymmetric(privKeyEncrypted, passphrase)
	if errDecrypt != nil {
		return "", errDecrypt
	}
	return cryptographyoperations.DecryptStringAsymmetric(encryptedContent, privClear)
}

func CreateTempKeyPair(encryptionId, passphrase string) error {
	valid, err := WriteNewKeyPairs(encryptionId+"_new", passphrase)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("error creating new keyPair")
	}
	return nil
}

func OverwriteKeyPairWithTempKeyPair(encryptionId string) {
	keyStore, _ := GetKeyStore()
	pubKeyFile, _ := keyStore.ReturnFile(encryptionId + pubKeyEnding)
	privateKeyFile, _ := keyStore.ReturnFile(encryptionId + privKeyEnding)

	newPubKeyFile, _ := keyStore.ReturnFile(encryptionId + "_new" + pubKeyEnding)
	newPrivateKeyFile, _ := keyStore.ReturnFile(encryptionId + "_new" + privKeyEnding)

	backUpPubKey := passwordstoreFilesystem.NewPasswordstoreContentFile(pubKeyFile.GetContent(), encryptionId+"_old"+pubKeyEnding, *keyStore)
	backUpPrivKey := passwordstoreFilesystem.NewPasswordstoreContentFile(privateKeyFile.GetContent(), encryptionId+"_old"+privKeyEnding, *keyStore)
	keyStore.AppendFile(backUpPubKey)
	keyStore.AppendFile(backUpPrivKey)
	pubKeyFile.SetContent(newPubKeyFile.GetContent())
	privateKeyFile.SetContent(newPrivateKeyFile.GetContent())
	keyStore.WriteFiles()
	keyStore.RemoveFile(newPubKeyFile)
	keyStore.RemoveFile(newPrivateKeyFile)

}

func EncryptContentWithTempEncryptionId(encryptionId, content string) (string, error) {
	tempEncryptionId := encryptionId + "_new"
	return EncryptContentWithEncryptionId(tempEncryptionId, content)
}

func ChangePasswordOfKeyPair(encryptionId, passphrase, newPassphrase string) error {
	keyStoreDir, err := GetKeyStore()
	if err != nil {
		return err
	}
	privateKeyFile, _ := keyStoreDir.ReturnFile(encryptionId + privKeyEnding)
	privClear, errDecrypt := cryptographyoperations.DecryptStringSymmetric(privateKeyFile.GetContent(), passphrase)
	if errDecrypt != nil {
		return errDecrypt
	}
	keyPair := keys.NewAsymmetricKeyPair(retrieveContentOutOfContentDir(*keyStoreDir, encryptionId+pubKeyEnding), privClear)
	valid, validationError := keyPair.CheckIfKeyPairIsValid()
	if validationError != nil {
		return validationError
	}
	if !valid {
		return errors.New("The KeyPair is not valid")
	}
	newEncryptedPrivateKey, enryptionErr := cryptographyoperations.EncryptStringSymmetric(privClear, newPassphrase)
	if enryptionErr != nil {
		return enryptionErr
	}
	privateKeyFile.SetContent(newEncryptedPrivateKey)
	keyStoreDir.WriteFiles()
	return nil
}
