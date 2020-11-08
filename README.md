# 目录
- algorithm 算法练习
- ARTS 每周打卡任务
- books 读书笔记
- diagram 图表
- practice 练习
- tools 工具

# 常用命令

## benchmark
```text
go test -bench=. -benchtime=1s -run=none -benchmem
```

## env
```text
go env -w GOPROXY=https://goproxy.io,direct
go env -w GOPRIVATE=*.corp.example.com
go env -w GO111MODULE=on
```