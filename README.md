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

- :art: when improving the format/structure of the code
- :rocket: when improving performance
- :pencil2: when writing docs
- :bulb: new idea
- :construction: work in progress
- :heavy_plus_sign: when adding feature
- :heavy_minus_sign: when removing feature
- :speaker: when adding logging
- :mute: when reducing logging
- :bug: when fixing a bug
- :white_check_mark: when adding tests
- :lock: when dealing with security
- :arrow_up: when upgrading dependencies
- :arrow_down: when downgrading dependencies