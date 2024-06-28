package config

type QQBot struct {
	// 地址
	Address string `mapstructure:"address" yaml:"address"`
	Qq      string `mapstructure:"qq" yaml:"qq"`
	// 激活的群号
	ActiveGroup    []int `mapstructure:"active-group" yaml:"active-group"`
	ActiveGroupMap map[int]bool
}

func (q *QQBot) InitMap() {
	q.ActiveGroupMap = make(map[int]bool)
	for _, item := range q.ActiveGroup {
		q.ActiveGroupMap[item] = true
	}
}

func (q *QQBot) IsActive(item int) bool {
	_, exists := q.ActiveGroupMap[item]
	return exists
}
