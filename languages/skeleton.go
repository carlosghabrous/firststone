package languages

import (
	"fmt"
	"os"
	"path"
)

type void struct{}

type languagesSet map[string]void

//TODO: fields don't need to be public
type projectItem struct {
	Name              string
	Permissions       os.FileMode
	Content           string
	ParentDir         string
	CreateParentFunc  func(itemName string, perm os.FileMode) error              //TODO: replace by interface
	CreateContentFunc func(itemName string, data []byte, perm os.FileMode) error //TODO: replace by interface
}

// Maps project item name to projectItem
type Project map[string]projectItem

// Maps languages to Projects
type Projects map[string]Project

var member void

var supportedLanguages languagesSet = make(languagesSet)

var projectsMetaData Projects = make(Projects)

func (pMetaData Projects) addProject(language string, project Project) {
	fmt.Printf("Adding project %v for language %v\n", project, language)
	pMetaData[language] = project

	// if _, ok := pMetaData[language]; !ok {
	// }
}

func (ls languagesSet) addLanguage(language string) {

	if _, ok := ls[language]; !ok {
		ls[language] = member
		fmt.Printf("Language %v added\n", language)
	}
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
func CreateProject(name, language string) error {

	switch language {
	case "python":
		pythonProjectMeta.SetProjectMeta(name)

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

func dirExists(directory string) bool {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return false
	}

	return true
}

// func addProjectItem(language string, projectItem projectItem) error {
// 	_, ok := projectsMetaData[language]

// 	if !ok {
// 		projectsMetaData[language][projectItem.Name] = projectItem
// 		return nil
// 	}

// 	item, ok := projectsMetaData[language][projectItem.Name]
// 	if !ok {
// 		projectsMetaData[language][projectItem.Name] = item
// 		return nil
// 	}

// 	return fmt.Errorf("Project item %v already exists!\n", item.Name)
// }
