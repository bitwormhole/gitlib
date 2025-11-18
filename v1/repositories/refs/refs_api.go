package refs

import "github.com/bitwormhole/gitlib/v1/repositories"

// 这个包提供一组针对 refs 的类型.
// 它们的实际类型均来自于 repositories 包中的 refs 部分.

// string

type Name = repositories.Refname

// interface

type Refs = repositories.Refs

type Ref = repositories.Ref
