package lib

import (
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/git/support"
	"github.com/bitwormhole/gitlib/git/support/algorithms"
	"github.com/bitwormhole/starter/markup"
)

// ConfigAlgorithms ...
type ConfigAlgorithms struct {
	markup.Component `class:"git-context-configurer"`

	Algorithms []store.AlgorithmRegistry `inject:".git-algorithm-registry"`
}

func (inst *ConfigAlgorithms) _Impl() store.ContextConfigurer {
	return inst
}

// Configure 。。。
func (inst *ConfigAlgorithms) Configure(c *store.Context) error {

	am := &support.AlgorithmManager{}
	am.Init(inst.Algorithms)

	c.Algorithms = inst.Algorithms
	c.AlgorithmManager = am
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// TheSHA1 ...
type TheSHA1 struct {
	markup.Component `class:"git-algorithm-registry"`
	algorithms.DigestSHA1
}

// TheSHA256  ...
type TheSHA256 struct {
	markup.Component `class:"git-algorithm-registry"`
	algorithms.DigestSHA256
}

// TheDeflate  ...
type TheDeflate struct {
	markup.Component `class:"git-algorithm-registry"`
	algorithms.CompressionDeflate
}

// ThePlain  ...
type ThePlain struct {
	markup.Component `class:"git-algorithm-registry"`
	algorithms.CompressionPlain
}
