# jenkins运维手册
## 安装jenkins
1. 检查是否有jdk环境
```
java -version
```
如果没有的话，先安装
```
yum install java-1.8.0-openjdk* -y
```
-y:会自动安装

默认安装位置：/usr/lib/jvm

2. 安装git插件
```
yum install git -y
```
安装完，检查
```
git --version
```
3. 安装jenkins
```
sudo wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat-stable/jenkins.repo
```
```
sudo rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io.key
```
```
yum install jenkins
```
4. 安装完，修改配置，比如修改用户名、端口等
```
cd /etc/sysconfig
vi jenkins
```
打开后，编辑你需要的信息，这里以修改用户名和端口为例
```
JENKINS_USER="root"
JENKINS_PORT="8888"
```