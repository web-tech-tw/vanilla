// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

import (
	"encoding/base64"
	"github.com/supersonictw/virtual_host-server/internal/auth"
	"github.com/supersonictw/virtual_host-server/internal/user/fs/middleware"
	"io/fs"
	"io/ioutil"
	"strings"
)

type ReadResponse struct {
	GeneralResponse
	Type int `json:"type"`
}

type File struct {
	Name         string `json:"name"`
	Type         int    `json:"type"`
	Size         int64  `json:"size"`
	Mode         string `json:"mode"`
	LastModified int64  `json:"lastModified"`
}

type Read struct {
	session *auth.Session
	path    string
}

func NewRead(session *auth.Session, path string) Interface {
	instance := new(Read)
	instance.session = session
	instance.path = middleware.FullPathExpression(path, session.Identification)
	return instance
}

func (r *Read) Validate() bool {
	if !middleware.RefactorPathValidator(r.path, r.session.Identification) {
		return false
	}
	return true
}

func getFilesInDirectory(files []fs.FileInfo) []*File {
	files_ := make([]*File, len(files))
	for i, f := range files {
		type_ := 0
		if f.IsDir() {
			type_ = 1
		}
		files_[i] = &File{
			Name:         f.Name(),
			Type:         type_,
			Size:         f.Size(),
			Mode:         f.Mode().String(),
			LastModified: f.ModTime().UnixNano(),
		}
	}
	return files_
}

func (r *Read) directoryHandler(response *ReadResponse) error {
	directory, err := ioutil.ReadDir(r.path)
	if err != nil {
		return err
	}
	response.Status = true
	response.Data = getFilesInDirectory(directory)
	return nil
}

func (r *Read) fileHandler(response *ReadResponse) error {
	content, err := ioutil.ReadFile(r.path)
	if err != nil {
		return err
	}
	response.Status = true
	response.Data = base64.StdEncoding.EncodeToString(content)
	return nil
}

func (r *Read) Refactor() Response {
	response := new(ReadResponse)
	response.Status = false
	response.Type = middleware.PathTypeDetector(r.path)
	if !r.Validate() {
		response.Data = "Not Allowed"
		return response
	}
	switch response.Type {
	case middleware.Directory:
		err := r.directoryHandler(response)
		if err != nil {
			response.Data = strings.Title(err.Error())
			return response
		}
		break
	case middleware.File:
		err := r.fileHandler(response)
		if err != nil {
			response.Data = strings.Title(err.Error())
			return response
		}
		break
	default:
		return response
	}
	r.session.Journalist("Read", r.path)
	return response
}
