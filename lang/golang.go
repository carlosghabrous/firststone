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

var golangProject = GolangProject{Language: golang, ProjectItems: golangProjectItems}

func init() {
	RegisterLanguage(golang, &golangProject)
}

func (p *GolangProject) CheckNamingConventions(name string) error {
	p.Name = name
	fmt.Printf("Naming conventions for project %s OK\n", p.Name)
	return nil
}

func (p *GolangProject) Build() error {
	fmt.Printf("Building %s project %s\n", p.Language, p.Name)
	return nil
}
