package support

import (
	"errors"
	"strings"

	"github.com/bitwormhole/gitlib/git"
	"github.com/bitwormhole/gitlib/git/store"
)

// AlgorithmManager ...
type AlgorithmManager struct {
	table map[string]*git.AlgorithmRegistration
}

func (inst *AlgorithmManager) _Impl() store.AlgorithmManager {
	return inst
}

// Init ...
func (inst *AlgorithmManager) Init(src []git.AlgorithmRegistry) {
	dst := inst.table
	dst = make(map[string]*git.AlgorithmRegistration)
	for _, ar1 := range src {
		some := ar1.ListRegistrations()
		for _, ar2 := range some {
			name := inst.normalizeName(ar2.Name)
			dst[name] = ar2
		}
	}
	inst.table = dst
}

func (inst *AlgorithmManager) normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return name
}

// Find ...
func (inst *AlgorithmManager) Find(name string) (git.Algorithm, error) {
	name = inst.normalizeName(name)
	ar := inst.table[name]
	if ar != nil {
		p := ar.Provider
		if p != nil {
			return p, nil
		}
	}
	return nil, errors.New("no algorithm with name:" + name)
}

// FindCompression ...
func (inst *AlgorithmManager) FindCompression(name string) (git.Compression, error) {
	a, err := inst.Find(name)
	if err != nil {
		return nil, err
	}
	p := a.(git.Compression)
	return p, nil
}

// FindDigest ...
func (inst *AlgorithmManager) FindDigest(name string) (git.Digest, error) {
	a, err := inst.Find(name)
	if err != nil {
		return nil, err
	}
	p := a.(git.Digest)
	return p, nil
}

// FindPathMapping 。。。
func (inst *AlgorithmManager) FindPathMapping(name string) (git.PathMapping, error) {
	a, err := inst.Find(name)
	if err != nil {
		return nil, err
	}
	p := a.(git.PathMapping)
	return p, nil
}
