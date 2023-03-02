package utils

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

func initMysql() error {
	sqlConfig, err := config.String("sqlconn")
	if err != nil {
		logs.Info("ERROR: initMySQL error:", err)
		return err
	}
	// orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", sqlConfig)
	// orm.RunSyncdb("default", false, true)
	return nil
}
