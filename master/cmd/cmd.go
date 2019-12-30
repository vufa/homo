//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, August 2019
//

package cmd

import (
	"github.com/countstarlight/homo/master/config"
	"github.com/countstarlight/homo/utils/logger"
	"github.com/urfave/cli/v2"
	"os"
)

// Execute execute
func Execute() {
	app := cli.NewApp()
	app.Name = config.AppName
	app.Version = config.AppVersion
	app.Usage = "Expand the combination of artificial intelligence applications and the IoT"
	app.Action = startInternal
	app.Flags = flags
	if err := app.Run(os.Args); err != nil {
		logger.S.Fatal(err)
	}
}
