package model

import "time"

type Tenant struct {
	Id          uint      `xorm:"bigint 'id' AUTO_INCREMENT pk" json:"id"`        // 租户 ID，主键自增
	UserId      uint      `xorm:"'user_id' int(11) notnull index" json:"user_id"` // 租户所属用户 ID，外键关联用户表
	Date        time.Time `xorm:"'date' date" json:"date"`                        // 租金缴纳日期
	Water       float64   `xorm:"'water' double(10,2)" json:"water"`              // 水费
	Electricity float64   `xorm:"'electricity' double(10,2)" json:"electricity"`  // 电费
	Paid        bool      `xorm:"'paid' bool" json:"paid"`                        // 是否已经缴费
	AmountDue   float64   `xorm:"'amount_due' double(10,2)" json:"amount_due"`    // 待缴费用
}

func (t *Tenant) TableName() string {
	return "tenant" // 租户表名
}
