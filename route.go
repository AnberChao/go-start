package main

import "go-start/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
