package lang

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func buildProject(projectItems *[]ProjectItem, replacer *strings.Replacer) (err error) {
	for _, pItem := range *projectItems {

		pItem.Name = (*replacer).Replace(pItem.Name)
		pItem.Parent = (*replacer).Replace(pItem.Parent)
		pItem.Content = (*replacer).Replace(pItem.Content)

		if _, err = os.Stat(pItem.Parent); os.IsNotExist(err) {
			err = createDir(&pItem)
			if err != nil {
				fmt.Errorf("could not create directory %s\n", pItem.Parent)
			}
		}

		err = createContent(&pItem)

		if err != nil {
			return err
		}
	}

	return nil
}

func createDir(pItem *ProjectItem) error {
	if err := os.Mkdir(pItem.Parent, os.ModeDir|0755); err != nil {
		return fmt.Errorf("could not create directory %s: %v", pItem.Name, err)
	}

	return nil
}

func createContent(pItem *ProjectItem) error {
	fh, err := os.Create(filepath.Join(pItem.Parent, pItem.Name))
	if err != nil {
		return fmt.Errorf("could not create file %s: %v", pItem.Name, err)
	}
	defer fh.Close()

	_, err = fh.WriteString(pItem.Content)
	if err != nil {
		return fmt.Errorf("could not write to file %s: %v", pItem.Name, err)
	}

	return nil
}
