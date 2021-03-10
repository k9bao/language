#include <chrono>   // std::chrono::seconds
#include <future>   // std::packaged_task, std::future
#include <iostream> // std::cout
#include <thread>   // std::thread, std::this_thread::sleep_for
#include <utility>  // std::move

int task_move() {
    std::cout << "--------------- task_move" << std::endl;
    std::packaged_task<int(int)> foo; // 默认构造函数.

    // 使用 lambda 表达式初始化一个 packaged_task 对象.
    std::packaged_task<int(int)> bar([](int x) { return x * 2; });

    foo = std::move(bar); // move-赋值操作，也是 C++11 中的新特性.

    // 获取与 packaged_task 共享状态相关联的 future 对象.
    std::future<int> ret = foo.get_future();

    std::thread(std::move(foo), 10).detach(); // 产生线程，调用被包装的任务.

    int value = ret.get(); // 等待任务完成并获取结果.
    std::cout << "The double of 10 is " << value << ".\n";

    std::cout << "--------------- task_move" << std::endl;
    return 0;
}

int packaged_task() {
    std::cout << "--------------- packaged_task" << std::endl;
    // count down taking a second for each value:
    auto countdown = [](int from, int to) {
        for (int i = from; i != to; --i) {
            std::cout << i << '\n';
            std::this_thread::sleep_for(std::chrono::seconds(1));
        }
        std::cout << "Finished!\n";
        return from - to;
    };

    std::packaged_task<int(int, int)> task(countdown); // 设置 packaged_task
    std::future<int> ret = task.get_future();          // 获得与 packaged_task 共享状态相关联的 future 对象.

    std::thread th(std::move(task), 3, 0); //创建一个新线程完成计数任务.

    int value = ret.get(); // 等待任务完成并获取结果.

    std::cout << "The countdown lasted for " << value << " seconds.\n";

    th.join();
    std::cout << "--------------- packaged_task" << std::endl;
    return 0;
}

int valid() {
    std::cout << "--------------- valid" << std::endl;

    // 在新线程中启动一个 int(int) packaged_task.
    auto launcher = [](std::packaged_task<int(int)> &tsk, int arg) {
        if (tsk.valid()) {
            std::future<int> ret = tsk.get_future();
            std::thread(std::move(tsk), arg).detach();
            return ret;
        } else
            return std::future<int>();
    };

    std::packaged_task<int(int)> tsk0([](int x) { return x * 2; });
    try {
        std::future<int> fut = launcher(tsk0, 25);
        auto result = fut.get();
        std::cout << "The double of 25 is " << result << ".\n";
    } catch (const std::future_error &e) {
        std::cout << "Caught a future_error with code \"" << e.code()
                  << "\"\nMessage: \"" << e.what() << "\"\n";
    }

    std::packaged_task<int(int)> tsk1; // package task

    try {
        std::future<int> fut = launcher(tsk1, 25);
        auto result = fut.get();
        std::cout << "The double of 25 is " << result << ".\n";
    } catch (const std::future_error &e) {
        std::cout << "Caught a future_error with code \"" << e.code()
                  << "\"\nMessage: \"" << e.what() << "\"\n";
    }
    std::cout << "--------------- valid" << std::endl;

    return 0;
}

int reset() {
    std::cout << "--------------- reset" << std::endl;

    std::packaged_task<int(int)> tsk([](int x) { return x * 3; }); // package task
    std::future<int> fut = tsk.get_future();
    std::thread(std::ref(tsk), 100).detach();
    std::cout << "The triple of 100 is " << fut.get() << ".\n";

    // re-use same task object:
    tsk.reset();
    fut = tsk.get_future();
    std::thread(std::move(tsk), 200).detach();
    std::cout << "Thre triple of 200 is " << fut.get() << ".\n";
    std::cout << "--------------- reset" << std::endl;

    return 0;
}

int pkg_operator() {
    std::cout << "--------------- pkg_operator" << std::endl;
    auto task = [](int i) {
        std::this_thread::sleep_for(std::chrono::seconds(2));
        return i + 100;
    };

    std::packaged_task<int(int)> package{task};
    std::future<int> f = package.get_future();
    package(1);
    std::cout << f.get() << "\n";
    std::cout << "--------------- pkg_operator" << std::endl;
    return 0;
}

int main() {
    packaged_task();
    task_move();
    valid();
    reset();
    pkg_operator();
    return 0;
}