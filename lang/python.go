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

var pythonProject = Project{Name: "", Language: "python", ProjectItems: pythonProjectItems}

func init() {
	RegisterLanguage(python, &pythonProject)
}

func (p *PythonProject) CheckNamingConventions() error {
	fmt.Printf("Checking naming conventions for project %s\n", p.Name)
	return nil
}

func (p *PythonProject) Build() (err error) {

	for _, pItem := range pythonProjectItems {

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
		return fmt.Errorf("Could not create directory %s: %v\n", pItem.Name, err)
	}

	return nil
}

func createContent(name, Content string, Permission os.FileMode) error {
	fh, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("Could not create file %s: %v\n", name, err)
	}
	defer fh.Close()

	_, err = fh.WriteString(Content)
	if err != nil {
		return fmt.Errorf("Could not write to file %s: %v\n", name, err)
	}

	return nil
}
