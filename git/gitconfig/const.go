// const.go
// 这里定义了一些 .git/config 属性名称的常量
// 参考 https://github.com/git/git/blob/master/Documentation/config/

package gitconfig

// KeyTemplate 表示仓库配置的属性名模板
type KeyTemplate string

// NamePlaceholder 名称占位符
const NamePlaceholder = "<name>"
