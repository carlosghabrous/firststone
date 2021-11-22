package lang

import "fmt"

func LanguageSupported(language string) error {
	fmt.Println("Language is supported")
	return nil
}

func CheckNamingConventions(name, language string) string {
	return "something checking here"
}

func BuildProject(name, language string) error {
	fmt.Println("Build project flawlessly")
	return nil
}
