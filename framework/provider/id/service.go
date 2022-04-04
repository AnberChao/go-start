package id

import (
	"github.com/rs/xid"
)

type StartIDService struct {
}

func NewStartIDService(params ...interface{}) (interface{}, error) {
	return &StartIDService{}, nil
}

func (s *StartIDService) NewID() string {
	return xid.New().String()
}