//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, March 2019
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
