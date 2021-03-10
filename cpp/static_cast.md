# C++的四种强制类型转换

## static_cast

用于良性转换，一般不会导致意外发生，风险很低。

## dynamic_cast

借助 RTTI，用于类型安全的向下转型（Downcasting）。

## const_cast

用于 const 与非 const、volatile 与非 volatile 之间的转换。

## reinterpret_cast

高度危险的转换，这种转换仅仅是对二进制位的重新解释，不会借助已有的转换规则对数据进行调整，但是可以实现最灵活的 C++ 类型转换。