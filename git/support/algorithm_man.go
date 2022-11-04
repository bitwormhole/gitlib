package support

import (
	"errors"
	"strings"

	"github.com/bitwormhole/gitlib/git/store"
)

// AlgorithmManager ...
type AlgorithmManager struct {
	table map[string]*store.AlgorithmRegistration
}

func (inst *AlgorithmManager) _Impl() store.AlgorithmManager {
	return inst
}

// Init ...
func (inst *AlgorithmManager) Init(src []store.AlgorithmRegistry) {
	dst := inst.table
	dst = make(map[string]*store.AlgorithmRegistration)
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
func (inst *AlgorithmManager) Find(name string) (store.Algorithm, error) {
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
func (inst *AlgorithmManager) FindCompression(name string) (store.Compression, error) {
	a, err := inst.Find(name)
	if err != nil {
		return nil, err
	}
	p := a.(store.Compression)
	return p, nil
}

// FindDigest ...
func (inst *AlgorithmManager) FindDigest(name string) (store.Digest, error) {
	a, err := inst.Find(name)
	if err != nil {
		return nil, err
	}
	p := a.(store.Digest)
	return p, nil
}

// FindPathMapping 。。。
func (inst *AlgorithmManager) FindPathMapping(name string) (store.PathMapping, error) {
	a, err := inst.Find(name)
	if err != nil {
		return nil, err
	}
	p := a.(store.PathMapping)
	return p, nil
}
