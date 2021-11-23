package lang

import "fmt"

type GolangProject Project

func (p GolangProject) CheckNamingConventions() error {
	fmt.Printf("Checking naming conventions for project %s\n", p.Name)
	return nil
}

func (p GolangProject) Build() error {
	fmt.Printf("Building %s project %s\n", p.Language, p.Name)
	return nil
}
