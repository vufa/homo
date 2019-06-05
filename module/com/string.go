//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
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
