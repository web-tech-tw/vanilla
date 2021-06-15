// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

import (
	"github.com/supersonictw/virtual_host-server/internal/auth"
	"github.com/supersonictw/virtual_host-server/internal/user/fs/middleware"
	"strings"
)

type Write struct {
	session *auth.Session
	path    string
}

func NewWrite(session *auth.Session, path string) Interface {
	instance := new(Write)
	instance.session = session
	instance.path = middleware.FullPathExpression(path, session.Identification)
	return instance
}

func (w *Write) Validate() bool {
	if !middleware.RefactorPathValidator(w.path, w.session.Identification) {
		return false
	}
	return true
}

func (w *Write) Refactor() Response {
	response := new(GeneralResponse)
	response.Status = false
	if !w.Validate() {
		response.Data = "Not Allowed"
		return response
	}
	context := w.session.Context
	file, _ := context.FormFile("file")
	err := context.SaveUploadedFile(file, w.path)
	if err != nil {
		response.Data = strings.Title(err.Error())
		return response
	}
	response.Status = true
	w.session.Journalist("Write", w.path)
	return response
}
