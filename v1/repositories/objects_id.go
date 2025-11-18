package repositories

// 表示一个对象ID
type ObjectID []byte

// 表示一个包文件的 ID
type PackID []byte

type PackIndex interface {
	GetPackID() PackID
}
