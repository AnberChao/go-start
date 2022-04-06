package id
//
//import (
//	"testing"
//
//	"github.com/JoeZhao1/go-start/framework"
//	"github.com/JoeZhao1/go-start/framework/contract"
//	"github.com/JoeZhao1/go-start/framework/provider/app"
//	"github.com/JoeZhao1/go-start/framework/provider/config"
//	"github.com/JoeZhao1/go-start/framework/provider/env"
//	"github.com/JoeZhao1/go-start/framework/util"
//
//	. "github.com/smartystreets/goconvey/convey"
//)
//
//func TestConsoleLog_Normal(t *testing.T) {
//	Convey("test start console log normal case", t, func() {
//		basePath := util.GetExecDirectory()
//		c := framework.NewStartContainer()
//		c.Singleton(&app.StartAppProvider{BasePath: basePath})
//		c.Singleton(&env.StartEnvProvider{})
//		c.Singleton(&config.StartConfigProvider{})
//
//		err := c.Singleton(&StartIDProvider{})
//		So(err, ShouldBeNil)
//
//		idService := c.MustMake(contract.IDKey).(contract.IDService)
//		xid := idService.NewID()
//		t.Log(xid)
//		So(xid, ShouldNotBeEmpty)
//	})
//}
