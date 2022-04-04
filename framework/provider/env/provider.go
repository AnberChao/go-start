package env

import (
	"github.com/JoeZhao1/go-start/framework"
	"github.com/JoeZhao1/go-start/framework/contract"
)

type StartEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *StartEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewStartEnv
}

// Boot will called when the service instantiate
func (provider *StartEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *StartEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *StartEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

/// Name define the name for this service
func (provider *StartEnvProvider) Name() string {
	return contract.EnvKey
}