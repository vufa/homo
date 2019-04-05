//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package com

// IfStringInArray return whether a string in string array
func IfStringInArray(a string, list []string) bool {
	for _, sub := range list {
		if sub == a {
			return true
		}
	}
	return false
}
