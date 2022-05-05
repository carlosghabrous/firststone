package lang

import "testing"

const (
	wrongName = ""
	goodName  = ""
)

func TestCheckNamingConventions(t *testing.T) {
	pythonProject = PythonProject{Name: wrongName, Language: python, ProjectItems: pythonProjectItems}
	err := pythonProject.CheckNamingConventions(pythonProject.Name)

	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}

	pythonProject.Name = goodName
	err = pythonProject.CheckNamingConventions(pythonProject.Name)

	if err != nil {
		t.Fatalf("expected nil error, got %s", err)
	}
}
