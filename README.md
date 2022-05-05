![GitHub Workflow Status](https://img.shields.io/github/workflow/status/carlosghabrous/firststone/Go)

# firststone
A simple util to automate the creation of projects in different languages

First Stone is a utility that automates the creation of project's directories and files for certain programming languages. 
Current supported languages are: Python, Go. 
    
## Why
Creating a project in a language typically requires certain infrastructure. This infrastructure is light (files and folders), but: 
1. it should be consistent across projects (all python projects should contain the same base directories and files, like setup.py, __init__.py for packages, test directory...)
2. it is tedious to create it manually (mkdir, touch)
3. if this task is left to copy-pasting from other existing projects, there are certain things that will need to be changed in anycase, like the project name. This leads to 
manually editting files to change strings here and there

The aim of first stone is just that, laying the first stone of a project quickly, automatically and without forgetting to edit files here and there. 

## Installation
1. Go get to install the latest version of the package
    ```
    go get -v github.com/carlosghabrous/firststone
    ```

2. Make sure GOPATH is included in your PATH. Run the tool: 
    ```
    firststone <command> [-flags] [arguments]
    ```

    For instance, to create a python project:
    ```
    mkdir my-new-project
    cd my-new-project
    firststone init my-new-project python
    ```

3. To get some help
    ```
    firststone --help
    ```

    You can also get help by command:
    ```
    firststone init --help
    ```

