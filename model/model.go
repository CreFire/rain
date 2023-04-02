package model

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

type Log struct {
	Compress         bool          `yaml:"compress"`
	Level            string        `yaml:"level"`
	OutputPaths      []string      `yaml:"outputPaths"`
	EncoderConfig    EncoderConfig `yaml:"encoderConfig"`
	Maxsize          int           `yaml:"maxsize"`
	MaxAge           int           `yaml:"maxage"`
	Filename         string        `yaml:"filename"`
	Encoding         string        `yaml:"encoding"`
	ErrorOutputPaths []string      `yaml:"errorOutputPaths"`
	FileMaxBackups   int           `yaml:"file_max_backups"`
	Stdout           bool          `yaml:"stdout"`
}

type EncoderConfig struct {
	MessageKey   string `yaml:"messageKey"`
	LevelKey     string `yaml:"levelKey"`
	LevelEncoder string `yaml:"levelEncoder"`
}

type Sqlite3 struct {
	Enable bool `yaml:"enable"`
}

type Mysql struct {
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}

type Rain struct {
	Mode    string `yaml:"mode"`
	WorkDir string `yaml:"workDir"`
	LogDir  string `yaml:"logDir"`
}
