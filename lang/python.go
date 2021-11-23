package lang

import "fmt"

type PythonProject Project

var pythonProjectItems = []projectItem{
	{name: "item1", parent: "parent1", permission: "permission1", content: "content1"},
}

func (p PythonProject) CheckNamingConventions() error {
	fmt.Printf("Checking naming conventions for project %s\n", p.Name)
	return nil
}

func (p PythonProject) Build() error {
	fmt.Printf("Building %s project %s\n", p.Language, p.Name)
	for _, pItem := range pythonProjectItems {
		fmt.Println(pItem)
	}
	return nil
}
