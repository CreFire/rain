package dal

import (
	"github.com/CreFire/rain/consts"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	DBType consts.DBType
)
