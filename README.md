# 开发环境

go1.9.2

GOOS="darwin"

GOARCH="amd64"

GoLand

# 运行说明

## 1、示例测试

```
go test main_test.go -v
```

## 2、终端交互运行

```
go run main.go -c ./testdata/input.txt

-c 配置输入文件路径
```

或者

```
go run main.go
```

两种运行后，都可在终端手动输入信息与系统交互。


# 设计说明

## 具体分析：

 * 罗马数字与十进制转换作为计算资源独立出来。抽象成Calculator，适配其它计算资源。
 * 解析输入信息，输出处理结果。抽象成Processor。
	 * Galaxy中输入信息都有独立规则，通过构造hander数组，遍历匹配处理。
	 * Galaxy中输入信息可分为Command与Question，Command会更新GalaxyGuider内部状态，Question需要Command提供的信息进行处理。
	 * 计算Command里信息，需要提供计算资源，通过注入Calculator到GalaxyGuider
 * 把问题处理抽象成一个Job，每种Worker包含一个Processor，通过调度器，实现不同种类信息的同时处理。

## 整理设计

流程：Console -> Dispatcher -> Worker(Processor[Calculator]) -> Console

* 用户通过Console输入信息到Dispatcher。
* Dispatcher将不同种类的信息分发到对应的Worker处理，每个Worker都有对应的Processor，每个Processor都注入了Calculator提供计算。
* Worker将处理结果输出到Console终端展示。
