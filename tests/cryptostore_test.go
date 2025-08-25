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
