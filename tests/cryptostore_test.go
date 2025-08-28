package tests

import (
	"testing"

	"github.com/Carlo451/vb-password-local-client/exec"
	"github.com/Carlo451/vb-password-local-client/storefuncs"
)

func TestChangePasswordOfEncryptionId(t *testing.T) {
	setUp()
	storefuncs.CreatePassStore("changPasswordStore", "main")
	exec.AddPasswordToStore("changPasswordStore", "test", "password123")
	passwordOldPassword, _ := exec.ReadPasswordFromStore("changPasswordStore", "test", masterpassword)

	storefuncs.ChangePasswordOfKeyPair("main", masterpassword, "hello123")
	passwordNewPassword, errorNewEncryptionIdPassword := exec.ReadPasswordFromStore("changPasswordStore", "test", "hello123")
	if errorNewEncryptionIdPassword != nil {
		t.Errorf("Error while reading content with new password")
	}
	if passwordOldPassword != passwordNewPassword {
		t.Error("something went wrong")
	}
	tearDown()
}

func TestCreateTempKeyPairAndOverwriteOldKeyPair(t *testing.T) {
	setUp()
	oldMasterKeyPair, _ := storefuncs.GetDecryptedKeyPair("main", masterpassword)
	storefuncs.CreateTempKeyPair("main", masterpassword)
	tempNewMasterKeyPair, _ := storefuncs.GetDecryptedKeyPair("main_new", masterpassword)
	storefuncs.OverwriteKeyPairWithTempKeyPair("main")
	newMasterKeyPair, _ := storefuncs.GetDecryptedKeyPair("main", masterpassword)
	backUpMasterKeyPair, _ := storefuncs.GetDecryptedKeyPair("main_old", masterpassword)

	if tempNewMasterKeyPair.PrivateKey != newMasterKeyPair.PrivateKey {
		t.Errorf("The temp private keys should match")
	}
	if oldMasterKeyPair.PrivateKey != backUpMasterKeyPair.PrivateKey {
		t.Errorf("The private keys of the backUp Old and the original key pairs should match")
	}
	if oldMasterKeyPair.PublicKey == newMasterKeyPair.PublicKey {
		t.Errorf("The old and new public keys should not match")
	}
	tearDown()
}

func TestChangeKeyPairAndReencryptContent(t *testing.T) {
	setUp()
	exec.CreatePasswordStoreWithMasterPassword("firstStore")
	exec.CreatePasswordStoreWithMasterPassword("secondStore")
	exec.CreatePasswordStoreWithMasterPassword("thirdStore")
	exec.AddPasswordToStore("firstStore", "test/youtube", "password123")
	exec.AddPasswordToStore("firstStore", "socialMedia", "password123")
	exec.AddPasswordToStore("firstStore", "test/amazon", "password123")
	exec.AddPasswordToStore("firstStore", "test/testTwo/youtube", "password123")
	exec.AddPasswordToStore("firstStore", "test/testTwo/testThree/amazon", "password123")
	exec.AddPasswordToStore("secondStore", "test/youtube", "password123")
	exec.AddPasswordToStore("secondStore", "socialMedia", "password123")
	exec.AddPasswordToStore("secondStore", "test/amazon", "password123")
	exec.AddPasswordToStore("thirdStore", "test/youtube", "password123")
	exec.AddPasswordToStore("thirdStore", "socialMedia", "password123")
	exec.AddPasswordToStore("thirdStore", "test/amazon", "password123")
	exec.ChangeKeyPairOfEncryptionId(masterpassword, storefuncs.GetMasterEncryptionId())

	one, _ := exec.ReadPasswordFromStore("firstStore", "test/youtube", masterpassword)
	two, _ := exec.ReadPasswordFromStore("firstStore", "socialMedia", masterpassword)
	three, _ := exec.ReadPasswordFromStore("firstStore", "test/amazon", masterpassword)

	four, _ := exec.ReadPasswordFromStore("secondStore", "test/youtube", masterpassword)
	five, _ := exec.ReadPasswordFromStore("secondStore", "socialMedia", masterpassword)
	six, _ := exec.ReadPasswordFromStore("secondStore", "test/amazon", masterpassword)
	seven, _ := exec.ReadPasswordFromStore("thirdStore", "test/youtube", masterpassword)
	eight, _ := exec.ReadPasswordFromStore("thirdStore", "socialMedia", masterpassword)
	nine, _ := exec.ReadPasswordFromStore("thirdStore", "test/amazon", masterpassword)

	ten, _ := exec.ReadPasswordFromStore("firstStore", "test/testTwo/youtube", masterpassword)
	eleven, _ := exec.ReadPasswordFromStore("firstStore", "test/testTwo/testThree/amazon", masterpassword)

	passwordList := [11]string{one, two, three, four, five, six, seven, eight, nine, ten, eleven}

	for _, password := range passwordList {
		if password != "password123" {
			t.Errorf("The password should match")
		}
	}
	tearDown()
}
