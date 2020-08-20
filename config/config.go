package config

type Service struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
}

type System struct {
	Name  string `yaml:"name"`
	Url   string `yaml:"url"`
	Addr  string `yaml:"addr"`
	Env   string `yaml:"env"`
	Debug bool   `yaml:"debug"`
}

type Log struct {
	Prefix  string `yaml:"prefix"`
	LogFile bool   `yaml:"log-file"`
	Stdout  string `yaml:"stdout"`
	File    string `yaml:"file"`
}

type Mysql struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Addr         string `yaml:"addr"`
	Database     string `yaml:"database"`
	Config       string `yaml:"config"`
	MaxIdleConns int    `yaml:"max-idle-conns"`
	MaxOpenConns int    `yaml:"max-open-conns"`
	LogMode      bool   `yaml:"log-mode"`
	Prefix       string `yaml:"prefix"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
