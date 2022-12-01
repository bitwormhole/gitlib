package pack

// FileType 表示pack文件的类型
type FileType string

// CheckFlag 表示检查标志位, 多个标志位可以通过'|'运算符组合
type CheckFlag int

////////////////////////////////////////////////////////////////////////////////

// 定义各种pack文件的类型
const (
	FileTypeIdx            FileType = ".idx"
	FileTypePack           FileType = ".pack"
	FileTypeRev            FileType = ".rev"
	FileTypeMtimes         FileType = ".mtimes"
	FileTypeMultiPackIndex FileType = "multi-pack-index"
)

// 定义各种检查标志位, 多个标志位可以通过'|'运算符组合
const (
	CheckSize CheckFlag = 0x0008 // 检查文件的大小
	CheckSum  CheckFlag = 0x0004 // 检查文件的 hash-sum
	CheckHead CheckFlag = 0x0002 // 检查文件的头部

	CheckAll CheckFlag = 0xffff // 进行全面的检查
)

// 定义魔数
const (
	MagicNumberIdxV2 = 0xff744f63 // aka `\xfftOc`
)
