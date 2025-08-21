package tests

import (
	"github.com/Carlo451/vb-password-local-client/exec"
	"github.com/Carlo451/vb-password-local-client/storefuncs"
	"testing"
)

func TestCreateAndReadPassword(t *testing.T) {
	setUp()
	storefuncs.CreatePassStore("testStore", "main")
	err := exec.AddPasswordToStore("testStore", "youtube", "personalYoutubePassword")
	if err != nil {
		t.Fatal(err)
	}
	password, readErr := exec.ReadPasswordFromStore("testStore", "youtube", masterpassword)
	if readErr != nil {
		t.Errorf("Error reading password: %s", readErr)
	}
	if password != "personalYoutubePassword" {
		t.Errorf("Error reading password")
	}
	tearDown()
}
