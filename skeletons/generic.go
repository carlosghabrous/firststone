// TODO: error handling
package skeletons

import (
	"fmt"
	"os"
	"path"
)

// projectItem provides the description of each item (file or directory) that belong to a project
//TODO: members do not need to be public
type projectItem struct {
	Name              string                                                     // Item's name
	Permissions       os.FileMode                                                // Item's permissions (0644 for files, 0755 for directories)
	Content           string                                                     // Item's content
	ParentDir         string                                                     // Item's parent directory
	CreateParentFunc  func(itemName string, perm os.FileMode) error              // Function signature to create the item's parent
	CreateContentFunc func(itemName string, data []byte, perm os.FileMode) error // Function signature to create the item's content
}

// Project is a collection of projectItems
type Project []projectItem

// Maps languages to Projects
//TODO: change name
type Projects map[string]Project

// type projectMetaData struct {
// 	name      string
// 	author    string
// 	email     string
// 	url       string
// 	shortDesc string
// }
// type ProjectBuilder interface {
// 	setProjectMetaData(metaData)
// 	buildProject() error
// }

// projectsMetaData maps languages to Projects (collection of projectItems)
//TODO: change name
var projectsMetaData Projects = make(Projects)

// addLanguage adds a language to the projectsMetaData map. It is used from the individual languages' modules
func (p Projects) addLanguage(language string) {

	if _, ok := p[language]; !ok {
		p[language] = Project{}
	}
}

// addProject is used to add a correspondance between a language and a Project(collection of projectItems)
func (pMetaData Projects) addProject(language string, project Project) {
	pMetaData[language] = project
}

// CreateProject runs predefined actions to create a project of a certain language
// TODO: refactor
// 1. replace switch/case by a map of language a type. This type should contain two members, one the SetProjectMeta function and the other
// the buildProject function
func CreateProject(name, language string) error {

	switch language {
	case "python":
		pythonProjectMetaData.setProjectMetaData(name)

		break

	case "go":
		// goProjectMetaData.setProjectMetaData(name)
		break

	default:
		return fmt.Errorf("Language %v not supported\n", language)
	}

	buildProject()
	project := projectsMetaData[language]

	for _, projectItem := range project {
		if !dirExists(projectItem.ParentDir) {
			projectItem.CreateParentFunc(projectItem.ParentDir, 0755)
		}

		projectItem.CreateContentFunc(path.Join(projectItem.ParentDir, projectItem.Name), []byte(projectItem.Content), projectItem.Permissions)
	}

	return nil
}

// dirExists checks whether a directory exists already
func dirExists(directory string) bool {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return false
	}

	return true
}
