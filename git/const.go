package git

// ObjectType 表示git对象类型
type ObjectType string

// 定义各种对象类型
const (
	ObjectTypeTag    ObjectType = "tag"
	ObjectTypeCommit ObjectType = "commit"
	ObjectTypeTree   ObjectType = "tree"
	ObjectTypeBLOB   ObjectType = "blob"
)

////////////////////////////////////////////////////////////////////////////////

func (v ObjectType) String() string {
	return string(v)
}
