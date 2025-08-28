package tests

import (
	"os"

	"github.com/Carlo451/vb-password-base-package/environment"
	"github.com/Carlo451/vb-password-local-client/storefuncs"
)

const basePath = "/home/carl-moritz/vB-Passwords"
const masterpassword = "password123"
const username = "carl-moritz"

func setUp() {
	_, found := environment.LookUpAndGetEnvValue("VB_PASSWORD_STORE_BASE_DIR_PATH")
	if !found {
		os.Setenv("VB_PASSWORD_STORE_BASE_DIR_PATH", basePath)
	}
	storefuncs.Init(masterpassword)
}
func tearDown() {
	os.RemoveAll(basePath)
}
