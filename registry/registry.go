package registry

import (
	"fmt"
	"os"
)

// projectItem represents an item in a project
type projectItem struct {
	name       string
	parent     string //TODO: does go have a "directory"/file object
	permission os.FileMode
	content    string
}

// Project represents a project in a certain language
type Project struct {
	Name         string
	Language     string
	projectItems []projectItem
}

// ProjectBuilder is an interface that builder objects implement within the lang package
type ProjectBuilder interface {
	CheckNamingConventions() error
	Build() error
}

// Projects type stores builders for every language
//TODO: replace ProjectBuilder by *Project
type Projects map[string]*Project

var languageRegistry Projects

// RegisterLanguage allows modules in package lang to register themselves
func RegisterLanguage(language string, project *Project) error {
	if languageRegistry == nil {
		languageRegistry = make(map[string]*Project)
	}

	_, ok := languageRegistry[language]
	if ok {
		return fmt.Errorf("Language %s already registered. Are you sure you want to overwrite it?\n", language)
	}

	languageRegistry[language] = project

	return nil
}

// LanguageSupported returns and error if argument language is not contained in the language registry
func LanguageSupported(language string) error {
	_, ok := languageRegistry[language]
	if !ok {
		return fmt.Errorf("Language %s is not supported\n", language)
	}
	return nil
}

func GetProject(language, name string) *Project {
	project := languageRegistry[language]
	project.Name = name
	return project
}
