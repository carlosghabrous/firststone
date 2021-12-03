package lang

import (
	"os"
)

type projectItem struct {
	name       string
	parent     string // Maybe not a string
	permission os.FileMode
	content    string
}

type Project struct {
	Name         string
	Language     string
	projectItems []projectItem
}

type ProjectBuilder interface {
	CheckNamingConventions() error
	Build() error
}
