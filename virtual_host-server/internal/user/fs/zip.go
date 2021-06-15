// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

import (
	"archive/zip"
	"github.com/supersonictw/virtual_host-server/internal/auth"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/supersonictw/virtual_host-server/internal/user/fs/middleware"
)

type Zip struct {
	session *auth.Session
	path    string
	origin  string
}

func NewZip(session *auth.Session, path string) Interface {
	instance := new(Zip)
	instance.session = session
	instance.path = middleware.FullPathExpression(path, session.Identification)
	origin := session.Context.PostForm("origin")
	originPath := filepath.Join(filepath.Dir(path), origin)
	instance.origin = middleware.FullPathExpression(originPath, session.Identification)
	return instance
}

func (z *Zip) Validate() bool {
	if !middleware.RefactorPathValidator(z.path, z.session.Identification) ||
		!middleware.RefactorPathValidator(z.origin, z.session.Identification) {
		return false
	}
	if _, err := os.Stat(z.origin); os.IsNotExist(err) {
		return false
	}
	return true
}

func (z *Zip) compress(response *GeneralResponse) {
	destinationFile, err := os.Create(z.path)
	if err != nil {
		response.Data = strings.Title(err.Error())
		return
	}

	w := zip.NewWriter(destinationFile)
	defer func() {
		err := w.Close()
		if err != nil {
			panic(err)
		}
	}()

	err = filepath.Walk(z.origin, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		relPath := strings.TrimPrefix(filePath, filepath.Dir(z.origin))
		zipFile, err := w.Create(relPath)
		if err != nil {
			return err
		}
		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		response.Data = strings.Title(err.Error())
		return
	}

	response.Status = true
}

func (z *Zip) Refactor() Response {
	response := new(GeneralResponse)
	response.Status = false
	if !z.Validate() {
		response.Data = "Not Allowed"
		return response
	}
	z.compress(response)
	z.session.Journalist("Zip", z.path)
	return response
}
