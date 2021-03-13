# 1. 异步调度

- [1. 异步调度](#1-异步调度)
  - [1.1. loop相关API](#11-loop相关api)
  - [1.2. 参考资料](#12-参考资料)

调度流程图

&emsp;&emsp;&emsp;&emsp;![调度分类](./img/loop.png)

&emsp;&emsp;上图是loop的while循环流程。

1. `Updata loop time` 为了减少系统时间依赖性，循环开始位置更新当前时间指。
2. `loop alive` 判断loop是否存活，存活的条件：1 激活且被引用的handle，2 激活的request，3 正在关闭中的handle
3. `Run due timers` 判断是否有存活的定时器handle
4. `Call pending callbacks` 执行上一轮遗留的 I/O 回调事件
5. `Run idle handles` 执行注册的 idle 回调事件
6. `Run prepare handles` 执行注册的 prepare 回调事件
7. `Poll for I/O` 这里会阻塞 timeout 时长，监听异步 I/O ，如果存在 I/O事件，则调用对应的回调函数,timeout计算方法如下
   - 如果 loop 运行时，使用 UV_RUN_NOWAIT 参数，timeout=0
   - 如果 uv_stop() 被调用，timeout=0
   - 如果没有激活的 handle 或 request，timeout=0
   - 如果存在 idle 的激活事件，timeout=0
   - 如果遗留 handle 被关闭，timeout=0
   - 如果不是上边情况，则 timeout 就是最近的 timer，如果没有激活的timer，则一直阻塞
8. `Run check handles` 与prepare对应，执行注册的 check 回调事件
9. `Call close callbacks` 关闭回调调用，比如使用 uv_close关闭，就是在这里返回回调事件

## 1.1. loop相关API

```C++
struct uv_loop_s {
  void* data;/* User data - use this for whatever. */
  unsigned int active_handles;/* Loop reference counting. */
  void* handle_queue[2];
  union {
    void* unused;
    unsigned int count;
  } active_reqs;
  void* internal_fields;/* Internal storage for future extensions. */
  unsigned int stop_flag;/* Internal flag to signal loop stop. */
  UV_LOOP_PRIVATE_FIELDS
};

uv_loop_t* uv_default_loop(void);
void* uv_loop_set_data(uv_loop_t* loop, void* data)
void* uv_loop_get_data(const uv_loop_t* loop)
size_t uv_loop_size(void);
int uv_loop_alive(const uv_loop_t* loop);

int uv_loop_init(uv_loop_t* loop);
int uv_loop_close(uv_loop_t* loop);
int uv_loop_configure(uv_loop_t* loop, uv_loop_option option, ...);
int uv_loop_fork(uv_loop_t* loop);

int uv_run(uv_loop_t*, uv_run_mode mode);
void uv_stop(uv_loop_t*);

void uv_update_time(uv_loop_t*);
uint64_t uv_now(const uv_loop_t*);
uint64_t uv_hrtime(void);

int uv_backend_fd(const uv_loop_t*);
int uv_backend_timeout(const uv_loop_t*);
```

- `uv_default_loop` 获取默认的loop
- `uv_loop_set_data` 设置设置用户数据
- `uv_loop_get_data` 获取用户数据
- `uv_loop_size` Returns the size of the uv_loop_t structure
- `uv_loop_alive` 是否存活的结果，Loop里边有任务(不管uv_run是否启动)，返回非0，否则返回0。
- `uv_loop_init` 队列初始化，参数loop必须allocate
- `uv_loop_close` 关闭队列，执行此函数的前提是所有任务必须都执行完成，如果有任务遗留，会返回UV_EBUSY.
- `uv_loop_configure` 设置loop的属性，此接口在uv_loop之前调用。目前option仅支持SIGPROF，阻塞指定信号。
- `uv_loop_fork` fork loop，windows不支持
- `uv_run` 运行loop,uv_run_mode:UV_RUN_DEFAULT/UV_RUN_ONCE/UV_RUN_NOWAIT
- `uv_stop` 停止uv_run运行，停止运行后，loop里的任务还是有效的，只是不在执行，再次调用uv_run还可以继续运行。
- `uv_update_time` 立刻刷新loop的当前时间loop->time，目的是使uv_now更精确
- `uv_now` 获取loop时间loop->time，单位毫秒，如果需要微妙进度，请使用uv_hrtime()
- `uv_hrtime` nanoseconds
- `uv_backend_fd` 返回对应文件句柄，windows不支持
- `uv_backend_timeout` 下一个Loop到来时间。单位毫秒。

## 1.2. 参考资料

1. [Design overview](http://docs.libuv.org/en/v1.x/design.html)
2. [loop](http://docs.libuv.org/en/v1.x/loop.html)
3. [uv_loop_t](http://docs.libuv.org/en/v1.x/loop.html)
4. [uv_handle_t](http://docs.libuv.org/en/v1.x/handle.html)
5. [libuv的源码分析](cnblogs.com/watercoldyi/p/5675180.html)
6. [IO模型及select、poll、epoll和kqueue的区别](cnblogs.com/linganxiong/p/5583415.html)
7. [Reactor模式--VS--Proactor模式](blog.csdn.net/wenbingoon/article/details/9880365)
