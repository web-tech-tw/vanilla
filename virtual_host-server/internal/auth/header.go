// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package auth

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func ReadAccessToken(c *gin.Context) string {
	authorization := strings.Split(c.GetHeader("Authorization"), " ")
	if len(authorization) != 2 {
		return ""
	}
	return authorization[1]
}
