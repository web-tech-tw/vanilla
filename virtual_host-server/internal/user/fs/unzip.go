// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

import (
	"archive/zip"
	"github.com/supersonictw/virtual_host-server/internal/auth"
	"github.com/supersonictw/virtual_host-server/internal/user/fs/middleware"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Unzip struct {
	session *auth.Session
	path    string
	target  string
}

func NewUnzip(session *auth.Session, path string) Interface {
	instance := new(Unzip)
	instance.session = session
	instance.path = middleware.FullPathExpression(path, session.Identification)
	instance.target = middleware.FullPathExpression(filepath.Dir(path), session.Identification)
	return instance
}

func (u *Unzip) Validate() bool {
	if !middleware.RefactorPathValidator(u.path, u.session.Identification) ||
		!middleware.RefactorPathValidator(u.target, u.session.Identification) {
		return false
	}
	if _, err := os.Stat(u.path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (u *Unzip) decompress(response *GeneralResponse) {
	var filenames []string

	r, err := zip.OpenReader(u.path)
	if err != nil {
		response.Data = map[string]interface{}{
			"error":    strings.Title(err.Error()),
			"effected": filenames,
		}
	}
	defer func() {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}()

	for _, f := range r.File {
		currentPath := filepath.Join(u.target, f.Name)

		// ZipSlip
		if !middleware.RefactorPathValidator(currentPath, u.session.Identification) {
			response.Data = map[string]interface{}{
				"error":    "Illegal file path requested in zip file",
				"effected": filenames,
			}
			return
		}

		filenames = append(filenames, currentPath)

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(currentPath, 0755)
			if err != nil {
				response.Data = map[string]interface{}{
					"error":    strings.Title(err.Error()),
					"effected": filenames,
				}
				return
			}
		} else {
			if err = os.MkdirAll(filepath.Dir(currentPath), 0755); err != nil {
				response.Data = map[string]interface{}{
					"error":    strings.Title(err.Error()),
					"effected": filenames,
				}
				return
			}

			outFile, err := os.OpenFile(currentPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				response.Data = map[string]interface{}{
					"error":    strings.Title(err.Error()),
					"effected": filenames,
				}
				return
			}

			rc, err := f.Open()
			if err != nil {
				response.Data = map[string]interface{}{
					"error":    strings.Title(err.Error()),
					"effected": filenames,
				}
				return
			}
			_, err = io.Copy(outFile, rc)
			if err := outFile.Close(); err != nil {
				response.Data = map[string]interface{}{
					"error":    strings.Title(err.Error()),
					"effected": filenames,
				}
				return
			}
			if err := rc.Close(); err != nil {
				response.Data = map[string]interface{}{
					"error":    strings.Title(err.Error()),
					"effected": filenames,
				}
				return
			}
		}
	}
	response.Status = true
	response.Data = map[string][]string{
		"effected": filenames,
	}
}

func (u *Unzip) Refactor() Response {
	response := new(GeneralResponse)
	response.Status = false
	if !u.Validate() {
		response.Data = "Not Allowed"
		return response
	}
	u.decompress(response)
	u.session.Journalist("Unzip", u.path)
	return response
}
