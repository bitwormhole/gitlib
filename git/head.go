package git

// HEAD 是表示一个 .git/HEAD 的实体
type HEAD struct {
	Name ReferenceName // the Primary-Key
	ID   ObjectID      // 不常用
}
