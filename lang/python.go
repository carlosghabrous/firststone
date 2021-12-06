package lang

import (
	"fmt"
	"os"
)

type PythonProject Project

const (
	python                = "python"
	projectNameDash       = "PROJECT-NAME"
	projectNameUnderscore = "PROJECT_NAME"
)

var readmeContent string = `# ` +
	projectNameDash +
	`
SHORT DESCRIPTION OF PROJECT

You can use [Github-flavored Markdown](https://guides.github.com/features/mastering-markdown/)
to write your content.

## Purpose of this project
## Getting started
##`

var setupContent string = `"""
setup.py for ` +
	projectNameDash +
	`.

For reference see
https://packaging.python.org/guides/distributing-packages-using-setuptools/

"""
from pathlib import Path
from setuptools import setup, find_packages


HERE = Path(__file__).parent.absolute()
with (HERE / 'README.md').open('rt') as fh:
    LONG_DESCRIPTION = fh.read().strip()


REQUIREMENTS: dict = {
    'core': [
        # 'mandatory-requirement1',
        # 'mandatory-requirement2',
    ],
    'test': [
        'pytest',
    ],
    'dev': [
        # 'requirement-for-development-purposes-only',
    ],
    'doc': [
        'sphinx',
        'acc-py-sphinx',
    ],
}


setup(
    name='` +
	projectNameDash +
	`',
    version="0.0.1.dev0",

    author='Carlos Ghabrous Larrea',
    author_email='carlos.ghabrous@gmail.com',
    description='SHORT DESCRIPTION OF PROJECT',
    long_description=LONG_DESCRIPTION,
    long_description_content_type='text/markdown',
    url='',

    packages=find_packages(),
    python_requires='>=3.6, <4',
    classifiers=[
        "Programming Language :: Python :: 3",
        "Operating System :: OS Independent",
    ],

    install_requires=REQUIREMENTS['core'],
    extras_require={
        **REQUIREMENTS,
        # The 'dev' extra is the union of 'test' and 'doc', with an option
        # to have explicit development dependencies listed.
        'dev': [req
                for extra in ['dev', 'test', 'doc']
                for req in REQUIREMENTS.get(extra, [])],
        # The 'all' extra is the union of all requirements.
        'all': [req for reqs in REQUIREMENTS.values() for req in reqs],
    },
)`

var initContent = `"""
Documentation for the ` +
	projectNameUnderscore +
	`package

"""

__version__ = "0.0.1.dev0"`

var testContent = `"""
High-level tests for the  package.

"""

import ` +
	projectNameUnderscore +
	`


def test_version():
    assert test_accpy.__version__ is not None`

var pythonProjectItems = []ProjectItem{
	{Name: "README.md",
		Parent:     ".",
		Permission: 0644,
		Content:    readmeContent},
	{Name: "setup.py",
		Parent:     ".",
		Permission: 0644,
		Content:    setupContent},
	{Name: projectNameUnderscore,
		Parent:     ".",
		Permission: os.ModeDir | 0755,
		Content:    ""},
	{Name: "__init__.py",
		Parent:     ".",
		Permission: 0644,
		Content:    initContent},
	{Name: "tests",
		Parent:     projectNameUnderscore,
		Permission: os.ModeDir | 0755,
		Content:    "",
	},
	{Name: "__init__.py",
		Parent:     "tests",
		Permission: 0644,
		Content:    "",
	},
	{Name: "test_" + projectNameUnderscore,
		Parent:     "tests",
		Permission: 0644,
		Content:    testContent},
}

var pythonProject = PythonProject{Language: python, ProjectItems: pythonProjectItems}

func init() {
	RegisterLanguage(python, &pythonProject)
}

func (p *PythonProject) CheckNamingConventions(name string) error {
	//TODO: after checking conventions are ok, assign name and return nil
	p.Name = name
	fmt.Printf("Naming conventions for project %s OK\n", p.Name)
	return nil
}

func (p *PythonProject) Build() (err error) {
	return buildProject(&pythonProjectItems)
}
