// future example
#include <chrono> // std::chrono::milliseconds
#include <cmath>
#include <future>   // std::async, std::future
#include <iostream> // std::cout

std::chrono::steady_clock::time_point now() {
    return std::chrono::steady_clock::now();
}

long long sub(std::chrono::steady_clock::time_point begin, std::chrono::steady_clock::time_point end) {
    return std::chrono::duration_cast<std::chrono::seconds>(end - begin).count();
}

int shared_future() {
    std::cout << "--------------- shared_future" << std::endl;
    std::future<int> fut = std::async([] { return 10; });
    std::shared_future<int> shared_fut = fut.share(); //after fut is not valid

    // 共享的 future 对象可以被多次访问.
    std::cout << "value: " << shared_fut.get() << '\n';
    std::cout << "its double: " << shared_fut.get() * 2 << '\n';

    std::cout << "--------------- shared_future" << std::endl;
    return 0;
}

int wait() {
    // a non-optimized way of checking for prime numbers:
    auto is_prime = [](int x) {
        for (int i = 2; i < x; ++i) {
            if (x % i == 0)
                return false;
        }
        return true;
    };
    // call function asynchronously:
    std::future<bool> fut = std::async(is_prime, 194232491);

    std::cout << "Checking...\n";
    fut.wait();

    std::cout << "\n194232491 ";
    if (fut.get()) // guaranteed to be ready (and not block) after wait returns
        std::cout << "is prime.\n";
    else
        std::cout << "is not prime.\n";

    return 0;
}

int wait_for() {
    std::cout << "--------------- wait_for" << std::endl;
    // a non-optimized way of checking for prime numbers:
    auto is_prime = [](int x) {
        for (int i = 2; i < x; ++i) {
            if (x % i == 0)
                return false;
        }
        return true;
    };
    // call function asynchronously:
    std::future<bool> fut = std::async(is_prime, 444444443);

    // do something while waiting for function to set future:
    std::cout << "checking, please wait";
    std::chrono::milliseconds span(100);
    while (fut.wait_for(span) == std::future_status::timeout)
        std::cout << '.';

    bool x = fut.get(); // retrieve return value

    std::cout << "\n444444443 " << (x ? "is" : "is not") << " prime.\n";

    std::cout << "--------------- wait_for" << std::endl;
    return 0;
}

int async() {
    std::cout << "--------------- async" << std::endl;

    auto ThreadTask = [](int n) {
        std::cout << std::this_thread::get_id()
                  << " start computing..." << std::endl;

        double ret = 0;
        for (int i = 0; i <= n; i++) {
            ret += std::sin(i); //正弦函數
        }

        std::cout << std::this_thread::get_id()
                  << " finished computing..." << std::endl;
        return ret;
    };

    std::future<double> f(std::async(std::launch::async, ThreadTask, 100000000));

#if 0
    while(f.wait_until(std::chrono::system_clock::now() + std::chrono::seconds(1))
            != std::future_status::ready) {
        std::cout << "task is running...\n";
    }
#else
    while (f.wait_for(std::chrono::seconds(1)) != std::future_status::ready) {
        std::cout << "task is running...\n";
    }
#endif

    std::cout << f.get() << std::endl;
    std::cout << "--------------- async" << std::endl;

    return EXIT_SUCCESS;
}

int deferred() {
    std::cout << "--------------- deferred" << std::endl;
    auto do_print_ten = [](char c, int ms) {
        for (int i = 0; i < 10; ++i) {
            std::this_thread::sleep_for(std::chrono::milliseconds(ms));
            std::cout << c;
        }
    };

    std::cout
        << "with launch::async:\n";
    std::future<void> foo =
        std::async(std::launch::async, do_print_ten, '*', 100);
    std::future<void> bar =
        std::async(std::launch::async, do_print_ten, '@', 200);
    // async "get" (wait for foo and bar to be ready):
    foo.get();
    bar.get();
    std::cout << "\n\n";

    std::cout << "with launch::deferred:\n";
    foo = std::async(std::launch::deferred, do_print_ten, '*', 100);
    bar = std::async(std::launch::deferred, do_print_ten, '@', 200);
    // deferred "get" (perform the actual calls):
    foo.get();
    bar.get();
    std::cout << '\n';
    std::cout << "--------------- deferred" << std::endl;

    return 0;
}

int asyncTime() {
    auto sleep = [](int s) { std::this_thread::sleep_for(std::chrono::seconds(s)); };

    auto start = now();
    std::async(std::launch::async, sleep, 2);               // 临时对象被析构，阻塞 2s
    std::cout << "cost:" << sub(start, now()) << std::endl; //cost 2s
    start = now();
    std::async(std::launch::async, sleep, 2);               // 临时对象被析构，阻塞 2s
    std::cout << "cost:" << sub(start, now()) << std::endl; //cost 2s

    start = now();
    auto f1 = std::async(std::launch::async, sleep, 2);
    std::cout << "cost:" << sub(start, now()) << std::endl; //cost 0s
    start = now();
    auto f2 = std::async(std::launch::async, sleep, 2);
    std::cout << "cost:" << sub(start, now()) << std::endl; //cost 0s
    f2.get();
    std::cout << "cost:" << sub(start, now()) << std::endl; //cost 2s
    return 0;
}

int main() {
    shared_future();
    wait_for();
    wait();
    async();
    deferred();
    auto start = now();
    asyncTime();
    std::cout << "main::cost:" << sub(start, now()) << std::endl; //cost 6s
    return 0;
}