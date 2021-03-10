# click

## 简介

Click 是一个利用很少的代码以可组合的方式创造优雅命令行工具接口的 Python 库。 它是高度可配置的，但却有合理默认值的“命令行接口创建工具”。
它致力于将创建命令行工具的过程变的快速而有趣，免除你因无法实现一个 CLI API 的挫败感。

- Click 的三个特性
  - 任意嵌套命令
  - 自动生成帮助页面
  - 支持在运行时延迟加载子命令

## 举例

```python
import click

@click.command()
@click.option('--count', default=1, help='Number of greetings.')
@click.option('--name', prompt='Your name', help='The person to greet.')
def hello(count, name):
    """Simple program that greets NAME for a total of COUNT times."""
    for x in range(count):
        click.echo('Hello %s!' % name)

if __name__ == '__main__':
    hello()
```

## 限制类型

- str / click.STRING:表示unicode字符串的默认参数类型。
- int / click.INT:只接受整数的参数。
- float / click.FLOAT:只接受浮点值的参数。
- bool / click.BOOL:接受布尔值的参数。这是自动使用布尔值的标志。如果字符值是: 1, yes, y 和 true 转化为 True ； 0, no, n ， false 转化为 False 。

```python
#"-f", "--foo-bar", the name is foo_bar
#"-x", the name is x
#"-f", "--filename", "dest", the name is dest
#"--CamelCase", the name is camelcase
#"-f", "-fb", the name is f
#"--f", "--foo-bar", the name is f
#"---f", the name is _f
```

## 参考资料

1. [中文文档](https://click-docs-zh-cn.readthedocs.io/zh/latest/)
2. [官网](https://click.palletsprojects.com/en/7.x/options/)
