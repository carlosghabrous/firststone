package lang

import "fmt"

const ReadmeContent = `SHORT DESCRIPTION OF PROJECT

You can use [Github-flavored Markdown](https://guides.github.com/features/mastering-markdown/)
to write your content.

## Purpose of this project
## Getting started
##%`

type PythonProject Project

//TODO: what to do with 'parent', 'permission'.
//TODO: Need to distinguish between file and directory?

var pythonProjectItems = []projectItem{
	{name: "README.md",
		parent:     "",
		permission: "permission1",
		content:    ReadmeContent},
	{name: "setup.py",
		parent:     "",
		permission: "",
		content:    "setup.py content"},
	{name: "",
		parent:     "",
		permission: "",
		content:    ""},
	{name: "__init__.py",
		parent:     "",
		permission: "",
		content:    ""},
	{name: "tests",
		parent:     "",
		permission: "",
		content:    "",
	},
	{name: "__init__.py",
		parent:     "tests",
		permission: "",
		content:    "",
	},
	{name: "test_",
		parent:     "tests",
		permission: "",
		content:    ""},
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
