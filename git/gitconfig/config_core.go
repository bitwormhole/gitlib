package gitconfig

// 定义仓库配置名称 (core.*)
const (
	CoreLogAllRefUpdates        KeyTemplate = "core.logAllRefUpdates"
	CoreSymlinks                KeyTemplate = "core.symlinks"
	CoreIgnoreCase              KeyTemplate = "core.ignoreCase"
	CoreBare                    KeyTemplate = "core.bare"
	CoreRepositoryFormatVersion KeyTemplate = "core.repositoryFormatVersion"
	CoreFileMode                KeyTemplate = "core.fileMode"

	CoreCompressionAlgorithm KeyTemplate = "core.compressionAlgorithm"
	CoreDigestAlgorithm      KeyTemplate = "core.digestAlgorithm"
	CoreObjectsPathPattern   KeyTemplate = "core.objectsPathPattern"
)

// Core ...
type Core struct {
	Exists bool

	Bare                    bool
	RepositoryFormatVersion int
}
