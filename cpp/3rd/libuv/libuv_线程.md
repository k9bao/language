
索引目录：https://blog.csdn.net/knowledgebao/article/details/84776754

Threads are used internally to fake the asynchronous nature of all of the system calls. libuv also uses threads to allow you, the application, to perform a task asynchronously that is actually blocking, by spawning a thread and collecting the result when it is done.

Today there are two predominant thread libraries: the Windows threads implementation and POSIX’s pthreads. libuv’s thread API is analogous to the pthreads API and often has similar semantics.（目前主要有二类线程，一种是window的POSIX，一种是linux的pthread，libuv封装的thread与pthread更相近）

A notable aspect of libuv’s thread facilities is that it is a self contained section within libuv. Whereas other features intimately depend on the event loop and callback principles, threads are complete agnostic, they block as required, signal errors directly via return values, and, as shown in the first example, don’t even require a running event loop.（线程在libuv中有被使用，很多情况和loop关联，当然线程也可以单独使用）

libuv’s thread API is also very limited since the semantics and syntax of threads are different on all platforms, with different levels of completeness.（因为不同线程的实现，封装的功能仅包含部分线程功能）

This chapter makes the following assumption: There is only one event loop, running in one thread (the main thread). No other thread interacts with the event loop (except using uv_async_send).（libuv中使用loop的话，只允许使用一个loop）

typedef void (*uv_thread_cb)(void* arg);

int uv_thread_create(uv_thread_t* tid, uv_thread_cb entry, void* arg);
uv_thread_t uv_thread_self(void);
int uv_thread_join(uv_thread_t *tid);
int uv_thread_equal(const uv_thread_t* t1, const uv_thread_t* t2);
int uv_key_create(uv_key_t* key);
void uv_key_delete(uv_key_t* key);
void* uv_key_get(uv_key_t* key);
void uv_key_set(uv_key_t* key, void* value);
uv_thread_t：Thread data type

void (*uv_thread_cb)(void* arg)：Callback that is invoked to initialize thread execution. arg is the same value that was passed to uv_thread_create().

uv_thread_create：第一个参数是线程ID，为输出参数。

uv_thread_t uv_thread_self：返回本线程ID，在线程内部调用。

uv_thread_join：等待线程结束。

uv_thread_equal：判断二个线程是否是同一个。如果相同，返回非0，不相同返回0

uv_key_t：Thread-local key data type.

uv_key_create:创建局部变量,第一个参数是返回值key

uv_key_delete:删除局部变量

uv_key_get:获取关联指针

uv_key_set:设置关联指针

int uv_mutex_init(uv_mutex_t* handle);
int uv_mutex_init_recursive(uv_mutex_t* handle);
void uv_mutex_destroy(uv_mutex_t* handle);
void uv_mutex_lock(uv_mutex_t* handle);
int uv_mutex_trylock(uv_mutex_t* handle);
void uv_mutex_unlock(uv_mutex_t* handle);
int uv_rwlock_init(uv_rwlock_t* rwlock);
void uv_rwlock_destroy(uv_rwlock_t* rwlock);
void uv_rwlock_rdlock(uv_rwlock_t* rwlock);
int uv_rwlock_tryrdlock(uv_rwlock_t* rwlock);
void uv_rwlock_rdunlock(uv_rwlock_t* rwlock);
void uv_rwlock_wrlock(uv_rwlock_t* rwlock);
int uv_rwlock_trywrlock(uv_rwlock_t* rwlock);
void uv_rwlock_wrunlock(uv_rwlock_t* rwlock);
互斥锁

1，uv_mutex_init：初始化锁，handle是返回值

2，uv_mutex_init_recursive：初始化线程循环锁

3，uv_mutex_destroy：销毁锁

4，uv_mutex_lock：上锁

5，uv_mutex_trylock:尝试上锁

6，uv_mutex_unlock：解锁

读写锁加锁操作

假设当前临界区没有任何进程，这时候read进程或者write进程都可以进来，但是只能是其一

如果当前临界区只有一个read进程，这时候任意的read进程都可以进入，但是write进程不能进入

如果当前临界区只有一个write进程，这时候任何read/write进程都无法进入。只能自旋等待

