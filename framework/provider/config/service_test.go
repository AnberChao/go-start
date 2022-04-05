package config

import (
	"path/filepath"
	"testing"

	"github.com/JoeZhao1/go-start/framework/contract"
	tests "github.com/JoeZhao1/go-start/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStartConfig_GetInt(t *testing.T) {
	Convey("test go-start env normal case", t, func() {
		basePath := tests.BasePath
		folder := filepath.Join(basePath, "config")
		serv, err := NewStartConfig(folder, map[string]string{}, contract.EnvDevelopment)
		So(err, ShouldBeNil)
		conf := serv.(*StartConfig)
		timeout := conf.GetInt("database.mysql.timeout")
		So(timeout, ShouldEqual, 1)
	})
}
