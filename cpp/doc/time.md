# 1. time

- [1. time](#1-time)
  - [1.1. 简介](#11-简介)
  - [time](#time)
  - [1.3. chrono](#13-chrono)
    - [1.3.1. duration](#131-duration)
    - [1.3.2. time_point](#132-time_point)
  - [1.4. 参考资料](#14-参考资料)

## 1.1. 简介

C++ 标准库没有提供所谓的日期类型。C++ 继承了 C 语言用于日期和时间操作的结构和函数。为了使用日期和时间相关的函数和结构，需要在 C++ 程序中引用 `<ctime>` 头文件。

## time

```C++
typedef /* unspecified */ time_t;//未定义，可以简单理解为 long
typedef long clock_t;//ticker数，CLOCKS_PER_SEC表示每秒ticker数

struct tm {
  int tm_sec;   // 秒，正常范围从 0 到 59，但允许至 61
  int tm_min;   // 分，范围从 0 到 59
  int tm_hour;  // 小时，范围从 0 到 23
  int tm_mday;  // 一月中的第几天，范围从 1 到 31
  int tm_mon;   // 月，范围从 0 到 11
  int tm_year;  // 自 1900 年起的年数
  int tm_wday;  // 一周中的第几天，范围从 0 到 6，从星期日算起
  int tm_yday;  // 一年中的第几天，范围从 0 到 365，从 1 月 1 日算起
  int tm_isdst; // 夏令时
};
```

```C++
time_t time(time_t *time);//返回系统的当前日历时间，自 1970 年 1 月 1 日以来经过的秒数。如果系统没有时间，则返回 -1。

struct tm *localtime(const time_t *time);//该函数返回一个指向表示本地时间的 tm 结构的指针。

char *ctime(const time_t *time);//该返回一个表示当地时间的字符串指针，字符串形式 day month year hours:minutes:seconds year\n\0。
char * asctime ( const struct tm * time );//该函数返回一个指向字符串的指针，字符串包含了 time 所指向结构中存储的信息，返回形式为：day month date hours:minutes:seconds year\n\0。

struct tm *gmtime(const time_t *time);//该函数返回一个指向 time 的指针，time 为 tm 结构，用协调世界时（UTC）也被称为格林尼治标准时间（GMT）表示。
time_t mktime(struct tm *time);//该函数返回日历时间，相当于 time 所指向结构中存储的时间。

double difftime ( time_t time2, time_t time1 );//time1 - time2

//https://www.cplusplus.com/reference/ctime/strftime/
size_t strftime();//该函数可用于格式化日期和时间为指定的格式。

//CLOCKS_PER_SEC
clock_t clock(void);//该函数返回程序执行起（一般为程序的开头），处理器时钟所使用的时间。如果时间不可用，则返回 -1。
```

```C++
#include <ctime>
#include <iostream>
#include <thread> // std::thread, std::thread::id, std::this_thread::get_id

using namespace std;

int main() {
    time_t now = time(0);                        //本机时钟当前时间，单位秒
    cout << "localtime:" << ctime(&now); //转化为字符串显示

    tm *gmtm = gmtime(&now);                     //转换为utc时间
    cout << "UTLtime:" << asctime(gmtm); //转化为字符串显示

    time_t rawtime;
    struct tm *timeinfo;
    time(&rawtime);
    timeinfo = localtime(&rawtime);
    printf("Current local time and date: %s", asctime(timeinfo));

    char buffer[80];
    strftime(buffer, 80, "Now it's %I:%M%p.", timeinfo);
    puts(buffer);

    clock_t t = clock(); //开机执行的ticker，一般1秒1000个ticker，取决于 CLOCKS_PER_SEC 定义，不依赖系统时间修改
    this_thread::sleep_for(chrono::seconds(2));
    cout << "difftime:" << difftime(clock(), t) << endl; //该函数返回 time1 和 time2 之间相差的秒数。
    t = clock() - t;
    printf("It took me %d clicks (%f seconds).\n", t, ((float)t) / CLOCKS_PER_SEC);
    return 0;
}

// localtime:Thu Feb 04 17:05:23 2021
// UTLtime:Thu Feb 04 09:05:23 2021
// Current local time and date: Thu Feb 04 17:05:23 2021
// Now it's 05:05PM.
// difftime:2009
// It took me 2009 clicks (2.009000 seconds).
```

## 1.3. chrono

```C++
#include <chrono>
using namespace chrono;
```

### 1.3.1. duration

duration 是 chrono 裡面，用来记录时间长度的类别，他基本上是一个 template class，可以自行定义他的意义；chrono 也有提供一些比较常见的时间类别，可以直接拿来使用，下面就是内建的 duration 的型别：

```C++
typedef duration<long long, nano> nanoseconds;
typedef duration<long long, micro> microseconds;
typedef duration<long long, milli> milliseconds;
typedef duration<long long> seconds;
typedef duration<int, ratio<60> > minutes;
typedef duration<int, ratio<3600> > hours;
```

```C++
std::chrono::minutes t1( 10 );
std::chrono::seconds t2( 60 );
std::chrono::seconds t3 = t1 - t2;
std::cout << t3.count() << " second" << std::endl;//540 .count()取的一个duration的值
cout << chrono::duration_cast<chrono::minutes>( t3 ).count() << endl;//9
```

### 1.3.2. time_point

相较于 duration 是用来记录时间的长度的
time_point 是用来记录一个特定时间点的资料类别。
他一样是一个 template class，需要指定要使用的 clock 与时间单位（duration）。

`system_clock` 是直接去抓系统的时间，1970-01-01 00:00:00 UTC逝去的时间单位是1个tick
`steady_clock` 系统开机到现在逝去的纳秒
`high_resolution_clock` typedef steady_clock high_resolution_clock;

```C++
#include <ctime>
#include <iostream>
#include <thread> // std::thread, std::thread::id, std::this_thread::get_id

using namespace std;

int main() {
    std::chrono::minutes t1(10);
    std::chrono::seconds t2(60);
    auto t3 = t1 - t2;                                                  //默认转化为更细粒度，t3是std::chrono::seconds
    std::cout << t3.count() << " second" << std::endl;                  //540 .count()取的一个duration的值
    cout << chrono::duration_cast<chrono::minutes>(t3).count() << endl; //9

    std::chrono::steady_clock::time_point now1 = std::chrono::steady_clock::now();
    this_thread::sleep_for(chrono::milliseconds(200));
    std::chrono::steady_clock::time_point now2 = std::chrono::steady_clock::now();
    std::cout << "cost " << std::chrono::duration_cast<std::chrono::microseconds>(now2 - now1).count() << " us.\n";//cost 211927 us.

    std::chrono::system_clock::time_point sysNow = std::chrono::system_clock::now();
    std::chrono::system_clock::time_point sysNow2 = sysNow + std::chrono::hours(2);
    std::time_t now_c = std::chrono::system_clock::to_time_t(sysNow);     //输出先转为time_t
    std::time_t now_c2 = std::chrono::system_clock::to_time_t(sysNow2); //输出先转为time_t
    std::cout << "system_clock now:" << std::ctime(&now_c) << std::endl;//system_clock now:Thu Feb 04 17:24:29 2021
    std::cout << "system_clock now+2h:" << std::ctime(&now_c2) << std::endl;// system_clock now+2h:Thu Feb 04 19:24:29 2021
    return 0;
}
```

## 1.4. 参考资料

1. [C++11 时间工具chrono](https://www.jianshu.com/p/170164adae0f)
