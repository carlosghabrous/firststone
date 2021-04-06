package skeletons

import (
	"fmt"
	"io/ioutil"
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

func TestIsCwdEmptyWithEmptyDir(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error while getting current directory\n")
	}

	dir, err := ioutil.TempDir(currentDir, "tempDir")
	if err != nil {
		fmt.Printf("Error while creating temp dir %v tempDir\n", dir)
	}
	defer os.RemoveAll(dir)

	if empty, _ := isCwdEmpty(dir); !empty {
		t.Errorf("Directory %v should be empty, but it is not\n", dir)
	}
}

func TestIsCwdEmptyWithNonEmptyDir(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error while getting current directory\n")
	}

	dir, err := ioutil.TempDir(currentDir, "tempDir")
	if err != nil {
		fmt.Printf("Error while creating temp dir %v tempDir\n", dir)
	}
	defer os.RemoveAll(dir)

	file, err := ioutil.TempFile(dir, "some_file")
	if err != nil {
		fmt.Printf("Error while creating file %v\n", file.Name())
	}
	defer os.Remove(file.Name())

	if empty, _ := isCwdEmpty(dir); empty {
		t.Errorf("Directory %v should NOT be empty, but it is\n", dir)
	}

}
