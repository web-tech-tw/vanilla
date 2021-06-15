// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package auth

type Identification struct {
	Identity    string `json:"identity"`
	DisplayName string `json:"name"`
	PictureURL  string `json:"picture"`
	Email       string `json:"email"`
}
