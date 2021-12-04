package lang

import (
	"fmt"
	"os"
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

func buildProject(projectItems *[]ProjectItem) (err error) {
	for _, pItem := range *projectItems {

		if pItem.Permission.IsDir() {
			err = createDir(&pItem)

		} else {
			err = createContent(pItem.Name, pItem.Content, pItem.Permission)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func createDir(pItem *ProjectItem) error {
	if err := os.Mkdir(pItem.Name, pItem.Permission); err != nil {
		return fmt.Errorf("could not create directory %s: %v", pItem.Name, err)
	}

	return nil
}

func createContent(name, Content string, Permission os.FileMode) error {
	fh, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("could not create file %s: %v", name, err)
	}
	defer fh.Close()

	_, err = fh.WriteString(Content)
	if err != nil {
		return fmt.Errorf("could not write to file %s: %v", name, err)
	}

	return nil
}
