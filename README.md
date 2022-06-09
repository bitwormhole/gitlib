# gitlib
gitlib 是一个可扩展的git库，用go实现，完全基于标注库，没有其它依赖


#### 项目内各个包的依赖关系：
    [gitlib] -> [commands] -> [instructions] -> [local] -> [remote] -
    -> [network] -> [repository] -> [git] -> [util]


| 层次 | 名称           | 含义             | 说明               |
| ---- | -------------- | ---------------- | ------------------ |
| 5    | git/context    | 环境（上下文）层 | 提供面向文件的存储 |
| 4    | git/command    | 命令层           | 提供cli接口        |
| 3    | git/service    | 服务层           | 提供高级api接口    |
| 2    | git/repository | 逻辑层           | 提供低级api接口    |
| 1    | git/file       | 物理（持久化）层 | 提供面向文件的存储 |


| 名称        | 含义 | 说明     |
| ----------- | ---- | -------- |
| git/network | 网络 | 网络模块 |
| git/data    | 数据 | 数据模型 |
