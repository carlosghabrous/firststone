package languages

import (
	"fmt"
	"io/ioutil"
	"os"
)

const moduleLanguage string = "python"

func init() {
	fmt.Println("Executing pythonSkeleton function ")
	addProjectItem(
		moduleLanguage,
		projectItem{
			Name:              "setup.py",
			permissions:       0644,
			Content:           "setup.py's content here",
			ParentDir:         ".",
			CreateParentFunc:  os.Mkdir,
			CreateContentFunc: ioutil.WriteFile},
	)
}
