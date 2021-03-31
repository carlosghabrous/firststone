package languages

import (
	"fmt"
	"os"
)

type void struct{}
type languagesSet map[string]void
type projectItem struct {
	Name              string
	permissions       os.FileMode
	Content           string
	ParentDir         string
	CreateParentFunc  func(itemName string, perm os.FileMode) error              //TODO: replace by interface
	CreateContentFunc func(itemName string, data []byte, perm os.FileMode) error //TODO: replace by interface
}
type Project map[string]projectItem

var member void
var supportedLanguages languagesSet = make(languagesSet)
var languageProjectItems = make(map[string]Project)

func init() {
	supportedLanguages["python"] = member
	supportedLanguages["go"] = member
}

// Checks whether a language is supported by the project
func IsSupportedLanguage(language string) bool {
	_, ok := supportedLanguages[language]
	if !ok {
		return false
	}

	return true
}

// Runs predefined actions to create a project in a certain language
func CreateProject(language string) error {
	projectItems := languageProjectItems[language]
	fmt.Printf("projectItems contains %v\n", projectItems)

	for _, projectItem := range projectItems {
		projectItem.CreateParentFunc(projectItem.Name, projectItem.permissions)
		projectItem.CreateContentFunc(projectItem.Name, []byte(projectItem.Content), projectItem.permissions)
	}

	return nil
}

func addProjectItem(language string, projectItem projectItem) error {
	_, ok := languageProjectItems[language]

	if !ok {
		languageProjectItems[language][projectItem.Name] = projectItem
		return nil
	}

	item, ok := languageProjectItems[language][projectItem.Name]
	if !ok {
		languageProjectItems[language][projectItem.Name] = item
		return nil
	}

	return fmt.Errorf("Project item %v already exists!\n", item.Name)
}
