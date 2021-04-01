package languages

import (
	"fmt"
	"io/ioutil"
	"os"
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
	fmt.Printf("building project: %v\n", pythonProjectMeta.ProjectName)
	pythonProject := Project{
		"setup.py": projectItem{
			Name:              "setup.py",
			permissions:       0644,
			Content:           setupContent(),
			ParentDir:         ".",
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile},
	}

	projectsMetaData.addProject(moduleLanguage, pythonProject)

}
func setupContent() string {
	content := "this is the project's " + pythonProjectMeta.ProjectName + "setup.py content"
	return content
}
