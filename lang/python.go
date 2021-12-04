package lang

import (
	"fmt"
	"os"
)

type PythonProject Project

const python = "python"

var ReadmeContent string = `#
	SHORT DESCRIPTION OF PROJECT

You can use [Github-flavored Markdown](https://guides.github.com/features/mastering-markdown/)
to write your Content.

## Purpose of this project
## Getting started
##%`

var pythonProjectItems = []ProjectItem{
	{Name: "README.md",
		Parent:     "",
		Permission: 0644,
		Content:    ReadmeContent},
	{Name: "setup.py",
		Parent:     "",
		Permission: 0644,
		Content:    "setup.py Content"},
	{Name: "_project_name_",
		Parent:     "",
		Permission: os.ModeDir | 0755,
		Content:    "test"},
	{Name: "__init__.py",
		Parent:     "",
		Permission: 0644,
		Content:    "test"},
	{Name: "tests",
		Parent:     "",
		Permission: os.ModeDir | 0755,
		Content:    "test",
	},
	{Name: "__init__.py",
		Parent:     "tests",
		Permission: 0644,
		Content:    "test",
	},
	{Name: "test_",
		Parent:     "tests",
		Permission: 0644,
		Content:    "test"},
}

var pythonProject = PythonProject{Language: python, ProjectItems: pythonProjectItems}

func init() {
	RegisterLanguage(python, &pythonProject)
}

func (p *PythonProject) CheckNamingConventions(name string) error {
	//TODO: after checking conventions are ok, assign name and return nil
	p.Name = name
	fmt.Printf("Naming conventions for project %s OK\n", p.Name)
	return nil
}

func (p *PythonProject) Build() (err error) {
	return buildProject(&pythonProjectItems)
}
