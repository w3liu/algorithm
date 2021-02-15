* 通过`improt "C"`语句开启CGO特性
* `/**/`中间是C代码,之后接 import "C" 如果存在空行 就会报错.could not determine kind of name for C.*
* CGO不仅仅用于Go语言中调用C语言函数，还可以用于导出Go语言函数给C语言函数调用
