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

func TestParseListOutputInvalidJSON(t *testing.T) {
	invalidJSON := `{invalid json}`

	_, err := ParseListOutput([]byte(invalidJSON))
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestStartCommand(t *testing.T) {
	cmd := StartCommand("test-project")

	if cmd.Path == "" {
		t.Error("command path should not be empty")
	}

	args := cmd.Args
	// Args[0] is the command itself
	if len(args) < 3 || args[1] != "start" || args[2] != "-s" {
		t.Errorf("unexpected args: %v", args)
	}
}

func TestStopCommand(t *testing.T) {
	cmd := StopCommand("test-project")

	if cmd.Path == "" {
		t.Error("command path should not be empty")
	}

	args := cmd.Args
	if len(args) < 2 || args[1] != "stop" {
		t.Errorf("unexpected args: %v", args)
	}
}
