package tests

import (
	"testing"

	"github.com/Carlo451/vb-password-local-client/exec"
	"github.com/Carlo451/vb-password-local-client/storefuncs"
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

func TestCreateAndReadPasswordInDeeperSubDir(t *testing.T) {
	setUp()
	storefuncs.CreatePassStore("testStore", "main")
	err := exec.AddPasswordToStore("testStore", "video/youtube", "personalYoutubePassword")
	if err != nil {
		t.Fatal(err)
	}
	password, readErr := exec.ReadPasswordFromStore("testStore", "video/youtube", masterpassword)
	if readErr != nil {
		t.Errorf("Error reading password: %s", readErr)
	}
	if password != "personalYoutubePassword" {
		t.Errorf("Error reading password")
	}
	tearDown()
}

func TestOverwriteAndReadPassword(t *testing.T) {
	setUp()
	storefuncs.CreatePassStore("testStore", "main")
	exec.AddPasswordToStore("testStore", "youtube", "personalYoutubePassword")
	err := exec.AddPasswordToStore("testStore", "youtube", "personalOverriddenYoutubePassword")
	if err != nil {
		t.Fatal(err)
	}
	password, readErr := exec.ReadPasswordFromStore("testStore", "youtube", masterpassword)
	if readErr != nil {
		t.Errorf("Error reading password: %s", readErr)
	}
	if password != "personalOverriddenYoutubePassword" {
		t.Errorf("Error reading password")
	}
	tearDown()
}
