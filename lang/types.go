package lang

import "os"

// projectItem represents an item in a project
type ProjectItem struct {
	Name       string
	Parent     string //TODO: does go have a "directory"/file object
	Permission os.FileMode
	Content    string
}

// Project represents a project in a certain language
type Project struct {
	Name         string
	Language     string
	ProjectItems []ProjectItem
}

// ProjectBuilder is an interface that builder objects implement within the lang package
type ProjectBuilder interface {
	CheckNamingConventions(name string) error
	Build() error
}

// Projects type stores builders for every language
type Registry map[string]ProjectBuilder
