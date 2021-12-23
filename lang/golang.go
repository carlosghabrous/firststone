package lang

import (
	"fmt"
	"path/filepath"
	"strings"
)

type GolangProject Project

const (
	golang      = "golang"
	projectName = "PROJECTNAME"
)

const mainContent = `package main 
func main() {
}`

var golangProjectItems = []ProjectItem{
	{Name: "README.md",
		Parent:     ".",
		Permission: 0644,
		Content:    ""},
	{Name: "main.go",
		Parent:     filepath.Join(".", "cmd", projectName+"-cli"),
		Permission: 0644,
		Content:    mainContent},
	{Name: "file01.go",
		Parent:     filepath.Join(".", "internal", "package01"),
		Permission: 0644,
		Content:    ""},
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
	replacer := strings.NewReplacer(projectName, p.Name)
	return buildProject(&golangProjectItems, replacer)
}
