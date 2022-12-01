package git

// Object 表示一个git对象
type Object struct {
	ID     ObjectID
	Type   ObjectType
	Length int64
}
