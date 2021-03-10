# python的包管理

## 概念介绍

test
├── sub1            //package sub1
    ├── test1.py        //module sub1.test1
    ├── test3.py        //module sub1.test3
├── sub2            //package sub2
    ├── test2.py        //module sub2.test2
├── main.py         //package main

- 程序目录(exepath)，运行时，程序所在目录
- 运行时目录(cwdpath)，执行python命令的目录。运行时目录是可以代码修改的
- `sys.path` 存放模块目录，从这里查找模块，可以手动 `sys.path.append("..")` 添加模块目录
  - 被运行文件的目录会被自动加入 sys.path
    - 在sub1运行 `python test1.py`,则 `/user/test/sub1`/`空格` 会被加入sys.path
    - 在test运行 `python sub1/test1.py`,则 `/user/test/sub1`/`sub1` 会被加入sys.path
  - python的安装目录会被加入 sys.path
- package:目录叫package
- module:py文件叫module
- 直接引用:不包含`./../...`的引用(不推荐)
  - `./../...`都是相对运行时目录而言的
  - `.`比较特殊，有时候package名称，有时候是`__main__`,假设test1.py中有代码`from .test3 import hello3`
    - 在sub1运行 `python test1.py`,则 test1.py文件中的 `.`==`__main__`
    - 在test运行 `python main.py`,则 test1.py文件中的 `.`=`sub1` 会被加入sys.path
- 间接引用:包含`./../...`的引用(不推荐)
- 绝对引用:定义top根目录，引用从top开始(不推荐)
  - 比如test1.py引用test2，和test3如下
    - `from sub1.test3 import hello3`
    - `from sub2.test2 import hello2`
  - 运行时，在top目录，通过 `python -m sub1.test1` 运行

## 参考资料

1. [关于package及module管理](https://note.youdao.com/ynoteshare1/index.html?id=a0a38d0dc7ecb4e609ebf6f7f952df77&type=note)