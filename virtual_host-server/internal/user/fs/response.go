// Virtual Host System - Server
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

type Response interface {
	GetStatus() bool
	GetData() interface{}
}

type GeneralResponse struct {
	Response
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func (g *GeneralResponse) GetStatus() bool {
	return g.Status
}

func (g *GeneralResponse) GetData() interface{} {
	return g.Data
}
