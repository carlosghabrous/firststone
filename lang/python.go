package lang

import "fmt"

type PythonProject Project

func (p PythonProject) CheckNamingConventions() error {
	fmt.Println("Checking Python naming conventions")
	return nil
}

func (p PythonProject) Build() error {
	fmt.Println("Building Python Project")
	return nil
}
