package service

type ServiceGroupEnter struct {
	DbService
}

var (
	ServiceGroup = new(ServiceGroupEnter)
)
