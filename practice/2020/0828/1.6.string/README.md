# 转换
    * []byte转string需要一次内存拷贝
    * string转[]byte需要一次内存拷贝
# 拼接
    * 空间是一次性分配完成
    * 性能消耗主要在拷贝数据上
# 为什么不能修改？
    * string不包含内存空间，只有一个内存指针
    * string通常指向字符串字面量，字符串字面量存储位置是只读字段
# []byte转换为string不一定会拷贝内存
    * 使用m[string(b)]来查找map（map是string为key，临时把切片b转成string）
    * 字符串拼接，如”<” + “string(b)” + “>”；
    * 字符串比较：string(b) == “foo”
# string和[]byte如何取舍
    * string擅长的场景
        * 需要字符串比较的场景
        * 不需要nil的场景
    * []byte擅长的场景：
        * 修改字符串的场景，尤其是修改粒度为1个字节
        * 函数返回值，需要用nil表示含义的场景
        * 需要切片操作的场景