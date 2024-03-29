package common

type DBType string

const (
	DBTypeMySQL  = "mysql"
	DBTypeSQLite = "sqlite3"
)

type AttachmentType int32

const (
	AttachmentTypeLocal AttachmentType = iota
	// AttachmentTypeUpOSS 又拍云
	AttachmentTypeUpOSS
	// AttachmentTypeQiNiuOSS 七牛云
	AttachmentTypeQiNiuOSS
	// AttachmentTypeSMMS sm.ms
	AttachmentTypeSMMS
	// AttachmentTypeAliOSS 阿里云OSS
	AttachmentTypeAliOSS
	// AttachmentTypeBaiDuOSS 百度云OSS
	AttachmentTypeBaiDuOSS
	// AttachmentTypeTencentCOS 腾讯COS
	AttachmentTypeTencentCOS
	// AttachmentTypeHuaweiOBS 华为OBS
	AttachmentTypeHuaweiOBS
	// AttachmentTypeMinIO AttachmentTypeMinIO
	AttachmentTypeMinIO
)

type Role int32

const (
	VISTOR Role = 0
	CUSTOM Role = 1
	DICTOR Role = 2
	ADMIN  Role = 3
)
