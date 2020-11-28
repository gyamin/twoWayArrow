# 開発の始め方

## 開発環境構築

|言語|Version|
|---|---|
|go|1.15.x|


### go
```
$ goenv install 1.15.3
$ goenv local 1.15.3
$ go version
go version go1.15.3 darwin/amd64
```

### go mod
```
$ cd ./src/github.com/gyamin/twoWayArrow

$ go mod init github.com/gyamin/twoWayArrow
go: creating new go.mod: module github.com/gyamin/twoWayArrow

$ cat go.mod 
module github.com/gyamin/twoWayArrow

go 1.15
```

### GoLand
- GOROOT
![](./img/2020-11-22%201.24.52.png)

- Go Modules
![](./img/2020-11-22%201.27.05.png)