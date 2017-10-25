# sunnnyini--ini的处理工具
读取ini文件，读取不同的节(section)——[section]。

读取不同的key，键(key)又名属性(property)，单独占一行用等号连接键名和键值，例如：

    name=value
## 下载安装
   go git github.com/sunnyregion/sunnyini
## 关于注释

下述几种情况的内容将被视为注释：

1. 所有以 `#` 或 `;` 开头的行
2. 所有在 `#` 或 `;` 之后的内容