package skeletons

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// pythonModuleLanguage contains the programming language of projects that will be created
const pythonModuleLanguage string = "python"

// pythonProjectMetaData is a variable of type projectMetaData
var pythonProjectMetaData ProjectMetaData

// init registers that this module's language is available
func init() {
	registerBuilder(pythonModuleLanguage, buildProjectPython)
}

// buildProject constructs a variable of type Project with all necessary projectItems
// TODO: CreateParentFunc and CreateContentFunc should contain these functions by default, instead of repeating them every time
func buildProjectPython(pMeta *ProjectMetaData) Project {

	pythonProjectMetaData.pName = pMeta.pName

	if pMeta.pAuthor == "" {
		pythonProjectMetaData.pAuthor = "Carlos Ghabrous Larrea"
	}

	if pMeta.pMail == "" {
		pythonProjectMetaData.pMail = "carlos.ghabrous@gmail.com"
	}

	pythonProject := Project{
		ProjectItem{
			itemName:          "setup.py",
			permissions:       0644,
			content:           setupContent(),
			parentDir:         ".",
			createParentFunc:  os.Mkdir,
			createContentFunc: ioutil.WriteFile},

		ProjectItem{
			itemName:          "README.md",
			permissions:       0644,
			content:           readMeContent(),
			parentDir:         ".",
			createParentFunc:  os.Mkdir,
			createContentFunc: ioutil.WriteFile},

		ProjectItem{
			itemName:          "LICENSE",
			permissions:       0644,
			content:           "",
			parentDir:         ".",
			createParentFunc:  os.Mkdir,
			createContentFunc: ioutil.WriteFile},

		ProjectItem{
			itemName:          "__init__.py",
			permissions:       0644,
			content:           initPyContent(),
			parentDir:         pythonProjectMetaData.pName,
			createParentFunc:  os.Mkdir,
			createContentFunc: ioutil.WriteFile,
		},

		ProjectItem{
			itemName:          "__init__.py",
			permissions:       0644,
			content:           "",
			parentDir:         path.Join(pythonProjectMetaData.pName, "tests"),
			createParentFunc:  os.Mkdir,
			createContentFunc: ioutil.WriteFile,
		},

		ProjectItem{
			itemName:          "test_" + pythonProjectMetaData.pName + ".py",
			permissions:       0644,
			content:           testProjectContent(),
			parentDir:         path.Join(pythonProjectMetaData.pName, "tests"),
			createParentFunc:  os.Mkdir,
			createContentFunc: ioutil.WriteFile,
		},
	}

	return pythonProject
}

func setupContent() string {
	content := []string{
		"'''",
		"setup.py for " + pythonProjectMetaData.pName + ".\n",
		"For reference see",
		"https://packaging.python.org/guides/distributing-packages-using-setuptools/",
		"'''\n",

		"from pathlib import Path",
		"from setuptools import setup, find_packages\n",

		"HERE = Path(__file__).parent.absolute()",
		"with (HERE / 'README.md').open('rt') as fh:",
		"\tLONG_DESCRIPTION = fh.read().strip()\n\n",

		"REQUIREMENTS: dict = {",
		"\t'core': [",
		"\t\t# 'mandatory-requirement1',",
		"\t\t# 'mandatory-requirement2',",
		"\t],",
		"\t'test': [",
		"\t\t'pytest',",
		"\t],",
		"\t'dev': [",
		"\t\t#'requirement-for-development-purposes-only',",
		"\t],",
		"\t'doc': [",
		"\t\t'sphinx',",
		"\t\t'acc-py-sphinx',",
		"\t],",
		"}\n",
		"setup(",
		"\tname='test-accpy',",
		"\tversion='0.0.1.dev0',\n",
		"\tauthor='" + pythonProjectMetaData.pAuthor + "',",
		"\tauthor_email='" + pythonProjectMetaData.pMail + "',",
		"\tdescription='SHORT DESCRIPTION OF PROJECT',",
		"\tlong_description=LONG_DESCRIPTION,",
		"\tlong_description_content_type='text/markdown',",
		"\turl='',\n",
		"\tpackages=find_packages(),",
		"\tpython_requires='>=3.6, <4',",
		"\tclassifiers=[",
		"\t\t'Programming Language :: Python :: 3',",
		"\t\t'Operating System :: OS Independent',",
		"\t],\n",
		"\tinstall_requires=REQUIREMENTS['core'],",
		"\textras_require={",
		"\t\t**REQUIREMENTS,",
		"\t\t# The 'dev' extra is the union of 'test' and 'doc', with an option",
		"\t\t# to have explicit development dependencies listed.",
		"\t\t\t'dev': [req",
		"\t\t\t\tfor extra in ['dev', 'test', 'doc']",
		"\t\t\t\tfor req in REQUIREMENTS.get(extra, [])],",
		"\t\t\t# The 'all' extra is the union of all requirements.",
		"\t\t\t'all': [req for reqs in REQUIREMENTS.values() for req in reqs],",
		"\t},",
		")",
	}

	return strings.Join(content, "\n")
}

func readMeContent() string {
	content := []string{
		"# " + pythonProjectMetaData.pName + "\n",
		"SHORT DESCRIPTION OF PROJECT\n",
		`You can use [Github-flavored Markdown](https://guides.github.com/features/mastering-markdown/)
to write your content.\n`,
		"## Purpose of this project",
		"## Getting started",
		"##",
	}

	return strings.Join(content, "\n")
}

func initPyContent() string {
	content := []string{"'''",
		"Documentation for the " + pythonProjectMetaData.pName + " package",
		"'''",
		"__version__ = '0.0.1.dev0'",
	}
	return strings.Join(content, "\n")
}

func testProjectContent() string {
	content := []string{"'''",
		"High-level tests for the  package.",
		"'''",
		"import " + pythonProjectMetaData.pName,
		"def test_version():",
		"\tassert " + pythonProjectMetaData.pName + ".__version__ is not None",
	}

	return strings.Join(content, "\n")
}
