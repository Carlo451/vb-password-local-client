package keyring

import "github.com/Carlo451/vb-password-local-client/storefuncs"

const (
	keyRingIdent             = "vB-Passwords"
	masterPasswordIdentifier = keyRingIdent + "-MasterPassword"
)

func Init(masterPassword string) {
	user := storefuncs.GetUsername()
	savePassword(masterPasswordIdentifier, user, masterPassword)
}

func GetMasterPassword() (string, error) {
	user := storefuncs.GetUsername()
	masterpassword, err := loadPassword(masterPasswordIdentifier, user)
	if err != nil {
		return "", err
	}
	return masterpassword, nil
}

func SaveEncryptionIdPassword(encryptionId, additionalPassword string) {
	completeIdent := buildIdentWithEncryptionId(encryptionId)
	savePassword(completeIdent, storefuncs.GetUsername(), additionalPassword)
}

func GetPasswordForEncryptionId(encryptionId string) (string, error) {
	completeIdent := buildIdentWithEncryptionId(encryptionId)
	masterpassword, err := loadPassword(completeIdent, storefuncs.GetUsername())
	if err != nil {
		return "", err
	}
	return masterpassword, nil
}

func DeleteEncryptionIdPassword(encryptionId string) {
	completeIdent := buildIdentWithEncryptionId(encryptionId)
	deletePassword(completeIdent, storefuncs.GetUsername())
}
