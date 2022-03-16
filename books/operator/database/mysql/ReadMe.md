# MySQL运维手册
## MySQL安装
1. 下载MySQL yum包
```
wget http://repo.mysql.com/mysql57-community-release-el7-10.noarch.rpm
```
2. 安装MySQL源
```
rpm -Uvh mysql57-community-release-el7-10.noarch.rpm
```
3. 安装MySQL服务端
```
rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2022
yum install -y mysql-community-server
```
4. 启动MySQL
```
systemctl start mysqld.service
```
5. 检查是否启动成功
```
systemctl status mysqld.service
```
6. 获取临时密码，MySQL5.7为root用户随机生成了一个密码
```
grep 'temporary password' /var/log/mysqld.log 
```
7. 通过临时密码登录MySQL，进行修改密码操作
```
mysql -uroot -p
```
```
alter user 'root'@'localhost' IDENTIFIED BY 'yourpassword';
```
8. 授权其他机器远程登录
```
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'yourpassword' WITH GRANT OPTION;
 
FLUSH PRIVILEGES;
```
9. 开启开机自启动
```
systemctl enable mysqld
systemctl daemon-reload
```
10. 防火墙开放3306端口
```
firewall-cmd --state
firewall-cmd --zone=public --add-port=3306/tcp --permanent
firewall-cmd --reload
```
