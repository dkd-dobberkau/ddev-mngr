package ddev

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Project struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	ShortRoot string `json:"shortroot"`
	HTTPUrl   string `json:"httpurl"`
}

type listOutput struct {
	Raw []Project `json:"raw"`
}

func ParseListOutput(data []byte) ([]Project, error) {
	var output listOutput
	if err := json.Unmarshal(data, &output); err != nil {
		return nil, err
	}
	return output.Raw, nil
}

func List() ([]Project, error) {
	cmd := exec.Command("ddev", "list", "--json-output")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list DDEV projects: %w", err)
	}
	projects, err := ParseListOutput(output)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DDEV list output: %w", err)
	}
	return projects, nil
}

func StartCommand(name string) *exec.Cmd {
	return exec.Command("ddev", "start", "-s", name)
}

func StopCommand(name string) *exec.Cmd {
	return exec.Command("ddev", "stop", name)
}

func Start(name string) error {
	cmd := StartCommand(name)
	return cmd.Run()
}

func Stop(name string) error {
	cmd := StopCommand(name)
	return cmd.Run()
}
