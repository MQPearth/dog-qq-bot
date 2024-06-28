package config

type Config struct {
	// gorm
	Mysql Mysql `mapstructure:"mysql" yaml:"mysql"`
	// Zap
	Zap Zap `mapstructure:"zap" yaml:"zap"`
	// Server
	Server Server `mapstructure:"server" yaml:"server"`
	// QQBot
	QQBot QQBot `mapstructure:"qq-bot" yaml:"qq-bot"`
	// ai
	Ai Ai `mapstructure:"ai" yaml:"ai"`
}
