package tests

import (
	"github.com/Carlo451/vb-password-local-client/storefuncs"
	"testing"
)

func TestConfigRead(t *testing.T) {
	setUp()
	storefuncs.CreatePassStore("newTeststore", "1")
	ownerString := storefuncs.ReadConfigEntry(storefuncs.OwnerIdent, "newTeststore", storefuncs.CreateHandler())
	if ownerString != username {
		t.Errorf("ownerString does not match")
	}
	tearDown()
}

func TestWriteAndReadContentToStoreContentDirDoesNotExists(t *testing.T) {
	setUp()
	storefuncs.CreatePassStore("newTeststore", "1")
	realtiveContentPath := "subdirOne/subdirTwo/contentDir"
	storefuncs.WriteNewContentToStore("newTeststore", realtiveContentPath, "password", "password123")
	password, err := storefuncs.ReadContentFromStore("newTeststore", realtiveContentPath, "password")
	if err != nil {
		t.Errorf("Error reading content")
	}
	if password != "password123" {
		t.Errorf("password does not match")
	}
}

func TestWriteAndReadContentToStoreContentDirDoesExist(t *testing.T) {
	setUp()
	storefuncs.CreatePassStore("newTeststore", "1")
	realtiveContentPath := "subdirOne/subdirTwo/contentDir"
	storefuncs.WriteNewContentToStore("newTeststore", realtiveContentPath, "password", "password123")
	storefuncs.WriteNewContentToStore("newTeststore", realtiveContentPath, "username", "username123")
	password, err := storefuncs.ReadContentFromStore("newTeststore", realtiveContentPath, "username")
	if err != nil {
		t.Errorf("Error reading content")
	}
	if password != "username123" {
		t.Errorf("password does not match")
	}
}

func TestWriteAndReadContentToStoreOverwritesFile(t *testing.T) {
	setUp()
	storefuncs.CreatePassStore("newTeststore", "1")
	realtiveContentPath := "subdirOne/subdirTwo/contentDir"
	storefuncs.WriteNewContentToStore("newTeststore", realtiveContentPath, "password", "password123")
	storefuncs.WriteNewContentToStore("newTeststore", realtiveContentPath, "password", "username123")
	password, err := storefuncs.ReadContentFromStore("newTeststore", realtiveContentPath, "password")
	if err != nil {
		t.Errorf("Error reading content")
	}
	if password != "username123" {
		t.Errorf("password does not match")
	}
	tearDown()
}
