package ddev

import (
	"testing"
)

func TestParseListOutput(t *testing.T) {
	jsonOutput := `{"raw":[{"name":"test-project","status":"running","shortroot":"~/test","httpurl":"https://test.ddev.site"}]}`

	projects, err := ParseListOutput([]byte(jsonOutput))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(projects) != 1 {
		t.Fatalf("expected 1 project, got %d", len(projects))
	}

	if projects[0].Name != "test-project" {
		t.Errorf("expected name 'test-project', got '%s'", projects[0].Name)
	}

	if projects[0].Status != "running" {
		t.Errorf("expected status 'running', got '%s'", projects[0].Status)
	}
}
