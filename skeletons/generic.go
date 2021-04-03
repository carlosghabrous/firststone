// TODO: error handling
package skeletons

import (
	"fmt"
	"os"
)

// voidMember is an empty struct, used to implement a set
type voidMember struct{}

// languageRegistry is a set
type languageRegistry map[string]voidMember

// languageRegisterer is an interface that defines a method to register a language
type languageRegisterer interface {
	registerLanguage(language string)
}

// projectItem provides the description of each item (file or directory) that belong to a project
type projectItem struct {
	itemName          string                                                     // Item's name
	permissions       os.FileMode                                                // Item's permissions (0644 for files, 0755 for directories)
	content           string                                                     // Item's content
	parentDir         string                                                     // Item's parent directory
	createParentFunc  func(itemName string, perm os.FileMode) error              // Function signature to create the item's parent
	createContentFunc func(itemName string, data []byte, perm os.FileMode) error // Function signature to create the item's content
}

// Project is a collection of projectItems
type Project []projectItem

// voidMember is a variable used in the assignment of new entries to the lanRegistry set
var void voidMember

// lanRegistry is the variable holding the unique set of languages registered
var lanRegistry languageRegistry = make(languageRegistry)

func (lreg languageRegistry) registerLanguage(language string) {
	if _, ok := lreg[language]; !ok {
		lreg[language] = void
	}
}

// // projectMetaData contains some fields every project should have
// type projectMetaData struct {
// 	pName   string // Project's name
// 	pAuthor string // Project's main author
// }

// metaDataSetter is an interface with a method to set a project's meta data
// type metaDataSetter interface {
// 	setProjectMetaData(pmd *projectMetaData)
// }

// addProject is used to add a correspondance between a language and a Project(collection of projectItems)
// func (pMetaData ProjectRegistry) addProject(language string, project Project) {
// 	pMetaData[language] = project
// }

// CreateProject runs predefined actions to create a project of a certain language
// TODO: refactor
// 1. replace switch/case by a map of language a type. This type should contain two members, one the SetProjectMeta function and the other
// the buildProject function
func CreateProject(name, language string) error {

	switch language {
	case "python":
		// pythonProjectMetaData.setProjectMetaDatae})

		break

	case "go":
		// goProjectMetaData.setProjectMetaData(name)
		break

	default:
		return fmt.Errorf("Language %v not supported\n", language)
	}

	// registerProject()
	// project := registry[language]

	// for _, projectItem := range project {
	// 	if !dirExists(projectItem.parentDir) {
	// 		projectItem.createParentFunc(projectItem.parentDir, 0755)
	// 	}

	// 	projectItem.createContentFunc(path.Join(projectItem.parentDir, projectItem.name), []byte(projectItem.content), projectItem.permissions)
	// }

	return nil
}

// dirExists checks whether a directory exists already
func dirExists(directory string) bool {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return false
	}

	return true
}
