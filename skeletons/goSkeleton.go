package skeletons

import "strings"

// goModuleLanguage contains the programming language of projects that will be created
const goModuleLanguage string = "go"

// goProjectMetaData is a variable of type projectMetaData
var goProjectMetaData ProjectMetaData

// init registers that this module's language is available
func init() {
	registerBuilder(goModuleLanguage, buildProjectGo)
}

// buildProjectGo constructs a variable of type Project with all necessary projectItems
func buildProjectGo(pMeta *ProjectMetaData) Project {

	goProjectMetaData.pName = pMeta.pName

	if pMeta.pAuthor == "" {
		goProjectMetaData.pAuthor = "Carlos Ghabrous Larrea"
	}

	if pMeta.pMail == "" {
		goProjectMetaData.pMail = "carlos.ghabrous@gmail.com"
	}

	goProject := Project{
		ProjectItem{
			itemName:    "main.go",
			permissions: 0644,
			content:     mainContent(),
			parentDir:   ".",
		},

		ProjectItem{
			itemName:    "README.md",
			permissions: 0644,
			content:     readMeContent(),
			parentDir:   ".",
		},

		ProjectItem{
			itemName:    "Dockerfile",
			permissions: 0644,
			content:     dockerFileContent(),
			parentDir:   ".",
		},
	}

	return goProject
}

func mainContent() string {

	content := []string{
		"// " + goProjectMetaData.pName + "brief description",
		"package main\n",
		"func main(){",
		"}",
	}

	return strings.Join(content, "\n")
}

func dockerFileContent() string {

	content := []string{
		"FROM golang:latest\n",
		"WORKDIR /src\n",
		"ADD . /src",
	}

	return strings.Join(content, "\n")
}
