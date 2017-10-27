# sunnnyini--ini的处理工具

![ini logo](./ini.png)

读取ini文件，读取不同的节(section)——[section]。

读取不同的key，键(key)又名属性(property)，单独占一行用等号连接键名和键值，例如：

    name=value
## 下载安装
   go git github.com/sunnyregion/sunnyini
## 关于注释

下述几种情况的内容将被视为注释：

1. 所有以 `#` 或 `;` 开头的行
2. 所有在 `#` 或 `;` 之后的内容
3. 在行尾以"//"结尾的部分



## git 标签的使用方法

- :art: when improving the format/structure of the code <br />
当代码提高format/structure
- :rocket: when improving performance <br />
当代码提高性能
- :pencil2: when writing docs <br />
写文档
- :bulb: new idea <br />
新的想法
- :construction: work in progress <br />
工作的进展
- :heavy_plus_sign: when adding feature <br />
  添加功能
- :heavy_minus_sign: when removing feature <br />
移除功能
- :speaker: when adding logging <br />
增加日志
- :mute: when reducing logging <br />
减少日志
- :bug: when fixing a bug <br />
修改bug
- :white_check_mark: when adding tests <br />
增加测试
- :lock: when dealing with security <br />
安全处理
- :arrow_up: when upgrading dependencies <br />
依赖升级
- :arrow_down: when downgrading dependencies <br />
降低依赖