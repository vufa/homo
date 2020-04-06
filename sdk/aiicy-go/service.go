//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that can
// be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, January 2020
//

package aiicy

import (
	"os"
	"runtime/debug"

	"github.com/aiicy/aiicy-go/logger"
)

type Service struct {
	CfgPath string
}

// Run service
func Run(s Service, handle func(Context) error) {
	defer func() {
		if r := recover(); r != nil {
			logger.S.Errorf("service is stopped with panic: %s\n%s", r, string(debug.Stack()))
		}
	}()
	c, err := newContext(s)
	if err != nil {
		logger.S.Errorw("failed to create context", logger.Error(err))
		return
	}
	c.log.Info("service starting: ", os.Args)
	err = handle(c)
	if err != nil {
		c.log.Errorw("service is stopped with error", logger.Error(err))
	} else {
		c.log.Info("service stopped")
	}
}
