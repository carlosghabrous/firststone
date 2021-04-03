package skeletons

// pythonModuleLanguage contains the programming language of projects that will be created
const pythonModuleLanguage string = "python"

// projectMetaDataPython embeds skeleton's generic projectMetaData and adds other fields
// type projectMetaDataPython struct {
// 	projectMetaData
// 	email     string
// 	url       string
// 	shortDesc string
// }

// pythonProjectMetaData is a variable of type projectMetaData
// var pythonProjectMetaData projectMetaDataPython

// setProjectMetaData sets data necessary for the project
// func (pn *projectMetaDataPython) setProjectMetaData(pmdPython *projectMetaDataPython) {
// 	// pn.
// }

// init registers that this module's language is available
func init() {
	lanRegistry.registerLanguage(pythonModuleLanguage)
}

// buildProject constructs a variable of type Project with all necessary projectItems
// TODO: CreateParentFunc and CreateContentFunc should contain these functions by default, instead of repeating them every time
// func registerProject() {
// 	pythonProject := Project{
// 		projectItem{
// 			itemName:          "setup.py",
// 			permissions:       0644,
// 			content:           setupContent(),
// 			parentDir:         ".",
// 			createParentFunc:  os.Mkdir,
// 			createContentFunc: ioutil.WriteFile},

// 		projectItem{
// 			itemName:          "README.md",
// 			permissions:       0644,
// 			content:           readMeContent(),
// 			parentDir:         ".",
// 			createParentFunc:  os.Mkdir,
// 			createContentFunc: ioutil.WriteFile},

// 		projectItem{
// 			itemName:          "LICENSE",
// 			permissions:       0644,
// 			content:           "",
// 			parentDir:         ".",
// 			createParentFunc:  os.Mkdir,
// 			createContentFunc: ioutil.WriteFile},

// 		projectItem{
// 			itemName:          "__init__.py",
// 			permissions:       0644,
// 			content:           initPyContent(),
// 			parentDir:         pythonProjectMetaData.pName,
// 			createParentFunc:  os.Mkdir,
// 			createContentFunc: ioutil.WriteFile,
// 		},

// 		projectItem{
// 			itemName:          "__init__.py",
// 			permissions:       0644,
// 			content:           "",
// 			parentDir:         path.Join(pythonProjectMetaData.pName, "tests"),
// 			createParentFunc:  os.Mkdir,
// 			createContentFunc: ioutil.WriteFile,
// 		},

// 		projectItem{
// 			itemName:          "test_" + pythonProjectMetaData.pName + ".py",
// 			permissions:       0644,
// 			content:           testProjectContent(),
// 			parentDir:         path.Join(pythonProjectMetaData.pName, "tests"),
// 			createParentFunc:  os.Mkdir,
// 			createContentFunc: ioutil.WriteFile,
// 		},
// 	}

// 	registry.addProject(pythonModuleLanguage, pythonProject)
// }

// func setupContent() string {
// 	content := "this is the project's " + pythonProjectMetaData.pName + "setup.py content"
// 	return content
// }

// func readMeContent() string {
// 	content := "this is the readme.md content"
// 	return content
// }

// func initPyContent() string {
// 	content := []string{"'''",
// 		"Documentation for the " + pythonProjectMetaData.pName + " package",
// 		"'''",
// 		"__version__ = '0.0.1.dev0'",
// 	}
// 	return strings.Join(content, "\n")
// }

// func testProjectContent() string {
// 	content := []string{"'''",
// 		"High-level tests for the  package.",
// 		"'''",
// 		"import " + pythonProjectMetaData.pName,
// 		"def test_version():",
// 		"\tassert " + pythonProjectMetaData.pName + ".__version__ is not None",
// 	}

// 	return strings.Join(content, "\n")
// }
