package lang

import "fmt"

type GolangProject Project

func (p GolangProject) CheckNamingConventions() error {
	fmt.Println("Checking Golang naming conventions")

	return nil
}

func (p GolangProject) Build() error {
	fmt.Println("Building Golang Project")

	return nil
}
