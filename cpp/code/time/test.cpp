#include <iostream>
#include <chrono>

std::time_t getTimeStamp()
{
    std::chrono::time_point<std::chrono::system_clock,std::chrono::milliseconds> tp = 
        std::chrono::time_point_cast<std::chrono::milliseconds>(std::chrono::system_clock::now());
    auto tmp=std::chrono::duration_cast<std::chrono::milliseconds>(tp.time_since_epoch());
    std::time_t timestamp = tmp.count();
    //std::time_t timestamp = std::chrono::system_clock::to_time_t(tp);
    return timestamp;
}

std::tm* gettm(std::time_t timestamp) {
    std::time_t milli = timestamp/*+ (std::time_t)8*60*60*1000*/;//此处转化为东八区北京时间，如果是其它时区需要按需求修改
    auto mTime = std::chrono::milliseconds(milli);
    auto tp=std::chrono::time_point<std::chrono::system_clock,std::chrono::milliseconds>(mTime);
    auto tt = std::chrono::system_clock::to_time_t(tp);
    std::tm* now = std::gmtime(&tt);
    printf("%4d%02d%02d_%02d%02d%02d.%d\n",now->tm_year+1900,now->tm_mon+1,now->tm_mday,now->tm_hour,now->tm_min,now->tm_sec, milli%1000);
   return now;
}

int main(){
    std::chrono::minutes t1( 10 );
    std::chrono::seconds t2( 60 );
    std::chrono::seconds t3 = t1 - t2;
    std::cout << t3.count() << " second" << std::endl;//540 .count()取的一个duration的值
    std::cout << std::chrono::duration_cast<std::chrono::minutes>( t3 ).count() << std::endl;//9
    {
        std::chrono::steady_clock::time_point t1 = std::chrono::steady_clock::now();
        std::cout << "Hello World\n";
        std::chrono::steady_clock::time_point t2 = std::chrono::steady_clock::now();
        std::cout << "Printing took " 
        << std::chrono::duration_cast<std::chrono::microseconds>(t2 - t1).count() 
        << "us." <<std::endl;
        std::chrono::steady_clock::time_point nt = t1 + std::chrono::hours(10);
    }
    {
        std::chrono::system_clock::time_point t1 = std::chrono::system_clock::now();
        std::time_t now_c = std::chrono::system_clock::to_time_t(t1);//输出先转为time_t
        std::cout << std::ctime( &now_c ) << std::endl; 
        gettm(getTimeStamp()); 
    }
    
    return 0;
}
