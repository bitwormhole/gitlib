package identity

import "github.com/bitwormhole/gitlib/git"

var theSha1IDFactory *sha1IDFactory
var theSha256IDFactory *sha256IDFactory

// GetSha1IDFactory 取sha-1-id 工厂（单例）
func GetSha1IDFactory() git.IdentityFactory {
	inst := theSha1IDFactory
	if inst == nil {
		inst = &sha1IDFactory{}
		theSha1IDFactory = inst
	}
	return inst
}

// GetSha256IDFactory 取sha-256-id 工厂（单例）
func GetSha256IDFactory() git.IdentityFactory {
	inst := theSha256IDFactory
	if inst == nil {
		inst = &sha256IDFactory{}
		theSha256IDFactory = inst
	}
	return inst
}
