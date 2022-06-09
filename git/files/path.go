package files

// Path 表示一个文件系统绝对路径
type Path interface {
	FS() FS
	String() string
	Name() string
	IsDir() bool
	IsFile() bool
	Exists() bool
	List() []string
	Parent() Path
	Child(name string) Path
}

// FS 表示一个抽象的文件系统
type FS interface {
	Path(p string) Path
}
