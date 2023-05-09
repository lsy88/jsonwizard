package config

type MongoDB struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`             //服务器地址:端口
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             //:端口
	Dbname   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // 数据库名
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	LogZap   bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`    // 是否通过zap写入日志文件
}

func (m *MongoDB) Dsn() string {
	return "mongodb://" + m.Username + ":" + m.Password + "@" + m.Path + ":" + m.Port + "/" + m.Dbname
}
