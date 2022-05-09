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

7. 优化性能的配置
a. Nginx worker进程个数
```
# 语法
worker_processes number;
# 默认
worker_processes 1;
```
在master/worker运行方式下，定义worker进程的个数。
如果模块调用不会出现阻塞式的调用，那么，有多少CPU内核就应该配置多少个进程；
反之，如果可能出现阻塞式调用，那么需要配置稍多一些的worker进程。
例如如果有大量的磁盘I/O操作的业务可以配置多一些的worker进程。
b. 绑定Nginx worker进程到指定的CPU内核
```
# 语法
worker_cpu_affinity cpumask[cpumask...]
```
例如，如果有4颗CPU内核，就可以进行如下配置：
```
worker_processes 4;
worker_cpu_affinity 1000 0100 0010 0001;
```
c. ssl硬件加速
```
# 语法
ssl_engine device;
```
d. 系统调用gettimeofday的执行频率
```
# 语法
timer_resolution t;
# 示例
timer_resolution 100ms;
```
e. Nginx worker进程优先级设置
```
# 语法
worker_priority nice;
# 默认
worker_priority 0;
```
该配置用于设置Nginx worker进程的nice优先级，nice的取值范围是：-20~+19，-20是最高优先级。
如果希望Nginx占有更多的系统资源，可以把nice值配置得更小一些，但不建议比内核进程的nice值（通常为-5）还要小。
8. 事件类配置项
a. 是否打开accept锁
```
# 语法
accept_nutex[on|off]
# 默认
accept_mutex on;
```
b. lock文件路径
```
# 语法
lock_file path/file;
# 默认
lock_file logs/nginx.lock;
```
c. 使用accept锁后到真正建立连接之间的延迟时间
```
# 语法
accept_mutex_delay Nms;
# 默认
accept_mutex_delay 500ms;
```
d. 批量建立新连接
```
# 语法
multi_accept[on|off];
# 默认
multi_accept off;
```
当时间模型通知有新连接时，尽可能地对本次调度中客户端发起的所有TCP请求都建立连接。
e. 选择事件模型
```
# 语法
use[kqueue|rtsig|epoll|/dev/poll|select|poll|eventport];
# 默认
Nginx会自动使用最合适的事件模型。
```
Linux系统可以选择的事件驱动模型有poll、select、epoll三种。其中，epoll性能最高。
f. 每个worker的最大连接数
```
# 语法
worker_connections number;
```
定义每个worker进程可以同时处理的最大连接数。