//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that can
// be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, December 2019
//

package logger

// LogConfig for logging
type LogConfig struct {
	Path   string `yaml:"path" json:"path"`
	Level  string `yaml:"level" json:"level" default:"info" validate:"regexp=^(info|debug|warn|error)$"`
	Mode   string `yaml:"mode" json:"mode" default:"console" validate:"regexp=^(console|file)$"`
	Format string `yaml:"format" json:"format" default:"text" validate:"regexp=^(text|json)$"`
	Age    struct {
		Max int `yaml:"max" json:"max" default:"15" validate:"min=1"`
	} `yaml:"age" json:"age"` // days
	Size struct {
		Max int `yaml:"max" json:"max" default:"50" validate:"min=1"`
	} `yaml:"size" json:"size"` // in MB
	Backup struct {
		Max int `yaml:"max" json:"max" default:"15" validate:"min=1"`
	} `yaml:"backup" json:"backup"`
}
