package lang

import (
	"fmt"
)

type GolangProject Project

const golang = "golang"

var golangProjectItems = []ProjectItem{
	{Name: "README.md",
		Parent:     "",
		Permission: 0644,
		Content:    ReadmeContent},
}

var golangProject = Project{Name: "", Language: golang, ProjectItems: golangProjectItems}

func init() {
	RegisterLanguage(golang, &golangProject)
}

func (p *GolangProject) CheckNamingConventions() error {
	fmt.Printf("Checking naming conventions for project %s\n", p.Name)
	return nil
}

func (p *GolangProject) Build() error {
	fmt.Printf("Building %s project %s\n", p.Language, p.Name)
	return nil
}
