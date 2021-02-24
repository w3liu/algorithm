* cgo命令会为每个包包含了cgo代码的Go文件创建两个中间文件
* 比如`main.go`会分别会分别创建`main.cgo1.go`和`main.cgo2.c`两个中间文件
* 并为整个包创建一个`_cgo_gotypes.go`Go文件，其中包含Go语言部分辅助代码
* 此外，还会创建一个`_cgo_export.h`和 `_cgo_export.c`文件，对应Go语言导出到C语言的类型和函数
* 编译可以供C语言调用的静态库命令：`go build -buildmode=c-archive -o sum.a sum.go` 