//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package com

import (
	"github.com/sirupsen/logrus"
	"io"
)

func IOClose(name string, c io.Closer) {
	err := c.Close()
	if err != nil {
		logrus.Warnf("close [%s] failed: %s", name, err.Error())
	}
}
