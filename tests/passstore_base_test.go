package tests

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/Carlo451/vb-password-local-client/storefuncs"
)

func TestBaseStoreInit(t *testing.T) {
	setUp()
	entrys, err := os.ReadDir(filepath.Join(basePath, "keystore"))
	if len(entrys) == 0 {
		t.Errorf("No entries in keystore directory -  should have entries")
	}
	if err != nil {
		t.Error(err)
	}
	tearDown()
}

func TestCreatePassStore(t *testing.T) {
	setUp()
	storefuncs.CreatePassStore("newTeststore", "1")
	infoo, err := os.Stat(filepath.Join(basePath, "newTeststore", "configs"))
	encryptionId, errEnc := os.Stat(filepath.Join(basePath, "newTeststore", "configs/encryptionId"))
	if err != nil || errEnc != nil {
		t.Error("Error creating configs")
	}
	if infoo.Size() < 2 {
		t.Errorf("New store created with wrong configs")
	}
	if encryptionId.Size() < 1 {
		t.Errorf("New store created with wrong encryptionId")
	}
	tearDown()
}

func TestCreateAndReturnPassStores(t *testing.T) {
	setUp()
	storeNames := []string{"newTeststoreOne", "newTeststoreTwo", "newTeststoreThree"}
	storefuncs.CreatePassStore(storeNames[0], "1")
	storefuncs.CreatePassStore(storeNames[1], "1")
	storefuncs.CreatePassStore(storeNames[2], "1")
	stores := storefuncs.GetAllPassStoreNames()
	if len(stores) != len(storeNames) {
		t.Errorf("Not all stores got created")
	}
	for _, store := range stores {
		if !slices.Contains(storeNames, store) {
			t.Errorf("THis store should not have been created")
		}
	}
	tearDown()
}
