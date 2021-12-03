package lang

import (
	"fmt"

	"github.com/carlosghabrous/firststone/registry"
)

const golang = "golang"

var golangProjectItems = []registry.ProjectItem{
	{Name: "README.md",
		Parent:     "",
		Permission: 0644,
		Content:    ReadmeContent},
}

var golangProject = registry.Project{Name: "", Language: "golang", ProjectItems: pythonProjectItems}

func init() {
	registry.RegisterLanguage(golang, &golangProject)
}

func (p *registry.Project) CheckNamingConventions() error {
	fmt.Printf("Checking naming conventions for project %s\n", p.Name)
	return nil
}

func (p *registry.Project) Build() error {
	fmt.Printf("Building %s project %s\n", p.Language, p.Name)
	return nil
}
