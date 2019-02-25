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
