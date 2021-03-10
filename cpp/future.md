# 1. future

- [1. future](#1-future)
  - [1.1. future相关结构](#11-future相关结构)
  - [1.2. async相关结构](#12-async相关结构)
  - [1.3. packaged_task 相关结构](#13-packaged_task-相关结构)
  - [1.4. promise相关结构](#14-promise相关结构)
  - [1.5. 参考资料](#15-参考资料)

## 1.1. future相关结构

Member functions

| fun           | desc                                                                    | pub/priv                 |
| ------------- | ----------------------------------------------------------------------- | ------------------------ |
| (constructor) | constructs the future object                                            | (public member function) |
| (destructor)  | destructs the future object                                             | (public member function) |
| operator=     | moves the future object                                                 | (public member function) |
| share         | transfers the shared state from *this to a shared_future and returns it | (public member function) |

Getting the result

| fun        | desc                                                                                             | pub/priv                 |
| ---------- | ------------------------------------------------------------------------------------------------ | ------------------------ |
| get        | returns the result                                                                               | (public member function) |
| valid      | checks if the future has a shared state                                                          | (public member function) |
| wait       | waits for the result to become available                                                         | (public member function) |
| wait_for   | waits for the result, returns if it is not available for the specified timeout duration          | (public member function) |
| wait_until | waits for the result, returns if it is not available until specified time point has been reached | (public member function) |

## 1.2. async相关结构

简单快捷，简易版std::thread
std::async 并不总会开启新的线程来执行任务，你可以指定 std::launch::async 来强制开启新线程
如果 std::async 返回值 std::future 被存放在一个临时对象中，那么std::async会立马阻塞，因为临时对象在返回后立马被析构了。
`std::async( std::launch::async, sleep, 5 ); // 临时对象被析构，阻塞 5s`
`auto f1 = std::async( std::launch::async, sleep, 5 );`如果不調用f1.get(),则std::async可能都不执行

```C++
async( Function&& f, Args&&... args );
//policy: std::launch::async,std::launch::deferred,立刻异步执行或调用时再触发异步执行
async( std::launch policy, Function&& f, Args&&... args );
```

```C++
int asyncTime() {
    auto sleep = [](int s) { std::this_thread::sleep_for(std::chrono::seconds(s)); };

    auto start = now();
    std::async(std::launch::async, sleep, 2); // 临时对象被析构，阻塞 2s
    std::cout << "cost:" << sub(start, now()) << std::endl;//cost 2s
    start = now();
    std::async(std::launch::async, sleep, 2); // 临时对象被析构，阻塞 2s
    std::cout << "cost:" << sub(start, now()) << std::endl;//cost 2s
    
    start = now();
    auto f1 = std::async(std::launch::async, sleep, 2);
    std::cout << "cost:" << sub(start, now()) << std::endl;//cost 0s
    start = now();
    auto f2 = std::async(std::launch::async, sleep, 2);
    std::cout << "cost:" << sub(start, now()) << std::endl;//cost 0s
    f2.get();
    std::cout << "cost:" << sub(start, now()) << std::endl;//cost 2s
    return 0;
}
```

## 1.3. packaged_task 相关结构

配合 std::thread，可以取代 async
实际上就是std::futrue的封装函数，没啥特殊函数，参数是指向一个函数
可以通过 packaged_task对象直接调用关联函数(支持`operator()`)。

```C++
    auto task = [](int i) { 
    std::this_thread::sleep_for(std::chrono::seconds(5)); return i+100; 
    };

    std::packaged_task< int(int) > package{ task };
    std::future<int> f = package.get_future();
    package(1);//直接调用
    std::cout << f.get() << "\n";

    std::packaged_task< int(int) > package2{ task };
    std::future<int> f = package.get_future();
    std::thread t { std::move(package2), 5 };//线程执行
    std::cout << f.get() << std::endl; // 阻塞，直到线程 t 结束
    t.join();
```

[packaged_task](https://en.cppreference.com/w/cpp/thread/packaged_task)

|Member functions

| fun                                     | desc                                                                                       | pub/priv                        |
| --------------------------------------- | ------------------------------------------------------------------------------------------ | ------------------------------- |
| (constructor)                           | constructs the task object                                                                 | (public member function)        |
| (destructor)                            | destructs the task object                                                                  | (public member function)        |
| operator=                               | moves the task object                                                                      | (public member function)        |
| valid                                   | checks if the task object has a valid function                                             | (public member function)        |
| swap                                    | swaps two task objects                                                                     | (public member function)        |
| get_future                              | returns a std::future associated with the promised result                                  | (public member function)        |
| operator()                              | executes the function                                                                      | (public member function)        |
| make_ready_at_thread_exit               | executes the function ensuring that the result is ready only once the current thread exits | (public member function)        |
| reset                                   | resets the state abandoning any stored results of previous executions                      | (public member function)        |
| std::swap(std::packaged_task)(C++11)    | specializes the std::swap algorithm                                                        | (function template)             |
| std::uses_allocator<std::packaged_task> | specializes the std::uses_allocator type trait                                             | (class template specialization) |

## 1.4. promise相关结构

更像一个线程间同步的封装类，一个调用set_value设置变量，一个调用get_future().get()获取变量

```C++
    auto task = [](std::future<int> i) {
        std::cout << i.get() << std::flush; // 阻塞，直到 p.set_value() 被调用
    };

    std::promise<int> p;
    std::thread t{ task, p.get_future() };

    std::this_thread::sleep_for(std::chrono::seconds(5));
    p.set_value(5);

    t.join();
```

[promise](https://en.cppreference.com/w/cpp/thread/promise)

Member functions

| fun           | desc                          | pub/priv                 |
| ------------- | ----------------------------- | ------------------------ |
| (constructor) | constructs the promise object | (public member function) |
| (destructor)  | destructs the promise object  | (public member function) |
| operator=     | assigns the shared state      | (public member function) |
| swap          | swaps two promise objects     | (public member function) |

Getting the result

| fun                               | desc                                                                                           | pub/priv                        |
| --------------------------------- | ---------------------------------------------------------------------------------------------- | ------------------------------- |
| get_future                        | returns a future associated with the promised result                                           | (public member function)        |
| set_value                         | sets the result to specific value                                                              | (public member function)        |
| set_value_at_thread_exit          | sets the result to specific value while delivering the notification only at thread exit        | (public member function)        |
| set_exception                     | sets the result to indicate an exception                                                       | (public member function)        |
| set_exception_at_thread_exit      | sets the result to indicate an exception while delivering the notification only at thread exit | (public member function)        |
| std::swap(std::promise)           | specializes the std::swap algorithm                                                            | (function template)             |
| std::uses_allocator<std::promise> | specializes the std::uses_allocator type trait                                                 | (class template specialization) |

## 1.5. 参考资料

1. [C++11并发指南系列](https://www.cnblogs.com/haippy/p/3284540.html)
2. [cppreference](https://en.cppreference.com/w/cpp/thread/future)
