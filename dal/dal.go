package dal

import (
	"fmt"
	"github.com/CreFire/rain/consts"
	"github.com/CreFire/rain/tools/config"
	log "github.com/CreFire/rain/tools/log"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

var (
	dbEngine *xorm.Engine
	DbType   consts.DBType
)

// mysqlDsn example  user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
const mysqlDsn = "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s&readTimeout" +
	"=1s&writeTimeout=1s&interpolateParams=true"

func NewDB() (err error) {
	//nolint:critic
	if config.Conf.Sqlite3 != nil && config.Conf.Sqlite3.Enable {
		initSQLite()
		DbType = consts.DBTypeSQLite
	} else if config.Conf.Mysql != nil {
		InitMysql()
		DbType = consts.DBTypeMySQL
	} else {
		log.Fatal("No database available")
	}
	if dbEngine == nil {
		log.Fatal("no available database")
	}
	sqlDB := dbEngine.DB()
	sqlDB.SetMaxIdleConns(200)
	sqlDB.SetMaxOpenConns(300)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	return nil
}

func initSQLite() {
	sqliteConfig := config.Conf.Sqlite3
	if sqliteConfig == nil {
		log.Fatal("nil SQLite config")
	}
	var err error
	dbEngine, err = xorm.NewEngine(consts.DBTypeSQLite, "./data")
	if err != nil {
		log.Fatal("nil SQLite config")
	}
}

func InitMysql() {
	mysqlConfig := config.Conf.Mysql
	if mysqlConfig == nil {
		log.Fatal("mysql config err")
	}
	dsn := fmt.Sprintf(mysqlDsn, mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Db)
	var err error
	fmt.Printf("dsn:%v \n", dsn)
	dbEngine, err = xorm.NewEngine(consts.DBTypeMySQL, dsn)
	if err != nil {
		log.Fatal("mysql NewEngine err", log.Err(err))
	}
}

func GetDb() *xorm.Engine {
	return dbEngine
}
