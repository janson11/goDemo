# goDemo
## 变量声明
var 变量名 类型 = 表达式

整型
在 Go 语言中，整型分为：

有符号整型：如 int、int8、int16、int32 和 int64。

无符号整型：如 uint、uint8、uint16、uint32 和 uint64。
它们的差别在于，有符号整型表示的数值可以为负数、零和正数，而无符号整型只能为零和正数。

浮点数
浮点数就代表现实中的小数。Go 语言提供了两种精度的浮点数，分别是 float32 和 float64。项目中最常用的是 float64，因为它的精度高，浮点计算的结果相比 float32 误差会更小。

布尔型
一个布尔型的值只有两种：true 和 false，它们代表现实中的“是”和“否”。

字符串
Go 语言中的字符串可以表示为任意的数据，比如以下代码，在 Go 语言中，字符串通过类型 string 声明：


零值
零值其实就是一个变量的默认值，在 Go 语言中，如果我们声明了一个变量，但是没有对其进行初始化，那么 Go 语言会自动初始化其值为对应类型的零值。比如数字类的零值是 0，布尔型的零值是 false，字符串的零值是 "" 空字符串等。

变量
变量简短声明
有没有发现，上面我们演示的示例都有一个 var 关键字，但是这样写代码很烦琐。借助类型推导，Go 语言提供了变量的简短声明 :=，结构如下：

复制代码
变量名:=表达式

指针
在 Go 语言中，指针对应的是变量在内存中的存储位置，也就说指针的值就是变量的内存地址。通过 & 可以获取一个变量的地址，也就是指针。

赋值
在讲变量的时候，我说过变量是可以修改的，那么怎么修改呢？这就是赋值语句要做的事情。最常用也是最简单的赋值语句就是 =

常量的定义
常量的定义和变量类似，只不过它的关键字是 const。

iota
iota 是一个常量生成器，它可以用来初始化相似规则的常量，避免重复的初始化