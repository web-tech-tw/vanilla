// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/supersonictw/virtual_host-server/internal/auth"
	"github.com/supersonictw/virtual_host-server/internal/user/fs/middleware"
)

func NewAccess(c *gin.Context) *auth.Session {
	accessToken := auth.ReadAccessToken(c)
	if accessToken == "" {
		return nil
	}
	authorization, err := auth.NewAuthorization(accessToken)
	if err != nil {
		return nil
	}
	identification := authorization.GetIdentification()
	fullPath := middleware.FullPathExpression("", identification)
	if middleware.PathTypeDetector(fullPath) != middleware.Directory {
		return nil
	}
	return authorization.GetSession(c)
}
