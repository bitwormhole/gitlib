package git

// ObjectType 表示git对象类型
type ObjectType string

// PackedObjectType 表示pack内对象的类型
type PackedObjectType int

////////////////////////////////////////////////////////////////////////////////

// 定义各种对象类型
const (
	ObjectTypeTag    ObjectType = "tag"
	ObjectTypeCommit ObjectType = "commit"
	ObjectTypeTree   ObjectType = "tree"
	ObjectTypeBLOB   ObjectType = "blob"

	ObjectTypeOfsDelta ObjectType = "OBJ_OFS_DELTA"
	ObjectTypeRefDelta ObjectType = "OBJ_REF_DELTA"
)

// 定义 pack 对象类型
const (
	PackedCommit   PackedObjectType = 1 // OBJ_COMMIT    = (1)
	PackedTree     PackedObjectType = 2 // OBJ_TREE      = (2)
	PackedBLOB     PackedObjectType = 3 // OBJ_BLOB      = (3)
	PackedTag      PackedObjectType = 4 // OBJ_TAG       = (4)
	PackedOfsDelta PackedObjectType = 6 // OBJ_OFS_DELTA = (6)
	PackedRefDelta PackedObjectType = 7 // OBJ_REF_DELTA = (7)
)

////////////////////////////////////////////////////////////////////////////////

func (v ObjectType) String() string {
	return string(v)
}

////////////////////////////////////////////////////////////////////////////////

// ToObjectType ... 转换为 [git.ObjectType] 类型
func (v PackedObjectType) ToObjectType() ObjectType {
	switch v {
	case PackedBLOB:
		return ObjectTypeBLOB
	case PackedCommit:
		return ObjectTypeCommit
	case PackedTag:
		return ObjectTypeTag
	case PackedTree:
		return ObjectTypeTree

	case PackedOfsDelta:
		return ObjectTypeOfsDelta
	case PackedRefDelta:
		return ObjectTypeRefDelta

	default:
		return ""
	}
}

////////////////////////////////////////////////////////////////////////////////
