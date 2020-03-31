//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that can
// be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, March 2019
//

package utils

// IfStringInArray return whether a string in string array
func IfStringInArray(a string, list []string) bool {
	for _, sub := range list {
		if sub == a {
			return true
		}
	}
	return false
}
