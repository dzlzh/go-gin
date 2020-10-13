package config

type Service struct {
	System System `mapstructure:"system"`
	Log    Log    `mapstructure:"log"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
}

type System struct {
	Name            string `mapstructure:"name"`
	Url             string `mapstructure:"url"`
	Addr            int    `mapstructure:"addr"`
	Tls             bool   `mapstructure:"tls"`
	Env             string `mapstructure:"env"`
	Debug           bool   `mapstructure:"debug"`
	RuntimeRootPath string `mapstructure:"runtime-root-path"`
}

type Log struct {
	Level string `mapstructure:"level"`
	Path  string `mapstructure:"path"`
}

type Mysql struct {
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Addr         string `mapstructure:"addr"`
	Database     string `mapstructure:"database"`
	Config       string `mapstructure:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode"`
	Prefix       string `mapstructure:"prefix"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
