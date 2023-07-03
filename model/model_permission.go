package model

type Permission struct {
	Id       uint   `xorm:"'id' bigint AUTO_INCREMENT pk" json:"id"` // 权限 ID，主键自增
	Name     string `xorm:"'name' varchar(50)" json:"name"`          // 权限名称
	Module   string `xorm:"'module' varchar(50)" json:"module"`      // 权限所属模块
	Resource string `xorm:"'resource' varchar(50)" json:"resource"`  // 权限所属资源
	Action   string `xorm:"'action' varchar(50)" json:"action"`      // 权限操作（CRUD）
	Desc     string `xorm:"'desc' varchar(255)" json:"desc"`         // 权限描述
}

func (p *Permission) TableName() string {
	return "permission" // 权限表名
}
