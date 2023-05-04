package common

import "time"

const (
	AccessTokenExpiredSeconds = 24 * 3600
	RefreshTokenExpiredDays   = 30
	TokenAccessCachePrefix    = "admin_access_token_"
	TokenRefreshCachePrefix   = "admin_refresh_token_"
	AdminTokenHeaderName      = "Admin-Authorization"
	AuthorizedUser            = "authorized_user"
	CodePrefix                = "code_"
	CodeValidDuration         = time.Second
	OneTimeTokenQueryName     = "ott"
	SessionID                 = "session_id"
	AccessPermissionKeyPrefix = "access_permission_"
	MySecretKey               = "key"
)

const (
	ROLE_TOURIST ROLE_TYPE = iota
	ROLE_MEMBER
	ROLE_NURSE
	ROLE_DICTOR
	ROLE_ADMIN
)

type ROLE_TYPE int32
