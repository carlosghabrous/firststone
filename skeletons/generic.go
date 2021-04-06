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

// ProjectCrap is a list of all items that must be cleaned
type ProjectCrap []string

// ProjectCleaner is a function that returns a list of items to be cleaned
type ProjectCleaner func(pName string) ProjectCrap

// ProjectCleanerRegistry maps languages to functions that perform the cleaning
type ProjectCleanerRegistry map[string]ProjectCleaner

// registry is a variable of type ProjectRegistry
// Associations between languages and project builder functions are stored here
var registry ProjectRegistry = make(ProjectRegistry)

// cleanRegistry is a variable of type ProjectCleanRegistry
// Associations between languages and project cleaner functions are stored here
var cleanRegistry ProjectCleanerRegistry = make(ProjectCleanerRegistry)

// registerBuilder maps a language to its builder function for later use
func registerBuilder(language string, builder ProjectBuilder) {
	if _, ok := registry[language]; !ok {
		registry[language] = builder
	}
}

func registerCleaner(language string, cleaner ProjectCleaner) {
	if _, ok := cleanRegistry[language]; !ok {
		cleanRegistry[language] = cleaner
	}
}

// func addToRegister(language string, registry map[string]interface{}, item interface{}) func() {
// 	return func() {
// 		if _, ok := registry[language]; !ok {
// 			registry[language] = item
// 		}
// 	}
// }

// CreateProject runs predefined actions to create a project of a certain language
func CreateProject(language string) error {

	var builder ProjectBuilder
	var ok bool

	if builder, ok = registry[language]; !ok {
		return fmt.Errorf("Language %v not supported\n", language)
	}

	name, err := getCurrentDirectory()
	if err != nil {
		return fmt.Errorf("Error while getting current directory: %v\n", err)
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

func CleanProject(language string) error {

	var cleaner ProjectCleaner
	var ok bool

	if cleaner, ok = cleanRegistry[language]; !ok {
		return fmt.Errorf("Language %v not supported\n", language)
	}

	name, err := getCurrentDirectory()
	if err != nil {
		return fmt.Errorf("Error while getting current directory: %v\n", err)
	}

	projectCrap := cleaner(name)

	for _, crap := range projectCrap {
		fmt.Printf("removing %v\n", crap)
		os.RemoveAll(crap)
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

func getCurrentDirectory() (cwd string, err error) {
	cdPath, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Could not get current directory's name\n")
	}

	cwd = path.Base(cdPath)

	return cwd, err
}
