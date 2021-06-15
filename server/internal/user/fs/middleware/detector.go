// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package middleware

import "os"

const (
	NotExists    = 0
	Directory    = 1
	File         = 2
	UnknownType  = 3
	UnknownError = 4
)

func PathTypeDetector(path string) int {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return NotExists
		}
		return UnknownError
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return Directory
	case mode.IsRegular():
		return File
	default:
		return UnknownType
	}
}
