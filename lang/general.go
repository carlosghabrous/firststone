package lang

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

const languageDirName = "lang"

var langDir = path.Join(".", languageDirName)
var supportedLangs = make(map[string]bool)
var filesToExclude map[string]bool

func init() {
	initFilesToExclude()

	files, err := ioutil.ReadDir(langDir)
	if err != nil {
		log.Fatal("Could not scan languages directory!")
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		if _, ok := filesToExclude[fileName]; ok {
			continue
		}

		if idx := strings.Index(fileName, ".go"); idx != -1 {
			supportedLangs[fileName[:idx]] = true
		}
	}
}

func initFilesToExclude() {
	filesToExclude = map[string]bool{
		"general.go": true,
	}
}

func LanguageSupported(language string) error {
	fmt.Println("Language is supported")
	return nil
}

func CheckNamingConventions(name, language string) string {
	return "something checking here"
}

func BuildProject(name, language string) error {
	fmt.Println("Build project flawlessly")
	return nil
}
