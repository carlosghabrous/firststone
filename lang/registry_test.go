package lang

import "testing"

const language = "some_language"

var someLanguageProject = Project{}

func cleanUpRegistry() func() {
	return func() {
		for k := range languageRegistry {
			delete(languageRegistry, k)
		}
	}
}

func TestLanguageRegistryIsInitializedToNil(t *testing.T) {
	if languageRegistry != nil {
		t.Fatalf("Variable %s should not be allocated yet!\n", "languageRegistry")
	}
}

func TestRegisterLanguage(t *testing.T) {
	err := RegisterLanguage(language, &someLanguageProject)

	cleanup := cleanUpRegistry()
	defer cleanup()

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAddSameLanguageTwiceFails(t *testing.T) {
	err := RegisterLanguage(language, &someLanguageProject)
	cleanup := cleanUpRegistry()
	defer cleanup()

	if err != nil {
		t.Fatalf(err.Error())
	}

	err = RegisterLanguage(language, &someLanguageProject)
	if err == nil {
		t.Fatalf(err.Error())
	}

}

func TestLanguageSupported(t *testing.T) {
	_ = RegisterLanguage(language, &someLanguageProject)
	cleanup := cleanUpRegistry()
	defer cleanup()

	err := LanguageSupported(language)
	if err != nil {
		t.Fatalf("Language %s has already been registered (%e)\n", language, err)
	}
}

func TestGetProject(t *testing.T) {
	projectName := "a_name"
	_ = RegisterLanguage(language, &someLanguageProject)
	cleanup := cleanUpRegistry()
	defer cleanup()

	project := GetProject(language, projectName)

	if project.Name != projectName {
		t.Fatalf("Project name not correctly set!")
	}
}
