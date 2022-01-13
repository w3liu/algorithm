# ClickHouse运维手册
## 安装ClickHouse
1. 系统要求
```
grep -q sse4_2 /proc/cpuinfo && echo "SSE 4.2 supported" || echo "SSE 4.2 not supported"
```
2. rpm安装包
首先，添加官方存储库
```
sudo yum install yum-utils
sudo rpm --import https://repo.clickhouse.com/CLICKHOUSE-KEY.GPG
sudo yum-config-manager --add-repo https://repo.clickhouse.com/rpm/stable/x86_64
```
然后运行命令安装：
```
sudo yum install clickhouse-server clickhouse-client
```
3. 启动
如果没有service，可以运行如下命令在后台启动服务：
```
sudo /etc/init.d/clickhouse-server start
```
日志文件将输出在/var/log/clickhouse-server/文件夹。
如果服务器没有启动，检查/etc/clickhouse-server/config.xml中的配置。
您也可以手动从控制台启动服务器:
```
clickhouse-server --config-file=/etc/clickhouse-server/config.xml
```
4. 命令行客户端
启动服务后，可以使用命令行客户端连接到它：
```
clickhouse-client
```
默认情况下，使用default用户并不携带密码连接到localhost:9000
5. 官方文档
```
https://clickhouse.com/docs/zh/getting-started/install/
```
## DBeaver
1. 下载地址
```
https://dbeaver.io/download/
```
2. 配置连接
驱动设置：
```
ru.yandex.clickhouse.ClickHouseDriver
```
JDBC URL:
```
jdbc:clickhouse://localhost:9000/default
```



