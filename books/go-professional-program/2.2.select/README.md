# 总结
    * select语句中除default外，每个case操作一个channel，要么读要么写
    * select语句中除default外，各case执行顺序是随机的
    * select语句中如果没有default语句，则会阻塞等待任一case
    * select语句中读写操作要判断是否成功读取，关闭的channel也可以读取