如果当前当前临界区有好多个read进程，同时read进程依然还会进入，这时候进入的write进程只能等待。直到临界区一个read进程都没有，才可进入

读写锁解锁操作

如果在read进程离开临界区的时候，需要根据情况决定write进程是否需要进入。只有当临界区没有read进程了，write进程方可进入。

如果在write进程离开临界区的时候，无论write进程或者read进程都可进入临界区，因为write进程是排它的。

1，uv_rwlock_init:初始化读写锁

2，uv_rwlock_destroy：销毁读写锁

3，uv_rwlock_rdlock：读上锁

4，uv_rwlock_tryrdlock：读尝试上锁

5，rwlock_rdunlock：读解锁

6，uv_rwlock_wrlock：写上锁

7，uv_rwlock_trywrlock：写尝试上锁

8，uv_rwlock_wrunlock：写解锁

int uv_sem_init(uv_sem_t* sem, unsigned int value);
void uv_sem_destroy(uv_sem_t* sem);
void uv_sem_post(uv_sem_t* sem);
void uv_sem_wait(uv_sem_t* sem);
int uv_sem_trywait(uv_sem_t* sem);
信号量允许多个线程同时进入临界区。

https://blog.csdn.net/qyz_og/article/details/47189219

1，uv_sem_init：初始化信号量，value：信号量值

2，uv_sem_destroy：销户

3，uv_sem_post：释放使value++

4，uv_sem_wait：等待获取，如果value>0,直接返回，并且value--。否则等待，知道value>0

5，uv_sem_trywait:尝试获取,if value>0,return 0;else return -1;

int uv_barrier_init(uv_barrier_t* barrier, unsigned int count);
void uv_barrier_destroy(uv_barrier_t* barrier);
int uv_barrier_wait(uv_barrier_t* barrier);
https://blog.csdn.net/qq405180763/article/details/23919191

把先后到达的多个线程挡在同一栏杆前，直到所有线程到齐，然后撤下栏杆同时放行。

1，uv_barrier_init：初始化屏障，count等待线程个数

2，uv_barrier_destroy：销毁

3，uv_barrier_wait：等待，线程内部调用。

int uv_cond_init(uv_cond_t* cond);
void uv_cond_destroy(uv_cond_t* cond);
void uv_cond_signal(uv_cond_t* cond);
void uv_cond_broadcast(uv_cond_t* cond);
void uv_cond_wait(uv_cond_t* cond, uv_mutex_t* mutex);
int uv_cond_timedwait(uv_cond_t* cond, uv_mutex_t* mutex,uint64_t timeout);
https://www.cnblogs.com/jiu0821/p/6424951.html

1，uv_cond_init：初始化条件变量，与pthread_cond_init的区别是缺少属性设置attr。

2，uv_cond_destroy：销毁条件变量

3，uv_cond_signal：释放被阻塞在指定条件变量上的一个线程。

4，uv_cond_broadcast：释放阻塞的所有线程

5，uv_cond_wait：函数将解锁mutex参数指向的互斥锁，并使当前线程阻塞在cond参数指向的条件变量上。

被阻塞的线程可以被uv_cond_signal函数，uv_cond_broadcast函数唤醒，也可能在被信号中断后被唤醒。

uv_cond_wait函数的返回并不意味着条件的值一定发生了变化，必须重新检查条件的值。

uv_cond_wait函数返回时，相应的互斥锁将被当前线程锁定，即使是函数出错返回。

一般一个条件表达式都是在一个互斥锁的保护下被检查。当条件表达式未被满足时，线程将仍然阻塞在这个条件变量上。当另一个线程改变了条件的值并向条件变量发出信号时，等待在这个条件变量上的一个线程或所有线程被唤醒，接着都试图再次占有相应的互斥锁。

阻塞在条件变量上的线程被唤醒以后，直到uv_cond_wait()函数返回之前条件的值都有可能发生变化。所以函数返回以后，在锁定相应的互斥锁之前，必须重新测试条件值。最好的测试方法是循环调用uv_cond_wait函数，并把满足条件的表达式置为循环的终止条件。如：

uv_mutex_lock();

