package short

import "github.com/seggga/cropurl/internal/cropurl/short/datatype"

type Controller struct {
	rw ShortReadWriter
}

type ShortReadWriter interface {
	ReadShort() (*datatype.ShortURL, error)
	WriteShort(short *datatype.ShortURL) error
}

func NewController() *Controller {
	return &Controller{
		//rw: rw,
	}
}

func (c *Controller) AddUser(user *datatype.ShortURL) error {
	return c.rw.WriteShort(user)
}
