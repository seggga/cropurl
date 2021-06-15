package user

import (
	"github.com/seggga/cropurl/internal/cropurl/user/datatype"
)

type Controller struct {
	rw UserReadWriter
}

type UserReadWriter interface {
	ReadUser(userID int64) (*datatype.User, error)
	WriteUser(user *datatype.User) error
}

func NewController() *Controller {
	return &Controller{
		//	rw: rw,
	}
}

func (c *Controller) AddUser(user *datatype.User) error {
	return c.rw.WriteUser(user)
}
