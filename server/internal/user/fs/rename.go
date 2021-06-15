// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

import (
	"github.com/supersonictw/virtual_host-server/internal/auth"
	"os"
	"path/filepath"
	"strings"

	"github.com/supersonictw/virtual_host-server/internal/user/fs/middleware"
)

type Rename struct {
	session *auth.Session
	path    string
	newPath string
}

func NewRename(session *auth.Session, path string) Interface {
	instance := new(Rename)
	instance.session = session
	instance.path = middleware.FullPathExpression(path, session.Identification)
	newName := session.Context.PostForm("name")
	newPath := filepath.Join(filepath.Dir(path), newName)
	instance.newPath = middleware.FullPathExpression(newPath, session.Identification)
	return instance
}

func (r *Rename) Validate() bool {
	if !middleware.RefactorPathValidator(r.path, r.session.Identification) ||
		!middleware.RefactorPathValidator(r.newPath, r.session.Identification) {
		return false
	}
	if _, err := os.Stat(r.path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (r *Rename) Refactor() Response {
	response := new(GeneralResponse)
	response.Status = false
	if !r.Validate() {
		response.Data = "Not Allowed"
		return response
	}
	err := os.Rename(r.path, r.newPath)
	if err != nil {
		response.Data = strings.Title(err.Error())
		return response
	}
	response.Status = true
	r.session.Journalist("Rename", r.path)
	return response
}
