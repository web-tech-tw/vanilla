// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

import (
	"github.com/supersonictw/virtual_host-server/internal/auth"
	"os"
	"strings"

	"github.com/supersonictw/virtual_host-server/internal/user/fs/middleware"
)

type Mkdir struct {
	session *auth.Session
	path    string
}

func NewMkdir(session *auth.Session, path string) Interface {
	instance := new(Mkdir)
	instance.session = session
	instance.path = middleware.FullPathExpression(path, session.Identification)
	return instance
}

func (m *Mkdir) Validate() bool {
	if !middleware.RefactorPathValidator(m.path, m.session.Identification) {
		return false
	}
	if _, err := os.Stat(m.path); !os.IsNotExist(err) {
		return false
	}
	return true
}

func (m *Mkdir) Refactor() Response {
	response := new(GeneralResponse)
	response.Status = false
	if !m.Validate() {
		response.Data = "Not Allowed"
		return response
	}
	err := os.MkdirAll(m.path, 0755)
	if err != nil {
		response.Data = strings.Title(err.Error())
		return response
	}
	response.Status = true
	m.session.Journalist("Mkdir", m.path)
	return response
}
