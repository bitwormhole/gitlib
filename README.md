# gitlib
gitlib 是一个可扩展的git库，用go实现。


#### 项目内各个包的依赖关系：
    [gitlib] -> [support] -> [network] -> [commands] -> [instructions] -> [store] -> [git]


| 层次 | 名称             | 含义     | 说明                             |
| ---- | ---------------- | -------- | -------------------------------- |
| 7    | .                | gitlib   | 提供面向[模块]的接口             |
| 6    | git/support      | 实现     | 提供对各种接口的实现             |
| 5    | git/network      | 网络     | 提供面向 [网络] 的接口           |
| 4    | git/commands     | 命令     | 提供面向 [command] 的接口（cli） |
| 3    | git/instructions | 指令框架 | 提供面向 [指令] 的接口           |
| 2    | git/store        | 存储库   | 提供面向 [repo] 的接口           |
| 1    | git/.            | 数据对象 | 基本的git数据对象                |
