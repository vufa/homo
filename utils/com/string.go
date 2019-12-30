//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, March 2019
//

package com

import "os"

// IfStringInArray return whether a string in string array
func IfStringInArray(a string, list []string) bool {
	for _, sub := range list {
		if sub == a {
			return true
		}
	}
	return false
}

// PathExists return true if given path exist.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}
