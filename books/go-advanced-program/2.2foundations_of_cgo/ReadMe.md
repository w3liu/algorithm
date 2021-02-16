* windows下需要安装MinGW工具，需要保证CGO_ENABLED被设置为1
* 当交叉构建时CGO默认时禁止的，比如要交叉构建ARM环境运行的Go程序，需要手工设置好C/C++交叉构建的工具链，同时开启CGO_ENABLED环境变量
* 在`import "C"`语句钱的注释可以通过 `#cgo`语句设置编译阶段和链接阶段的相关参数
* `#cgo CFLAGS: -DPNG_DEBUG=1 -I./include` `CFLAGS` 表示指定编译阶段参数，`-D`部分定义了宏PNG_DEBUG，值为1；-I定义了头文件包含的检索目录
* `#cgo LDFLAGS: -L/usr/local/lib -lpng` `LDFLAGS` 表示指定链接阶段参数，`-L`指定了链接时库文件检索目录，`-l` 指定了链接时需要链接png库
* 因为C/C++遗留问题，C头文件检索目录可以是相对目录，但是库文件检索目录则需要绝对路径
* 在库文件的检索目录中可以通过`${SRCDIR}`变量表示当前包目录的绝对路径：`// #cgo LDFLAGS: -L${SRCDIR}/libs -lfoo`
* `#cgo` 语句主要影响CFLAGS、CPPFLAGS、CXXFLAGS、FFLAGS、LDFLAGS几个编译器环境变量
* `// +build debug` 表示只有在设置debug构建标志时才会被构建，例如：`go build -tags="debug"`
* `-tags` 命令行参数同时指定多个build标志，它们之间用逗号分割
* `// +build linux,386 darwin,!cgo` 其中linux,386中linux和386用逗号拼接表示AND的意思；而linux,386和darwin,!cgo之间通过空格来分割来表示OR的意思