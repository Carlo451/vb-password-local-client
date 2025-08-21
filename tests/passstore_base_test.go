package tests

import (
	"github.com/Carlo451/vb-password-local-client/storefuncs"
	"os"
	"path/filepath"
	"testing"
)

const basePath = "/home/carl-moritz/vB-Passwords"
const masterpassword = "password123"

func setUp() {
	os.Setenv("VB_PASSWORD_STORE_BASE_DIR_PATH", basePath)
	storefuncs.Init(masterpassword)
}
func tearDown() {
	os.RemoveAll(basePath)
}

func TestBaseStoreInit(t *testing.T) {
	setUp()
	_, err := os.ReadDir(filepath.Join(basePath, "keystore"))
	if err != nil {
		t.Error(err)
	}
	tearDown()
}

func TestCreatePassStore(t *testing.T) {
	setUp()
	storefuncs.CreatePassStore("newTeststore", "1")
	infoo, err := os.Stat(filepath.Join(basePath, "newTeststore", "configs"))
	if err != nil {
		t.Error(err)
	}
	if infoo.Size() < 2 {
		t.Errorf("New store created with wrong configs")
	}
}
