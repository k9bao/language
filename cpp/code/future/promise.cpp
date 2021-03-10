#include <chrono>     // std::chrono::seconds
#include <exception>  // std::exception, std::current_exception
#include <functional> // std::ref
#include <future>     // std::promise, std::future
#include <iostream>   // std::cout
#include <thread>     // std::thread

void promise() {
    std::cout << "--------------- promise" << std::endl;
    auto print_int = [](std::future<int> &fut) {
        int x = fut.get();
        std::cout << "value: " << x << '\n';
    };
    std::promise<int> prom;                   // 生成一个 std::promise<int> 对象.
    std::future<int> fut = prom.get_future(); // 和 future 关联.
    std::thread t(print_int, std::ref(fut));  // 将 future 交给另外一个线程t.
    prom.set_value(10);                       // 设置共享状态的值, 此处和线程t保持同步.
    t.join();
    std::cout << "--------------- promise" << std::endl;
}

int promise_move() {
    std::cout << "--------------- promise_move" << std::endl;
    auto print_int_promise = [](std::promise<int> &prom) {
        try {
            std::cout << "before get: " << '\n';
            std::future<int> fut = prom.get_future();
            int x = fut.get();
            std::cout << "value: " << x << '\n';
        } catch (std::exception &e) {
            std::cout << "[exception caught: " << e.what() << "]\n";
        }
        std::cout << "print_promise over" << std::endl;
    };

    std::promise<int> prom;
    std::thread th1(print_int_promise, std::ref(prom));
    prom.set_value(10);
    th1.join();

    prom = std::promise<int>(); // prom 被move赋值为一个新的 promise 对象.

    std::thread th2(print_int_promise, std::ref(prom));
    prom.set_value(20);
    th2.join();
    std::cout << "--------------- promise_move" << std::endl;

    return 0;
}

int set_value_at_thread_exit() {
    std::cout << "--------------- set_value_at_thread_exit" << std::endl;
    auto print_int = [](std::future<int> &fut) {
        int x = fut.get();
        std::cout << "value: " << x << '\n';
    };

    auto set_thread_exit = [](std::promise<int> &prom) {
        prom.set_value_at_thread_exit(100);
        std::cout << "set_thread_exit over" << std::endl;
    };

    std::promise<int> prom;
    std::future<int> fut = prom.get_future();

    std::thread th1(set_thread_exit, std::ref(prom));
    std::thread th2(print_int, std::ref(fut));

    th1.join();
    th2.join();
    std::cout << "--------------- set_value_at_thread_exit" << std::endl;
    return 0;
}

int get_exceptions() {
    std::cout << "--------------- get_exceptions" << std::endl;
    auto get_int = [](std::promise<int> &prom) {
        int x;
        std::cout << "Please, enter an integer value: ";
        std::cin.exceptions(std::ios::failbit); // throw on failbit
        try {
            std::cin >> x; // sets failbit if input is not int
            prom.set_value(x);
        } catch (std::exception &) {
            prom.set_exception(std::current_exception());
        }
    };
    auto print_int = [](std::future<int> &fut) {
        try {
            std::cout << "before get: " << '\n';
            int x = fut.get();
            std::cout << "value: " << x << '\n';
        } catch (std::exception &e) {
            std::cout << "[exception caught: " << e.what() << "]\n";
        }
        std::cout << "print_int over" << std::endl;
    };
    std::promise<int> prom;
    std::future<int> fut = prom.get_future();

    std::thread th1(get_int, std::ref(prom));
    std::thread th2(print_int, std::ref(fut));

    th1.join();
    th2.join();
    std::cout << "--------------- get_exceptions" << std::endl;
    return 0;
}

int main() {
    promise();
    promise_move();
    set_value_at_thread_exit();
    get_exceptions();
    return 0;
}