package models

// MongoConf mongo的配置
type mongoConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Config 主配置
type Config struct {
	MongoConfig mongoConf `yaml:"mongo"`
	Env         string    `yaml:"env"`
	JwtKey      string    `yaml:"jwt-key"`
}
