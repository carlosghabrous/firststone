package registry

import "fmt"

type SupportedLanguages map[string]bool

var languageRegistry SupportedLanguages

// RegisterLanguage allows modules in package lang to register themselves
func RegisterLanguage(language string) {
	if languageRegistry == nil {
		languageRegistry = make(map[string]bool)
	}

	languageRegistry[language] = true
}

// LanguageSupported returns and error if argument language is not contained in the language registry
func LanguageSupported(language string) error {
	_, ok := languageRegistry[language]
	if !ok {
		return fmt.Errorf("Language %s is not supported\n", language)
	}
	return nil
}
