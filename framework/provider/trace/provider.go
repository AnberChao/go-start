package trace

import (
	"github.com/JoeZhao1/go-start/framework"
	"github.com/JoeZhao1/go-start/framework/contract"
)

type StartTraceProvider struct {
	c framework.Container
}

// Register registe a new function for make a service instance
func (provider *StartTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewStartTraceService
}

// Boot will called when the service instantiate
func (provider *StartTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *StartTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *StartTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

/// Name define the name for this service
func (provider *StartTraceProvider) Name() string {
	return contract.TraceKey
}