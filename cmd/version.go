package cmd

import (
	"fmt"
	"github.com/aiicy/aiicy/utils"
	"github.com/urfave/cli/v2"
	"runtime"
)

// Compile parameter
var (
	Version  string
	Revision string
)

func version(c *cli.Context) error {
	fmt.Printf("Version:      %s\nGit revision: %s\nGo version:   %s\nPlatform:     %s\n\n", Version, Revision, runtime.Version(), utils.GetHostInfo().FormatPlatformInfo())
	return nil
}
