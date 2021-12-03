package lang

import (
	"os"
	"path"
)

const languageDirName = "lang"

var langDir = path.Join(".", languageDirName)
var supportedLangSlice []string
var supportedLangs = make(map[string]bool)
var filesToExclude map[string]bool

type projectItem struct {
	name       string
	parent     string // Maybe not a string
	permission os.FileMode
	content    string
}

type Project struct {
	Name         string
	Language     string
	projectItems []projectItem
}

type ProjectBuilder interface {
	CheckNamingConventions() error
	Build() error
	SetAppName(appName string)
}
