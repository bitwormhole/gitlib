# gitlib
一个用go实现的git库


#### 项目内各个包的依赖关系：
    [gitlib] -> [commands] -> [instructions] -> [local] -> [remote] -
    -> [network] -> [repository] -> [git] -> [util]


| 层次 | 名称        | 含义                 | 说明                   |
| ---- | ----------- | -------------------- | ---------------------- |
| 5    | git/cmd     | 命令层               | 提供cli接口            |
| 4    | git/service | 服务层               | 提供api接口            |
| 3    | git/network | 网络层               | 提供面向网络的接口     |
| 2    | git/data    | 数据层               | 提供面向数据结构的操作 |
| 1    | git/fs      | 文件系统（持久化）层 | 提供面向文件的存储     |
