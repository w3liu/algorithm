# 命令

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