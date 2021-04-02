package languages

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// moduleLanguage contains the programming language of projects that will be created
const moduleLanguage string = "python"

// projectMetaData contains project meta data (duh)
type projectMetaData struct {
	projectName string // Project's name
	author      string // Project's author
}

// pythonProjectMetaData is a variable of type projectMetaData
var pythonProjectMetaData projectMetaData

// setProjectMetaData sets data necessary for the project
// TODO: set author, what else?
// TODO: author should always be me. Implement default value for parameter author
func (pn *projectMetaData) setProjectMetaData(name string) {
	pn.projectName = name
}

// init registers that this module's language is available
func init() {
	supportedLanguages.addLanguage(moduleLanguage)
}

// buildProject constructs a variable of type Project with all necessary projectItems
// TODO: CreateParentFunc and CreateContentFunc should contain these functions by default, instead of repeating them every time
func buildProject() {
	pythonProject := Project{
		"setup": projectItem{
			Name:              "setup.py",
			Permissions:       0644,
			Content:           setupContent(),
			ParentDir:         ".",
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile},

		"readme": projectItem{
			Name:              "README.md",
			Permissions:       0644,
			Content:           readMeContent(),
			ParentDir:         ".",
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile},

		"license": projectItem{
			Name:              "LICENSE",
			Permissions:       0644,
			Content:           "",
			ParentDir:         ".",
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile},

		"init": projectItem{
			Name:              "__init__.py",
			Permissions:       0644,
			Content:           initPyContent(),
			ParentDir:         pythonProjectMetaData.projectName,
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile,
		},

		"initTest": projectItem{
			Name:              "__init__.py",
			Permissions:       0644,
			Content:           "",
			ParentDir:         path.Join(pythonProjectMetaData.projectName, "tests"),
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile,
		},

		"testProject": projectItem{
			Name:              "test_" + pythonProjectMetaData.projectName + ".py",
			Permissions:       0644,
			Content:           testProjectContent(),
			ParentDir:         path.Join(pythonProjectMetaData.projectName, "tests"),
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile,
		},
	}

	projectsMetaData.addProject(moduleLanguage, pythonProject)

}

func setupContent() string {
	content := "this is the project's " + pythonProjectMetaData.projectName + "setup.py content"
	return content
}

func readMeContent() string {
	content := "this is the readme.md content"
	return content
}

func initPyContent() string {
	content := []string{"'''",
		"Documentation for the " + pythonProjectMetaData.projectName + " package",
		"'''",
		"__version__ = '0.0.1.dev0'",
	}
	return strings.Join(content, "\n")
}

func testProjectContent() string {
	content := []string{"'''",
		"High-level tests for the  package.",
		"'''",
		"import " + pythonProjectMetaData.projectName,
		"def test_version():",
		"\tassert " + pythonProjectMetaData.projectName + ".__version__ is not None",
	}

	return strings.Join(content, "\n")
}
