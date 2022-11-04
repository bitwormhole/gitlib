package instructions

// Service is the Runner for Command
type Service interface {
	Name() string // return command name
}

// ServiceRegistration ...
type ServiceRegistration struct {
	Name    string
	Service Service
}

// ServiceRegistry ...
// [inject:".git-instruction-registry"]
type ServiceRegistry interface {
	ListRegistrations() []*ServiceRegistration
}

// ServiceManager 用来管理已注册的服务
type ServiceManager interface {
	Find(name string) (Service, error)
}
