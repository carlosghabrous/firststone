package lang

import (
	"fmt"
)

var languageRegistry Registry

// RegisterLanguage allows modules in package lang to register themselves
func RegisterLanguage(language string, project ProjectBuilder) error {
	if languageRegistry == nil {
		languageRegistry = make(map[string]ProjectBuilder)
	}

	_, ok := languageRegistry[language]
	if ok {
		return fmt.Errorf("language %s already registered. Are you sure you want to overwrite it?", language)
	}

	languageRegistry[language] = project

	return nil
}

// LanguageSupported returns and error if argument language is not contained in the language registry
func LanguageSupported(language string) error {
	_, ok := languageRegistry[language]
	if !ok {
		return fmt.Errorf("language %s is not supported", language)
	}
	return nil
}

func GetProject(language string) ProjectBuilder {
	project := languageRegistry[language]
	return project
}
