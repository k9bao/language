# 1. thread

- [1. thread](#1-thread)
  - [1.1. thread](#11-thread)
  - [1.2. this_thread](#12-this_thread)
  - [1.3. 参考资料](#13-参考资料)

```C++
#include <thread>
using namespace std::thread;
using namespace std::this_thread;
```

## 1.1. thread

## 1.2. this_thread

```C++
thread::id get_id() noexcept;//  Get thread id (function )
void yield() noexcept;//   Yield to other threads (function )
template <class Clock, class Duration>
  void sleep_until (const chrono::time_point<Clock,Duration>& abs_time);//Sleep until time point (function )绝对时间
template <class Rep, class Period>
  void sleep_for (const chrono::duration<Rep,Period>& rel_time);  //Sleep for time span (function )
```

```C++
#include <chrono>
#include <thread>
#include <iostream>

int main(){
    auto start = std::chrono::steady_clock::now();//time_point
    std::chrono::milliseconds wait(500);//duration 50 ms
    while(1){
        auto delta = std::chrono::steady_clock::now()-start;
        if(delta > wait){
            std::cout << "time out";
            break;
        }
        std::cout << "cost" << std::chrono::duration_cast<std::chrono::milliseconds>(delta).count() << "ms" << std::endl;
        std::this_thread::sleep_for(std::chrono::milliseconds(10));//sleep 10 ms
    }
}
```

## 1.3. 参考资料

1. [this_thread官网](http://www.cplusplus.com/reference/thread/this_thread/)
2. [thread官网](http://www.cplusplus.com/reference/thread/thread/)
3. [C++11 并发指南系列](https://www.cnblogs.com/haippy/p/3284540.html)
