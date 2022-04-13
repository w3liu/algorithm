# Nginx运维手册
## 源码安装Nginx
1. 下载nginx
```
wget http://nginx.org/download/nginx-1.20.2.tar.gz
```
2. 安装依赖环境
```
yum install gcc-c++
yum install -y pcre pcre-devel
yum install -y openssl openssl-devel
```
3. 解压
```
tar -zxvf  nginx-1.20.2.tar.gz
```
4. 创建nginx临时目录
```
mkdir /var/temp/nginx -p
```
5. 在nginx的目录，输入如下命令进行配置，目的是为了创建Makefile文件
```
./configure \
--prefix=/usr/local/nginx \
--pid-path=/var/run/nginx/nginx.pid \
--lock-path=/var/lock/nginx.lock \
--error-log-path=/var/log/nginx/error.log \
--http-log-path=/var/log/nginx/access.log \
--with-http_gzip_static_module \
--http-client-body-temp-path=/var/temp/nginx/client \
--http-proxy-temp-path=/var/temp/nginx/proxy \
--http-fastcgi-temp-path=/var/temp/nginx/fastcgi \
--http-uwsgi-temp-path=/var/temp/nginx/uwsgi \
--http-scgi-temp-path=/var/temp/nginx/scgi
```
## yum安装Nginx
1. 安装epel-release源
```
yum -y install epel-release
```
2. 设置nginx安装源
```
vi /etc/yum.repos.d/nginx.repo

[nginx-stable]
name=nginx stable repo
baseurl=http://nginx.org/packages/centos/7/$basearch/
gpgcheck=1
enabled=1
gpgkey=https://nginx.org/keys/nginx_signing.key
```
3. 安装nginx
```
yum install nginx
```
4. 查看版本
```
nginx -v
```
5. 查看编译参数
```
nginx -V
```
6. 查看安装目录
```
rpm -ql nginx
```
7. 查看配置文件
```
more /etc/nginx/nginx.conf
```
8. 创建nginx服务
```
systemctl enable nginx
```
9. 启动nginx服务
```
systemctl start nginx
```
10. 停止nginx服务
```
systemctl stop nginx
```
11. 重载nginx服务
```
nginx -s reload /etc/nginx/nginx.conf
```