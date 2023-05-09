package config

type Server struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Local   Local   `mapstructure:"local" json:"local" yaml:"local"`
	MongoDB MongoDB `mapstructure:"mongodb" json:"mongodb" yaml:"mongodb"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
