package languages

import (
	"io/fs"
	"path/filepath"
)

type projectItem struct {
	fileName        string
	filePermissions fs.FileMode
	directoryName   string
	fileContent     string
}

type Project map[string]projectItem

func (project *Project) addItem(item projectItem) {
	itemKey := filepath.Join(projectItem.directoryName, projectItem.fileName)
	_, ok := project[itemKey]

	if !ok {
		project[itemKey] = projectItem
	}
}
