// Package VHS: Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/supersonictw/virtual_host-server/internal/user"
	"github.com/supersonictw/virtual_host-server/internal/user/fs"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()

	if os.Getenv("CORS_SUPPORT") == "yes" {
		var frontendURI string
		if hostname := os.Getenv("FRONTEND_HOSTNAME"); os.Getenv("FRONTEND_SSL") == "yes" {
			frontendURI = fmt.Sprintf("https://%s", hostname)
		} else {
			frontendURI = fmt.Sprintf("http://%s", hostname)
		}

		corsConfig := cors.DefaultConfig()
		corsConfig.AllowOrigins = []string{frontendURI}
		corsConfig.AllowCredentials = true
		corsConfig.AddAllowHeaders("Authorization")

		router.Use(cors.New(corsConfig))
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "virtual_host-system",
			"copyright":   "(c)2021 SuperSonic(https://github.com/supersonictw)",
		})
	})

	router.GET("/profile", func(c *gin.Context) {
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   session.Identification,
		})
	})

	router.GET("/user/*path", func(c *gin.Context) {
		path := c.Param("path")
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := fs.NewRead(session, path)
		if result := handler.Refactor().(*fs.ReadResponse); result.Status {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"data":   result,
			})
		} else if result.Type == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status": 404,
				"code":   result.Type,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"reason": result.GetData(),
			})
		}
	})

	router.POST("/user/*path", func(c *gin.Context) {
		path := c.Param("path")
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := fs.NewMkdir(session, path)
		if result := handler.Refactor(); result.GetStatus() {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"reason": result.GetData(),
			})
		}
	})

	router.PUT("/user/*path", func(c *gin.Context) {
		path := c.Param("path")
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := fs.NewWrite(session, path)
		if result := handler.Refactor(); result.GetStatus() {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"reason": result.GetData(),
			})
		}
	})

	router.PATCH("/user/*path", func(c *gin.Context) {
		path := c.Param("path")
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := fs.NewRename(session, path)
		if result := handler.Refactor(); result.GetStatus() {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"reason": result.GetData(),
			})
		}
	})

	router.DELETE("/user/*path", func(c *gin.Context) {
		path := c.Param("path")
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := fs.NewRemove(session, path)
		if result := handler.Refactor(); result.GetStatus() {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"reason": result.GetData(),
			})
		}
	})

	router.POST("/zip/*path", func(c *gin.Context) {
		path := c.Param("path")
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := fs.NewZip(session, path)
		if result := handler.Refactor(); result.GetStatus() {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"reason": result.GetData(),
			})
		}
	})

	router.DELETE("/zip/*path", func(c *gin.Context) {
		path := c.Param("path")
		session := user.NewAccess(c)
		if session == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
			})
			return
		}
		handler := fs.NewUnzip(session, path)
		if result := handler.Refactor(); result.GetStatus() {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"reason": result.GetData(),
			})
		}
	})

	exposePort := fmt.Sprintf(":%s", os.Getenv("EXPOSE_PORT"))
	if err := router.Run(exposePort); err != nil {
		panic(err)
	}
}
