# gitlib
gitlib 是一个可扩展的git库，用go实现。


#### 项目内各个包的依赖关系：
    [gitlib] -> [support] -> [store] -> [network]  -> [commands] -> [git] -> [instructions] -> [data]


| 层次 | 名称             | 含义     | 说明                             |
| ---- | ---------------- | -------- | -------------------------------- |
| 4    | .                | gitlib   | 提供面向模块的接口               |
| 2    | git/support      | 实现     | 提供对各种接口的实现             |
| 4    | git/store        | 仓库     | 提供面向 [repo] 的接口           |
| 4    | git/network      | 网络     | 提供面向 [网络] 的接口           |
| 4    | git/commands     | 命令     | 提供面向 [command] 的接口（cli） |
| 1    | git/.            | 指令api  | 提供面向 [指令] 的接口           |
| 4    | git/instructions | 指令框架 |                                  |
| 4    | git/data         | 数据对象 |                                  |
