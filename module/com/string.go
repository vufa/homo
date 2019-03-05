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
