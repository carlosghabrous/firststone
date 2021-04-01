package languages

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const moduleLanguage string = "python"

type projectMeta struct {
	ProjectName string
}

var pythonProjectMeta projectMeta

func (pn *projectMeta) SetProjectMeta(name string) {
	pn.ProjectName = name
}

func init() {
	supportedLanguages.addLanguage(moduleLanguage)
}

func buildProject() {
	pythonProject := Project{
		"setup": projectItem{
			Name:              "setup.py",
			Permissions:       0644,
			Content:           setupContent(),
			ParentDir:         ".",
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile},

		"README": projectItem{
			Name:              "README.md",
			Permissions:       0644,
			Content:           readMeContent(),
			ParentDir:         ".",
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile},

		"init": projectItem{
			Name:              "__init__.py",
			Permissions:       0644,
			Content:           initPyContent(),
			ParentDir:         pythonProjectMeta.ProjectName,
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile,
		},

		"initTest": projectItem{
			Name:              "__init__.py",
			Permissions:       0644,
			Content:           "",
			ParentDir:         path.Join(pythonProjectMeta.ProjectName, "tests"),
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile,
		},

		"testProject": projectItem{
			Name:              "test_" + pythonProjectMeta.ProjectName + ".py",
			Permissions:       0644,
			Content:           testProjectContent(),
			ParentDir:         path.Join(pythonProjectMeta.ProjectName, "tests"),
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile,
		},
	}

	projectsMetaData.addProject(moduleLanguage, pythonProject)

}

func setupContent() string {
	content := "this is the project's " + pythonProjectMeta.ProjectName + "setup.py content"
	return content
}

func readMeContent() string {
	content := "this is the readme.md content"
	return content
}

func initPyContent() string {
	content := []string{"'''",
		"Documentation for the " + pythonProjectMeta.ProjectName + " package",
		"'''",
		"__version__ = '0.0.1.dev0'",
	}
	return strings.Join(content, "\n")
}

func testProjectContent() string {
	content := []string{"'''",
		"High-level tests for the  package.",
		"'''",
		"import " + pythonProjectMeta.ProjectName,
		"def test_version():",
		"\tassert " + pythonProjectMeta.ProjectName + ".__version__ is not None",
	}

	return strings.Join(content, "\n")
}
