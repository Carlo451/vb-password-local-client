package storefuncs

import (
	"github.com/Carlo451/vb-password-base-package/cryptography/cryptographyoperations"
	"github.com/Carlo451/vb-password-base-package/cryptography/keygenerator"
	"github.com/Carlo451/vb-password-base-package/passwordstore/passwordstoreFilesystem"
	"os"
)

const dirName = "keystore"

func CreateCryptoStore(baseDirectory passwordstoreFilesystem.PasswordStoreDir, masterpassword string) (bool, error) {
	contentDir := passwordstoreFilesystem.NewEmptyContentDirecotry(baseDirectory, dirName)
	contentDir.WriteDirectory()
	WriteNewKeyPairs("main", masterpassword)
	_, err := os.Stat(contentDir.GetAbsoluteDirectoryPath())
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetKeyStore() (*passwordstoreFilesystem.PasswordStoreContentDir, error) {
	handler := CreateHandler()
	return handler.ReadContentDir(handler.GetPath(), dirName)
}

func WriteNewKeyPairs(encryptionId, passhphrase string) (bool, error) {
	keyStoreDir, err := GetKeyStore()
	if err != nil {
		return false, err
	}
	keyPair := keygenerator.GenerateAsymmetricKey()
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
	pubKey := retriveContentOutOfContentDir(*keyStoreDir, encryptionId+".pub")
	return cryptographyoperations.EncryptStringSymmetric(pubKey, clearContent)
}

func DecryptContentWithEncryptionIdAndPassword(encryptionId, encryptedContent, passphrase string) (string, error) {
	keyStoreDir, err := GetKeyStore()
	if err != nil {
		return "", err
	}
	privKeyEncrypted := retriveContentOutOfContentDir(*keyStoreDir, encryptionId+".priv.age")
	privClear, errDecrypt := cryptographyoperations.DecryptStringSymmetric(privKeyEncrypted, passphrase)
	if errDecrypt != nil {
		return "", errDecrypt
	}
	return cryptographyoperations.DecryptStringAsymmetric(encryptedContent, privClear)
}
