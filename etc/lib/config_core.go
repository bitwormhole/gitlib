package lib

import (
	"github.com/bitwormhole/gitlib/git/gitconfig"
	"github.com/bitwormhole/gitlib/git/store"
	"github.com/bitwormhole/gitlib/git/support"
	"github.com/bitwormhole/gitlib/git/support/algorithms"
	"github.com/bitwormhole/gitlib/git/support/sessions"
	"github.com/bitwormhole/starter/markup"
)

// ConfigCore ...
type ConfigCore struct {
	markup.Component `class:"git-core-configurer"`
}

func (inst *ConfigCore) _Impl() store.CoreConfigurer {
	return inst
}

// Configure 。。。
func (inst *ConfigCore) Configure(c *store.Core) error {
	steps := make([]func(c *store.Core) error, 0)
	steps = append(steps, inst.initBase)
	steps = append(steps, inst.initSessionFactory)
	steps = append(steps, inst.initCompression)
	steps = append(steps, inst.initDigest)
	steps = append(steps, inst.initPathMapping)
	for _, step := range steps {
		err := step(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *ConfigCore) initBase(c *store.Core) error {
	c1 := &support.BaseCoreConfigurer{}
	return c1.Configure(c)
}

func (inst *ConfigCore) initCompression(c *store.Core) error {
	name := inst.getConfigProperty(c, gitconfig.CoreCompressionAlgorithm)
	am := c.Context.AlgorithmManager
	alg, err := am.FindCompression(name)
	if err != nil {
		return err
	}
	c.Compression = alg
	return nil
}

func (inst *ConfigCore) initDigest(c *store.Core) error {
	name := inst.getConfigProperty(c, gitconfig.CoreDigestAlgorithm)
	am := c.Context.AlgorithmManager
	alg, err := am.FindDigest(name)
	if err != nil {
		return err
	}
	c.Digest = alg
	return nil
}

func (inst *ConfigCore) initPathMapping(c *store.Core) error {
	pattern := inst.getConfigProperty(c, gitconfig.CoreObjectsPathPattern)
	pathmap1 := algorithms.PathPatternMapping{}
	pathmap2 := pathmap1.WithPattern(pattern)
	c.PathMapping = pathmap2
	return nil
}

func (inst *ConfigCore) getConfigProperty(c *store.Core, key gitconfig.KeyTemplate) string {
	cfg := c.Config.Mix().Config()
	value := cfg.GetProperty(key.String())
	if value == "" {
		panic("no config property: " + key)
	}
	return value
}

func (inst *ConfigCore) initSessionFactory(c *store.Core) error {
	c.SessionFactory = &sessions.Factory{}
	return nil
}
