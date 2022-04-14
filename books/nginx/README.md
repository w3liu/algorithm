# Nginx学习笔记
## 安装Nginx的准备工作
* Linux操作系统
```
uname -a
```
* 使用nginx的必备软件
    - gcc 可以用来编译C语言程序
    - gcc-c++ 可以用来编译C++程序
    - pcre 支持正则表达式的库
    - pcre-devel 使用pcre做二次开发时所需要的开发库
    - zlib 用于对HTTP包的内容做gzip格式的压缩
    - zlib-devel 使用zlib做二次开发时所需要的开发库
    - openssl 支持ssl协议的库
    - openssl-devel 使用openssl做二次开发时所需要的开发库
* 内核参数优化
```
cat /etc/sysctl.conf
```
```
file-max:这个参数表示进城可以同时打开的最大句柄数
tcp_tw_reuse:这个参数设置为1，表示允许TIME_WAIT状态的socket重新用于新的连接
tcp_keepalive_time:这个参数表示当keepalive启用时，TCP发送keepalive消息的频度
tcp_fin_timeout:这个参数表示当服务器主动关闭连接时，socket保持在FIN-WAIT-2状态的最大时间
tcp_max_tw_buckets:这个参数表示操作系统允许TIME_WAIT套接字数量的最大值
tcp_max_syn_backlog:这个参数表TCP三次握手建立阶段接收SYN请求队列的最大长度
ip_local_port_range:这个参数定义了在UDP和TCP连接中本地端口的取值范围
net.ipv4.tcp_rmem:这个参数定义了TCP接收缓存
net.ipv4.rcp_wmem:这个参数定义了TCP发送缓存
netdev_max_backlog:这个参数表示网卡接收数据包队列的最大值
rmem_default:这个参数表示内核套接字接收缓存区默认的大小
wmem_default:这个参数表示内核套接字接发送存区默认的大小
rmem_max:这个参数表示内核套接字接收缓存区最大大小
wmem_max:这个参数表示内核套接字发送缓存区最大大小
```

## Nginx的配置
1. 配置项的语法格式
```
配置项名
    配置项值
    ...
```
2. 如果配置项值中包括语法符号，需要使用单引号或双引号括住配置项值，例如：
```
log_format main '$remote_addr - $remote_user [$time_local] "$request"'
```
3. 如果一个配置项暂时许愿哦注释掉，那么可以加“#”注释掉这一行配置，例如：
```
#pid    logs/nginx.pid;
```
4. 配置项的单位
a. 指定空间大小时，可以使用的单位：
- `k`或者`K`
- `M`或者`m`
b. 指定时间时，可以使用的单位：
- `ms`，`s`，`m`，`h`，`d`，`w`，`M`，`y`
5. 在配置中使用变量，例如：
```
log_format main '$remote_addr - $remote_user "$request"';
```
6. Nginx服务的配置分类：
- 用于调试、定位问题的配置
- 正常运行的必备配置项
- 优化性能的配置项
- 事件类配置项
7. 用于调试、定位问题的配置项
a. 是否以守护进程方式运行Nginx
- 语法：daemon on | off;
- 默认：daemon on;
b. 是否以master/worker方式工作
- 语法：master_process on | off;
- 默认：master_process on;
c. 日志的设置
- 语法：error_log pathfile level;
- 默认：error_log logs/error.log error;
d. 是否处理几个特殊的调试定
- 语法：`debug_points[stop|abort]`
```
这个配置项用来帮助用户跟踪调试Nginx
如果设置为stop，那么Nginx的代码执行到这些调试点的时候就会发出SIGSTOP信号以用于调试
如果设置为abort，那么Nginx则会产生一个coredump文件，可以使用gdb来查看Nginx当时的各种信息

PS：通常不会使用这个配置项
```
e. 仅对指定的客户端输出debug级别的日志
- 语法：`debug_connection[IP|CIDR]`
```
这个配置项实际上属于事件类配置，因此，它必须放在events{···}中才有效。
```
它的值可以是IP地址或CIDR地址，例如：
```
events {
    debug_connection 10.224.66.14;
    debug_connection 10.244.57.0/24;
}
```
f. 限制coredump核心转存储文件的大小
- 语法：worker_rlimit_core size;
```
在Linux系统中，当进程发生错误或收到信号而终止时，系统会将进程执行时的内存内容（核心映像）写入一个文件（core文件），以作为调试之用。
由于core文件可能会占用很大空间，因此需要通过该配置加以限制。
```
g. 指定coredump文件生成目录
- 语法：working_directory path;
