// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package middleware

import (
	"os"
	"sort"
	"strings"

	"github.com/supersonictw/virtual_host-server/internal/auth"
)

func UserDirectoryPrefix(identification *auth.Identification) string {
	if os.Getenv("STORAGE_USER_DIRECTORY_NAME_METHOD") == "email" {
		split := strings.Split(identification.Email, "@")
		sort.Sort(sort.Reverse(sort.StringSlice(split)))
		return strings.Join(split, "/")
	}
	return identification.Identity
}
