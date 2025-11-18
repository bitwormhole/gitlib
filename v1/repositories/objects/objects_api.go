package objects

import "github.com/bitwormhole/gitlib/v1/repositories"

// 这个包提供一组针对 objects 的类型.
// 它们的实际类型均来自于 repositories 包中的 objects 部分.

// ids

type ObjectID = repositories.ObjectID

type PackID = repositories.PackID

// interface

type SparseObject = repositories.SparseObject
