package languages

import (
	"io/ioutil"
	"os"
)

const moduleLanguage string = "python"

var pythonProject Project

pythonProject[moduleLanguage] = {
	Name:              "setup.py",
	permissions:       0644,
	Content:           "setup.py's content here",
	ParentDir:         ".",
	CreateParentFunc:  os.Mkdir,
	CreateContentFunc: ioutil.WriteFile}

func init() {
	supportedLanguages.addLanguage(moduleLanguage)
	projectsMetaData.addProject(moduleLanguage, project)
}
