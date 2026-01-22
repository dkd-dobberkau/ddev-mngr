package main

import "github.com/dkd-dobberkau/ddev-mngr/cmd"

var version = "dev"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
