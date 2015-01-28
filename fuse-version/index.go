package main

import (
	"fmt"
	"os"

	fuseversion "github.com/jbenet/go-fuse-version"
)

func main() {
	sys, err := fuseversion.LocalFuseSystems()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	for name, s := range *sys {
		fmt.Printf("%s.FuseVersion: %s\n", name, s.FuseVersion)
		fmt.Printf("%s.AgentVersion: %s\n", name, s.AgentVersion)
		fmt.Printf("%s.AgentName: %s\n", name, s.AgentName)
	}
}
