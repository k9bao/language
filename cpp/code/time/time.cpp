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
    std::cout << "cost " << std::chrono::duration_cast<std::chrono::microseconds>(now2 - now1).count() << " us.\n";

    std::chrono::system_clock::time_point sysNow = std::chrono::system_clock::now();
    std::chrono::system_clock::time_point sysNow2 = sysNow + std::chrono::hours(2);
    std::time_t now_c = std::chrono::system_clock::to_time_t(sysNow);     //输出先转为time_t
    std::time_t now_c2 = std::chrono::system_clock::to_time_t(sysNow2); //输出先转为time_t
    std::cout << "system_clock now:" << std::ctime(&now_c) << std::endl;
    std::cout << "system_clock now+2h:" << std::ctime(&now_c2) << std::endl;

    time_t now = time(0);                //本机时钟当前时间，单位秒
    cout << "localtime:" << ctime(&now); //转化为字符串显示

    tm *gmtm = gmtime(&now);             //转换为utc时间
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