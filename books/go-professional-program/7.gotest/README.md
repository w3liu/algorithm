# 单元测试
* 测试文件名必须以”_test.go”结尾
* 测试函数名必须以“TestXxx”开始
* 命令行下使用”go test”即可启动测试

# 性能测试
* 测试文件名必须以”_test.go”结尾
* 函数名必须以“BenchmarkXxx”开始
* 使用命令“go test -bench=.”即可开始性能测试

# 示例测试
* 例子测试函数名需要以”Example”开头
* 检测单行输出格式为“// Output: <期望字符串>”
* 检测多行输出格式为“// Output: \ <期望字符串> \ <期望字符串>”，每个期望字符串占一行
* 检测无序输出格式为”// Unordered output: \ <期望字符串> \ <期望字符串>”，每个期望字符串占一 行
* 测试字符串时会自动忽略字符串前后的空白字符
* 如果测试函数中没有“Output”标识，则该测试函数不会被执行
* 执行测试可以使用 go test ，此时该目录下的其他测试文件也会一并执行
* 执行测试可以使用 go test <xxx_test.go> ，此时仅执行特定文件中的测试函数

# 子测试
* 子测试适用于单元测试和性能测试
* 子测试可以控制并发
* 子测试提供一种类似table-driven风格的测试
* 子测试可以共享setup和tear-down

# main测试
* 用于主动执行各种测试，可以测试前后做setup和tear-down操作