while (condition_is_false)

        uv_cond_wait();

uv_mutex_unlock();

阻塞在同一个条件变量上的不同线程被释放的次序是不一定的。

注意：uv_cond_wait()函数是退出点，如果在调用这个函数时，已有一个挂起的退出请求，且线程允许退出，这个线程将被终止并开始执行善后处理函数，而这时和条件变量相关的互斥锁仍将处在锁定状态。

6，uv_cond_timedwait：在uv_cond_wait的基础上增加超时功能。

7.唤醒丢失问题

在线程未获得相应的互斥锁时调用pthread_cond_signal或pthread_cond_broadcast函数可能会引起唤醒丢失问题。

唤醒丢失往往会在下面的情况下发生：

一个线程调用pthread_cond_signal或pthread_cond_broadcast函数；

另一个线程正处在测试条件变量和调用pthread_cond_wait函数之间；

没有线程正在处在阻塞等待的状态下。

void uv_once(uv_once_t* guard, void (*callback)(void));
Multiple threads can attempt to call uv_once() with a given guard and a function pointer, only the first one will win, the function will be called once and only once

/* Initialize guard */

static uv_once_t once_only = UV_ONCE_INIT;

int i = 0;

void increment() { i++;}

void thread1() {

    /* ... work */

    uv_once(once_only, increment);

}

void thread2() {

    /* ... work */

    uv_once(once_only, increment);

}

int main() {

    /* ... spawn threads */

}

After all threads are done, i == 1.



线程池调用
int uv_queue_work(uv_loop_t* loop,
                            uv_work_t* req,
                            uv_work_cb work_cb,
                            uv_after_work_cb after_work_cb);

int uv_cancel(uv_req_t* req);


http://docs.libuv.org/en/v1.x/guide/threads.html

libuv provides a threadpool which can be used to run user code and get notified in the loop thread. This thread pool is internally used to run all file system operations, as well as getaddrinfo and getnameinfo requests.

Its default size is 4, but it can be changed at startup time by setting the UV_THREADPOOL_SIZE environment variable to any value (the absolute maximum is 128).

The threadpool is global and shared across all event loops. When a particular function makes use of the threadpool (i.e. when using uv_queue_work()) libuv preallocates and initializes the maximum number of threads allowed by UV_THREADPOOL_SIZE. This causes a relatively minor memory overhead (~1MB for 128 threads) but increases the performance of threading at runtime.


，可用于运行用户代码并在循环线程中得到通知。此线程池在内部用于运行所有文件系统操作，以及getaddrinfo和getnameinfo请求。

其默认大小为4，但可以通过将UV_THREADPOOL_SIZE环境变量设置为任何值（绝对最大值为128，putenv("UV_THREADPOOL_SIZE=128")）在启动时更改 。

线程池是全局的，并在所有事件循环中共享。当特定函数使用线程池时（即使用时），libuv会预分配并初始化允许的最大线程数。这会导致相对较小的内存开销（128个线程约为1MB），但会增加运行时的线程性能。

uv_queue_work：

Initializes a work request which will run the given work_cb in a thread from the threadpool. Once work_cb is completed, after_work_cb will be called on the loop thread.

This request can be cancelled with uv_cancel().

Note that even though a global thread pool which is shared across all events loops is used, the functions are not thread safe.

从线程池中获取一个线程，执行work_cb任务。执行完成后，通过after_work_cb通知。work_cb不可以为空，否则after_work_cb不会执行。

Loop主要用来调用after_work_cb回调结果，如果想要after_work_cb返回，loop必须调用uv_run.

从线程池获取线程，然后执行对应的任务，无需自己创建线程。

uv_cancel:取请求执行

/*
uv_work_t:Work request type.
void (*uv_work_cb)(uv_work_t* req)
void (*uv_after_work_cb)(uv_work_t* req, int status)
*/
int uv_queue_work(uv_loop_t* loop, uv_work_t* req, uv_work_cb work_cb, uv_after_work_cb after_work_cb)
注意事项：

创建子进程用spawn，不要fork，fork的子进程中调用uv_queue_work，不执行



有任何问题，请联系：knowledgebao@163.com