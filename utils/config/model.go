package config

import "time"

type Config struct {
	Server  *Server  `yaml:"server" json:"server"`
	Log     *Log     `yaml:"log" json:"log"`
	Sqlite3 *Sqlite3 `yaml:"sqlite3" json:"sqlite3"`
	Mysql   *Mysql   `yaml:"mysql"`
	Rain    *Rain    `yaml:"rain"`
	Model   string   `yaml:"model"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Sqlite3 struct {
	DatareSource string `yaml:"datareSource"`
	Enable       bool   `yaml:"enable"`
}

type Mysql struct {
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Charset  string `yaml:"charset"`
}

type Rain struct {
	Mode    string `yaml:"mode"`
	WorkDir string `yaml:"workDir"`
	LogDir  string `yaml:"logDir"`
}

type Log struct {
	OutputPaths      []string      `yaml:"outputPaths"`
	EncoderConfig    EncoderConfig `yaml:"encoderConfig"`
	ErrorOutputPaths []string      `yaml:"errorOutputPaths"`

	RotateConfig
	Compress       bool   `yaml:"compress"`
	Level          string `yaml:"level"`
	Filename       string `yaml:"filename"`
	Maxsize        int    `yaml:"maxsize"`
	MaxAge         int    `yaml:"maxage"`
	Encoding       string `yaml:"encoding"`
	FileMaxBackups int    `yaml:"file_max_backups"`
	Stdout         bool   `yaml:"stdout"`
}

type RotateConfig struct {
	// 共用配置
	Filename string // 完整文件名
	MaxAge   int    // 保留旧日志文件的最大天数

	// 按时间轮转配置
	RotationTime time.Duration // 日志文件轮转时间

	// 按大小轮转配置
	MaxSize    int  // 日志文件最大大小（MB）
	MaxBackups int  // 保留日志文件的最大数量
	Compress   bool // 是否对日志文件进行压缩归档
	LocalTime  bool // 是否使用本地时间，默认 UTC 时间
}

type EncoderConfig struct {
	MessageKey   string `yaml:"messageKey"`
	LevelKey     string `yaml:"levelKey"`
	LevelEncoder string `yaml:"levelEncoder"`
}
