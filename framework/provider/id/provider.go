package id

import (
	"github.com/JoeZhao1/go-start/framework"
	"github.com/JoeZhao1/go-start/framework/contract"
)

type StartIDProvider struct {
}

// Register registe a new function for make a service instance
func (provider *StartIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewStartIDService
}

// Boot will called when the service instantiate
func (provider *StartIDProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *StartIDProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *StartIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *StartIDProvider) Name() string {
	return contract.IDKey
}
