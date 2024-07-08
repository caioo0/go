# 排序案例

用命令行实现指定输入数据文件的排序，并输出到指定的文件，要求可以选择算法实现。
程序用法如下：
```shell
USAGE: sorter -i <in> -o <out> -a <qsort|bubblesort>
```
具体的执行结果如下：
```shell
go clean --modcache
export GOPATH="E:\code\go\exec\sorter"  
## windows 
$ set GO111MODULE=auto
$ set GOFLAGS=-mod=vendor
$ set GOPROXY = "https://goproxy.cn"
## linux 
$ export GO111MODULE=on 

$ go build sorter.go 
$ ./sorter -i in.dat -o out.dat -a qsort 
The sorting process costs 10us to complete.
```

演示运行：
```shell 
$ go mod init src 
$ go mod tidy
$ go build src/sorter
$ ./sorter -i in.dat -o out.dat -a qsort

$ go build src/algorithms/bubblesort
$ go build src/algorithms/qsort
$ go install src/algorithms/bubblesort
$ go install src/algorithms/qsort

$ go test src/algorithms/qsort
$ go test src/algorithms/bubblesort
$ go build src/sorter
$ go install src/sorter

$ ./sorter -i in.dat -o sorted.dat -a qsort
```

