# 1. bitset

- [1. bitset](#1-bitset)
  - [1.1. 简介](#11-简介)
  - [1.2. 函数](#12-函数)
  - [1.3. 举例](#13-举例)
  - [1.4. 参考资料](#14-参考资料)

## 1.1. 简介

1. bitset存储二进制数位。
2. bitset就像一个bool类型的数组一样，但是有空间优化——bitset中的一个元素一般只占1 bit，相当于一个char元素所占空间的八分之一。
3. bitset中的每个元素都能单独被访问，例如对于一个叫做foo的bitset，表达式foo[3]访问了它的第4个元素，就像数组一样。
4. bitset有一个特性：整数类型和布尔数组都能转化成bitset。
5. bitset的大小在编译时就需要确定。如果你想要不确定长度的bitset，请使用（奇葩的）`vector<bool>`。

## 1.2. 函数

- foo.size() 返回大小（位数）
- foo.count() 返回1的个数
- foo.any() 返回是否有1
- foo.none() 返回是否没有1
- foo.set() 全都变成1
- foo.set(p) 将第p + 1位变成1
- foo.set(p, x) 将第p + 1位变成x
- foo.reset() 全都变成0
- foo.reset(p) 将第p + 1位变成0
- foo.flip() 全都取反
- foo.flip(p) 将第p + 1位取反
- foo.to_ulong() 返回它转换为unsigned long的结果，如果超出范围则报错
- foo.to_ullong() 返回它转换为unsigned long long的结果，如果超出范围则报错
- foo.to_string() 返回它转换为string的结果

## 1.3. 举例

```C++
// constructing bitsets
#include <iostream>       // std::cout
#include <string>         // std::string
#include <bitset>         // std::bitset

int main ()
{
  std::bitset<16> foo;
  std::bitset<16> bar (0xfa2);
  std::bitset<16> baz (std::string("0101111001"));

  std::cout << "foo: " << foo << '\n'; //foo: 0000000000000000
  std::cout << "bar: " << bar << '\n'; //bar: 0000111110100010
  std::cout << "baz: " << baz << '\n'; //baz: 0000000101111001

  return 0;
}
```

bitset的运算就像一个普通的整数一样，可以进行与(&)、或(|)、异或(^)、左移(<<)、右移(>>)等操作。

```C++
// bitset operators
#include <iostream>       // std::cout
#include <string>         // std::string
#include <bitset>         // std::bitset

int main ()
{
  std::bitset<4> foo (std::string("1001"));
  std::bitset<4> bar (std::string("0011"));

  std::cout << (foo^=bar) << '\n';       // 1010 (XOR,assign)
  std::cout << (foo&=bar) << '\n';       // 0010 (AND,assign)
  std::cout << (foo|=bar) << '\n';       // 0011 (OR,assign)

  std::cout << (foo<<=2) << '\n';        // 1100 (SHL,assign)
  std::cout << (foo>>=1) << '\n';        // 0110 (SHR,assign)

  std::cout << (~bar) << '\n';           // 1100 (NOT)
  std::cout << (bar<<1) << '\n';         // 0110 (SHL)
  std::cout << (bar>>1) << '\n';         // 0001 (SHR)

  std::cout << (foo==bar) << '\n';       // false (0110==0011)
  std::cout << (foo!=bar) << '\n';       // true  (0110!=0011)

  std::cout << (foo&bar) << '\n';        // 0010
  std::cout << (foo|bar) << '\n';        // 0111
  std::cout << (foo^bar) << '\n';        // 0101

  return 0;
}
```

## 1.4. 参考资料

1. [std::bitset](http://www.cplusplus.com/reference/bitset/bitset/)
2. [C++ bitset](https://www.cnblogs.com/RabbitHu/p/bitset.html)