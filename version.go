package main

import (
	"fmt"
)

type VersionCommand struct{}

var Version VersionCommand

func (v *VersionCommand) Execute(args []string) error {
	fmt.Println(HubVersion)
	return nil
}

func (v *VersionCommand) Usage() string {
	return ""
}