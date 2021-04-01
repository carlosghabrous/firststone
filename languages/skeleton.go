// TODO: change file's name
// TODO: error handling
package languages

import (
	"os"
	"path"
)

// void is an empty struct, used to implement a set type using a map
//TODO: not necessary
type void struct{}

// languageSet is used to implement a set
//TODO: not necessary
type languagesSet map[string]void

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

// Project is a collection of projectItems, addressed by a name
type Project map[string]projectItem

// Maps languages to Projects
//TODO: change name
type Projects map[string]Project

// emptyValue is a var of type void(struct{}) used in the set implementation
//TODO: not necessary
var emptyValue void

// supportedLanguages contains the languages supported by firststone
//TODO: not necessary
var supportedLanguages languagesSet = make(languagesSet)

// projectsMetaData maps languages to Projects (collection of projectItems)
//TODO: change name
var projectsMetaData Projects = make(Projects)

// addLanguage adds a language to the supportedLanguages set. It is used from the individual languages' modules
func (ls languagesSet) addLanguage(language string) {

	if _, ok := ls[language]; !ok {
		ls[language] = emptyValue
	}
}

// addProject is used to add a correspondance between a language and a Project(collection of projectItems)
func (pMetaData Projects) addProject(language string, project Project) {
	pMetaData[language] = project
}

// IsSupportedLanguage checks whether a language is supported by the project
//TODO: not necessary. CreateProject does a check on languages too. Could use it to check that a language is supported
func IsSupportedLanguage(language string) bool {
	_, ok := supportedLanguages[language]
	if !ok {
		return false
	}

	return true
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

	// case "go":
	// 	goProjectName.SetProjectName(name)

	default:
		break
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
