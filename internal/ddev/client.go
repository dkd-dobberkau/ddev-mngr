package ddev

import (
	"encoding/json"
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
		return nil, err
	}
	return ParseListOutput(output)
}
