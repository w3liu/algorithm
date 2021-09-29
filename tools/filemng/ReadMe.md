## 编译命令
```
set GOARCH=amd64
set GOOS=linux
go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'" -o filemng
```

### 启动命令
```
chmod +x filemng
nohup ./filemng >> nohup.log 2>&1 &
```