package config

//定义config类型接收yaml数据

type ServerConfig struct {
	Name        string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	Mysqlinfo   MysqlConfig `mapstructure:"mysql"`
	LogsAddress string      `mapstructure:"logsAddress"`
	JWTKey      JwtConfig   `mapstructure:"jwt"`
	MinioInfo   minIO       `mapstructure:"minio"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type JwtConfig struct {
	SigningKey string `mapstructure:"key" json:"key" `
}

type minIO struct {
	Endpoint string `mapstructure:"endpoint"`
	ID       string `mapstructure:"ID"`
	Key      string `mapstructure:"Key"`
}
