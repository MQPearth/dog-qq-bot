package config

type Ai struct {
	AccessKey string `mapstructure:"access-key" yaml:"access-key"`
	SecretKey string `mapstructure:"secret-key" yaml:"secret-key"`
}
