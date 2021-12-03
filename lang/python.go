package lang

import (
	"fmt"
	"os"

	"github.com/carlosghabrous/firststone/skeleton"
)

const python = "python"

const flag = "_project-name_"

var ReadmeContent string = `# ` +
	flag +
	`
	SHORT DESCRIPTION OF PROJECT

You can use [Github-flavored Markdown](https://guides.github.com/features/mastering-markdown/)
to write your content.

## Purpose of this project
## Getting started
##%`

type PythonProject Project

var pythonProjectItems = []projectItem{
	{name: "README.md",
		parent:     "",
		permission: 0644,
		content:    ReadmeContent},
	{name: "setup.py",
		parent:     "",
		permission: 0644,
		content:    "setup.py content"},
	{name: "_project_name_",
		parent:     "",
		permission: os.ModeDir | 0755,
		content:    "test"},
	{name: "__init__.py",
		parent:     "",
		permission: 0644,
		content:    "test"},
	{name: "tests",
		parent:     "",
		permission: os.ModeDir | 0755,
		content:    "test",
	},
	{name: "__init__.py",
		parent:     "tests",
		permission: 0644,
		content:    "test",
	},
	{name: "test_",
		parent:     "tests",
		permission: 0644,
		content:    "test"},
}

func init() {
	skeleton.RegisterLanguage(python)
}

func (p *PythonProject) CheckNamingConventions() error {
	fmt.Printf("Checking naming conventions for project %s\n", p.Name)
	return nil
}

func (p *PythonProject) Build() (err error) {

	for _, pItem := range pythonProjectItems {

		if pItem.permission.IsDir() {
			err = createDir(&pItem)

		} else {
			err = createContent(pItem.name, pItem.content, pItem.permission)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func createDir(pItem *projectItem) error {
	if err := os.Mkdir(pItem.name, pItem.permission); err != nil {
		return fmt.Errorf("Could not create directory %s: %v\n", pItem.name, err)
	}

	return nil
}

func createContent(name, content string, permission os.FileMode) error {
	fh, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("Could not create file %s: %v\n", name, err)
	}
	defer fh.Close()

	_, err = fh.WriteString(content)
	if err != nil {
		return fmt.Errorf("Could not write to file %s: %v\n", name, err)
	}

	return nil
}
