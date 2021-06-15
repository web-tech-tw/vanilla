// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package middleware

import (
	"path/filepath"
	"strings"

	"github.com/supersonictw/virtual_host-server/internal/auth"
)

func RefactorPathValidator(path string, identification *auth.Identification) bool {
	if !filepath.IsAbs(path) {
		return false
	}
	userDirectoryPath := FullPathExpression("", identification)
	if !strings.HasPrefix(path, userDirectoryPath) {
		return false
	}
	return true
}
