package dal

import (
	"fmt"
	"github.com/CreFire/rain/common"
	"github.com/CreFire/rain/utils/config"
	log "github.com/CreFire/rain/utils/log"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

var (
	dbEngine *xorm.Engine
	DbType   common.DBType
)

// mysqlDsn example  user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
const mysqlDsn = "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s&readTimeout" +
	"=1s&writeTimeout=1s&interpolateParams=true"

func NewDB() {
	//nolint:critic
	if config.Conf.Sqlite3 != nil && config.Conf.Sqlite3.Enable {
		initSQLite()
		DbType = common.DBTypeSQLite
	} else if config.Conf.Mysql != nil {
		InitMysql()
		DbType = common.DBTypeMySQL
	} else {
		log.Fatal("No Set database conf")
		return
	}
	if dbEngine == nil {
		log.Fatal("dbEngine == nil no available database")
		return
	}
	sqlDB := dbEngine.DB()
	sqlDB.SetMaxIdleConns(200)
	sqlDB.SetMaxOpenConns(300)
	sqlDB.SetConnMaxIdleTime(time.Hour)

}

func InitMysql() {
	mysqlConfig := config.Conf.Mysql
	if mysqlConfig == nil {
		log.Fatal("mysql config err")
	}
	dsn := fmt.Sprintf(mysqlDsn, mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Db)
	var err error
	log.Info("mysql", "dsn", dsn, "time", time.Now())
	dbEngine, err = xorm.NewEngine(common.DBTypeMySQL, dsn)
	if err != nil {
		log.Fatal("mysql NewEngine err", "err", err)
	}
	if err = dbEngine.Ping(); err != nil {
		fmt.Println("Error on pinging database: ", err)
		return
	}
	fmt.Println("Database connected successfully!")
}

func GetDb() *xorm.Engine {
	return dbEngine
}
