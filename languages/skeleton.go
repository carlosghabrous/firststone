package languages

type void struct{}
type languagesSet map[string]void
type ProjectItem struct {
	name           string
	permissions    string
	content        string
	parentDir      string
	createFunc     func(itemName, itemPermissions, itemParentDir string) error
	addContentFunc func(itemName, itemContent string) error
}
type Project map[string]ProjectItem

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
	for _, projectItem := range projectItems {
		projectItem.createFunc(projectItem.name, projectItem.permissions, projectItem.parentDir)
		projectItem.addContentFunc(projectItem.name, projectItem.content)
	}

	return nil
}
