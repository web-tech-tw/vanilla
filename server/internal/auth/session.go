// Vanilla
// (c)2021 SuperSonic (https://github.com/supersonictw)

package auth

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type Session struct {
	Identification *Identification
	Context        *gin.Context
}

func (s *Session) Journalist(action string, target string) {
	logRootPath := os.Getenv("LOG_DIRECTORY_PATH")
	timestamp := time.Now().Format("2006-01-02")
	logPath := filepath.Join(logRootPath, fmt.Sprintf("%s.log", timestamp))
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	actionCode := fmt.Sprintf("[%s]", action)
	logger := log.New(file, actionCode, log.Ltime)
	logger.Printf(
		"%s (%s, %s, %s)\n",
		target,
		s.Identification.DisplayName,
		s.Identification.Email,
		s.Identification.Identity,
	)
}
