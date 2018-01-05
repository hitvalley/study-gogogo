# fmt

> fmt 包是用于格式化输入和输出，类似于 C 语言的 scanf 和 printf。fmt 的 verbs 和 C 类似，但是相对简单。

[fmt document](https://golang.org/pkg/fmt/)

## 引入 fmt

```go
import "fmt"
```

## Verbs

### 通用

* %v    展示值的默认形式
* %#v   以 Go 的方式展示数据
* %T    以 Go 的方式展示数据类型
* %%    百分号，不做转义

### 布尔

* %t    true 或 false

### 整形数字

* %b    数字的二进制展示
* %c    数字对应的Unicode值
* %d    数字的十进制展示
* %o    数字的八进制展示
* %q    用单引号包裹的Unicode值
* %x    数字十六进制展示，同时，a-f 是小写
* %X    数字十六进制展示，同时，A-F 是大写
* %U    Unicode 格式：U+1234

### 浮点型

* %b    浮点数的科学计数法展示
