package repository

import "github.com/bitwormhole/starter/collection"

// 配置的作用域(scope)
const (
	ConfigForDefault = iota
	ConfigForSystem
	ConfigForUser
	ConfigForRepository
	ConfigForSolid
	ConfigForHybrid
)

// Config the file of [.git/config]
type Config interface {
	GetProperties(scope int) collection.Properties
	SaveProperties(scope int) error
}
