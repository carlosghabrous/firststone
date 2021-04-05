// TODO: error handling
package skeletons

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// ProjectItem provides the description of each item (file or directory) that belong to a project
type ProjectItem struct {
	itemName    string      // Item's name
	permissions os.FileMode // Item's permissions (0644 for files, 0755 for directories)
	content     string      // Item's content
	parentDir   string      // Item's parent directory
}

// Project is a collection of projectItems
type Project []ProjectItem

// ProjectMetaData contains some fields every project should have
type ProjectMetaData struct {
	pName   string // Project's name
	pAuthor string // Project's main author
	pMail   string // Project's main author email
}

// ProjectBuilder is a function that takes care of building projects
// It is defined in the different language specific skeleton modules
type ProjectBuilder func(pMeta *ProjectMetaData) Project

// ProjectRegistry maps languages to project builder functions
type ProjectRegistry map[string]ProjectBuilder

// registry is a variable of type ProjectRegistry
// Associations between languages and project builder functions are stored here
var registry ProjectRegistry = make(ProjectRegistry)

// registerBuilder maps a language to its builder function for later use
func registerBuilder(language string, builder ProjectBuilder) {
	if _, ok := registry[language]; !ok {
		registry[language] = builder
	}
}

// CreateProject runs predefined actions to create a project of a certain language
func CreateProject(name, language string) error {

	var builder ProjectBuilder
	var ok bool

	if builder, ok = registry[language]; !ok {
		return fmt.Errorf("Language %v not supported\n", language)
	}

	project := builder(&ProjectMetaData{pName: name})

	for _, projectItem := range project {
		if !dirExists(projectItem.parentDir) {
			os.Mkdir(projectItem.parentDir, 0755)
		}

		ioutil.WriteFile(
			path.Join(projectItem.parentDir, projectItem.itemName),
			[]byte(projectItem.content),
			projectItem.permissions,
		)
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
