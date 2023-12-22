package dal

import (
	"fmt"
	"github.com/CreFire/rain/common"
	"github.com/CreFire/rain/utils/config"
	"github.com/CreFire/rain/utils/log"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

// sqlite 初始化
func initSQLite() {
	sqliteConfig := config.Conf.Sqlite3
	if sqliteConfig == nil {
		log.Fatal("nil SQLite config")
	}
	var err error
	dbEngine, err = xorm.NewEngine(common.DBTypeSQLite, sqliteConfig.DatareSource)
	if err != nil {
		log.Fatal(fmt.Sprintf("nil SQLite config,sqliteConfig:%v", sqliteConfig.DatareSource))
	}
}
