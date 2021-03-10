# log

## 简介

当前，go 本身并没有良好的错误处理机制

然后层层传递，最终将错误传递到最上层，这里面存在着两个问题：

- 没有错误发生时的上下文信息（或者叫堆栈信息）
- 在层层的错误传递过程中，有可能已经将原始错误转化，丢失了最原始的 error

当前，比较优雅的方式是使用 github.com/pkg/errors 这个包来进行错误处理

- 在错误的发生点使用errors.New生成错误，或者使用errors.Wrap封装错误，这时候会纪录堆栈信息，在之后就不需要在添加堆栈信息了
- 如果需要对错误添加错误信息，使用errors.WithMessage方法
- 顶层使用errors.Cause获取原始错误信息，使用%+v纪录整条链路的错误信息和堆栈信息

## 举例

## 参考资料

1. [使用 pkg/errors 进行错误处理](http://wangtingkui.com/2020/04/errors/)
2. [Golang 错误处理最佳实践](https://medium.com/@dche423/golang-error-handling-best-practice-cn-42982bd72672)
3. [doc](https://godoc.org/github.com/pkg/errors)
4. [git](https://github.com/pkg/errors)
