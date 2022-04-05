package env

import (
	"github.com/JoeZhao1/go-start/framework"
	"github.com/JoeZhao1/go-start/framework/contract"
	"github.com/JoeZhao1/go-start/framework/provider/app"
	tests "github.com/JoeZhao1/go-start/test"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStartEnvProvider(t *testing.T) {
	Convey("test start env normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewStartContainer()
		sp := &app.StartAppProvider{BaseFolder: basePath}

		err := c.Bind(sp)
		So(err, ShouldBeNil)

		sp2 := &StartEnvProvider{}
		err = c.Bind(sp2)
		So(err, ShouldBeNil)

		envServ := c.MustMake(contract.EnvKey).(contract.Env)
		So(envServ.AppEnv(), ShouldEqual, "development")
		// So(envServ.Get("DB_HOST"), ShouldEqual, "127.0.0.1")
		// So(envServ.AppDebug(), ShouldBeTrue)
	})
}