package skeletons

import (
	"os"
	"testing"
)

func TestDirExistsFakeDir(t *testing.T) {
	fakeDir := "/some/fake/dir"
	if dirExists(fakeDir) {
		t.Errorf("dirExists(%s) = true; want false", fakeDir)
	}
}

func TestDirExistsRealDir(t *testing.T) {
	homeDir := os.Getenv("HOME")
	if !dirExists(homeDir) {
		t.Errorf("dirExists(%s) = false; want true", homeDir)
	}
}
