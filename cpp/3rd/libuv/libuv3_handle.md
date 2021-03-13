# 1. handle

- [1. handle](#1-handle)
  - [1.1. handle对象](#11-handle对象)
  - [1.2. 通用接口API](#12-通用接口api)
  - [1.3. uv_async_t](#13-uv_async_t)
  - [1.4. idle/prepare/check](#14-idlepreparecheck)
  - [定时器](#定时器)
  - [1.5. 参考资料](#15-参考资料)

&emsp;&emsp;loop的调度一般通过二种模式调度，一种是Handle，一种是Request。Handle一般关联一个长期任务，request关联一个一次性任务。比如一个socket的监听任务，就是uv_udp_t（handle）,而一次数据的发送就是一个请求任务，关联请求句柄uv_udp_send_t（request）。

## 1.1. handle对象

```c++
#define UV_HANDLE_FIELDS                                                      \
  /* public */                                                                \
  void* data;                                                                 \
  /* read-only */                                                             \
  uv_loop_t* loop;                                                            \
  uv_handle_type type;                                                        \
  /* private */                                                               \
  uv_close_cb close_cb;                                                       \
  void* handle_queue[2];                                                      \
  union {                                                                     \
    int fd;                                                                   \
    void* reserved[4];                                                        \
  } u;                                                                        \
  UV_HANDLE_PRIVATE_FIELDS                                                    \

/* The abstract base class of all handles. */
struct uv_handle_s {
  UV_HANDLE_FIELDS
};
/* Handle types. */
typedef struct uv_loop_s uv_loop_t;
typedef struct uv_handle_s uv_handle_t;
typedef struct uv_dir_s uv_dir_t;
typedef struct uv_stream_s uv_stream_t;
typedef struct uv_tcp_s uv_tcp_t;
typedef struct uv_udp_s uv_udp_t;
typedef struct uv_pipe_s uv_pipe_t;
typedef struct uv_tty_s uv_tty_t;
typedef struct uv_poll_s uv_poll_t;
typedef struct uv_timer_s uv_timer_t;       //
typedef struct uv_prepare_s uv_prepare_t;   //
typedef struct uv_check_s uv_check_t;       //
typedef struct uv_idle_s uv_idle_t;         //
typedef struct uv_async_s uv_async_t;       //
typedef struct uv_process_s uv_process_t;
typedef struct uv_fs_event_s uv_fs_event_t;
typedef struct uv_fs_poll_s uv_fs_poll_t;
typedef struct uv_signal_s uv_signal_t;
```

## 1.2. 通用接口API

```C++
int uv_is_active(const uv_handle_t* handle);
int uv_is_closing(const uv_handle_t* handle);
void uv_close(uv_handle_t* handle, uv_close_cb close_cb);

void uv_ref(uv_handle_t*);
void uv_unref(uv_handle_t*);
int uv_has_ref(const uv_handle_t*);

size_t uv_handle_size(uv_handle_type type);
uv_handle_type uv_handle_get_type(const uv_handle_t* handle);
const char* uv_handle_type_name(uv_handle_type type);

uv_loop_t* uv_handle_get_loop(const uv_handle_t* handle);
void* uv_handle_get_data(const uv_handle_t* handle);
void uv_handle_set_data(uv_handle_t* handle, void* data);

int uv_send_buffer_size(uv_handle_t* handle, int* value);
int uv_recv_buffer_size(uv_handle_t* handle, int* value);

int uv_fileno(const uv_handle_t* handle, uv_os_fd_t* fd);
```

- `uv_is_active` 请求是否处于活动状态,活动状态返回非0.判读是否是活动状态标准如下
  - uv_async_t 启动后就处于active状态，uv_close 后，处于非active状态
  - uv_pipe_t, uv_tcp_t, uv_udp_t, etc. handle等I/O handle当处于读，写，状态时，处于active状态
  - 异步handle执行uv_*start后，就处于active状态，uv_*stop后，处于非active状态
  - 其他的hangle,比如 uv_foo_t,执行 vu_foo_start()函数的时候激活，执行uv_foo_stop()处于非激活状态
- `uv_is_closing`: 是否关闭，如果正在关闭或已经关闭返回非0
- `uv_close`:从endgame队列中删除,关闭指定句柄，回调可以为空（同步返回）,比如关闭定时器、idle、prepare等，一般和各种init成对出现。

- `uv_ref`: 指定句柄为ref状态，如果有处于ref状态的 loop 就不会退出，多次调用无效
- `uv_unref`: 与uv_ref相反，取消引用，多次无效(标志设置)，设置后不作为loop循环运行的条件
  - 一个应用场景就是，loop中，创建一个timer用于回收一些其他req或者handle的资源，创建timer之后，就可以通过uv_unref调用，减除timer的引用，当其他handle或者req全部退出的时候，loop也可以正常退出。
- `uv_has_ref`: Returns non-zero if the handle referenced, zero otherwise.

- `uv_handle_get_type`: Returns handle->type
  - `uv_handle_size`: Returns the size of the given handle type.
  - `uv_handle_type_name`: 类型名称, e.g. “pipe” (as in uv_pipe_t) for UV_NAMED_PIPE.

- `uv_handle_get_loop`：Returns handle->loop.
- `uv_handle_get_data`：Returns handle->data.
- `uv_handle_set_data`：Sets handle->data to data.

- `uv_send_buffer_size`：设置，发送缓存大小，注意：linux会设置二倍大小
- `uv_recv_buffer_size` : 获取，发送缓存大小，注意：linux会设置二倍大小

- `uv_fileno` : 提取句柄。比如：handle->io_watcher.fd;

## 1.3. uv_async_t

```C++
//uv_req_s详见 req.md
#define UV_ASYNC_PRIVATE_FIELDS                                               \
  struct uv_req_s async_req;                                                  \
  uv_async_cb async_cb;                                                       \
  /* char to avoid alignment issues */                                        \
  char volatile async_sent;
typedef struct uv_async_s {
  UV_HANDLE_FIELDS
  UV_ASYNC_PRIVATE_FIELDS
}uv_async_t;
typedef void (*uv_async_cb)(uv_async_t* handle);
int uv_async_init(uv_loop_t*, uv_async_t* async, uv_async_cb async_cb);
int uv_async_send(uv_async_t* async);
```

异步通信handle，从结构来看，既包含uv_req_s，也包含 uv_handle_s(UV_HANDLE_FIELDS)

- uv_async_init：初始化async句柄。 A NULL callback is allowed.
  - Returns:0 on success, or an error code < 0 on failure.Note
  - Unlike other handle initialization functions, it immediately starts the handle.

- uv_async_send：通知async句柄，此函数是线程安全的，但是多次发送uv_async_send有可能只触发一次uv_async_cb。
  - Returns:0 on success, or an error code < 0 on failure.Note
  - It’s safe to call this function from any thread. The callback will be called on the loop thread.

## 1.4. idle/prepare/check

```C++
#define UV_IDLE_PRIVATE_FIELDS                                                \
  uv_idle_t* idle_prev;                                                       \
  uv_idle_t* idle_next;                                                       \
  uv_idle_cb idle_cb;
typedef struct uv_idle_s {
  UV_HANDLE_FIELDS
  UV_IDLE_PRIVATE_FIELDS
}uv_idle_t;
void (*uv_idle_cb)(uv_idle_t* handle)
int uv_idle_init(uv_loop_t*, uv_idle_t* idle);
int uv_idle_start(uv_idle_t* idle, uv_idle_cb cb);
int uv_idle_stop(uv_idle_t* idle);

#define UV_PREPARE_PRIVATE_FIELDS                                             \
  uv_prepare_t* prepare_prev;                                                 \
  uv_prepare_t* prepare_next;                                                 \
  uv_prepare_cb prepare_cb;

typedef struct uv_prepare_s {
  UV_HANDLE_FIELDS
  UV_PREPARE_PRIVATE_FIELDS
}uv_prepare_t;
typedef void (*uv_prepare_cb)(uv_prepare_t* handle);
int uv_prepare_init(uv_loop_t*, uv_prepare_t* prepare);
int uv_prepare_start(uv_prepare_t* prepare, uv_prepare_cb cb);
int uv_prepare_stop(uv_prepare_t* prepare);

#define UV_CHECK_PRIVATE_FIELDS                                               \
  uv_check_t* check_prev;                                                     \
  uv_check_t* check_next;                                                     \
  uv_check_cb check_cb;
typedef struct uv_check_s {
  UV_HANDLE_FIELDS
  UV_CHECK_PRIVATE_FIELDS
}uv_check_s;
typedef void (*uv_check_cb)(uv_check_t* handle);
int uv_check_init(uv_loop_t*, uv_check_t* check);
int uv_check_start(uv_check_t* check, uv_check_cb cb);
int uv_check_stop(uv_check_t* check);
```

uv__handle_init() -> uv__handle_start() -> uv__handle_stop() - > uv__handle_closing() -> uv_want_endgame()-> uv__handle_close()

1, init，初始化，返回句柄。
2，start，添加句柄到执行队列，准备执行，设置回调，如果handle已经存在，添加无效。不管是否有效，都返回0
3，stop，从执行队列移到endgame队列。

一般idle作为辅助消息起作用，一般和其他消息一块运行，用来检测其他消息是否真正执行。或者是否执行完成。

## 定时器

```c++
#define UV_TIMER_PRIVATE_FIELDS                                               \
  void* heap_node[3];                                                         \
  int unused;                                                                 \
  uint64_t timeout;                                                           \
  uint64_t repeat;                                                            \
  uint64_t start_id;                                                          \
  uv_timer_cb timer_cb;

typedef struct uv_timer_s {
  UV_HANDLE_FIELDS
  UV_TIMER_PRIVATE_FIELDS
}uv_timer_t;

void (*uv_timer_cb)(uv_timer_t* handle);
int uv_timer_init(uv_loop_t*, uv_timer_t* handle);

int uv_timer_start(uv_timer_t* handle, uv_timer_cb cb, uint64_t timeout, uint64_t repeat);
int uv_timer_stop(uv_timer_t* handle);

int uv_timer_again(uv_timer_t* handle);

void uv_timer_set_repeat(uv_timer_t* handle, uint64_t repeat);
uint64_t uv_timer_get_repeat(const uv_timer_t* handle);
```

实现定时器功能

- `uv_timer_init` 出事后初始化定时器
- `uv_timer_start`
  - 开始计时器，timeout:首次触发时间，repeat：重复间隔时长(milliseconds)
  - 如果连续调用二次start，第二次覆盖第一次。
  - 如果两个定时间触发时间点相同，按照start执行的先后顺序触发。
  - cb不可以为空，否则失败
  - uv_timer_t在uv_timer_stop之前必须有效，否则触发崩溃
- `uv_timer_stop` 停止计时器，可以使用uv_close(uv_handle_t*, NULL)代替。
- `uv_timer_again` 重新开始
- `uv_timer_set_repeat` 设置 repeat 变量， 单位：milliseconds
- `uv_timer_get_repeat` 获取 repeat 变量， 单位：milliseconds

## 1.5. 参考资料

1. [uv_handle_t](http://docs.libuv.org/en/v1.x/handle.html)
2. [uv_async_t](http://docs.libuv.org/en/v1.x/async.html)
3. [prepare](http://docs.libuv.org/en/v1.x/prepare.html)
4. [check](http://docs.libuv.org/en/v1.x/check.html)
5. [idle](http://docs.libuv.org/en/v1.x/idle.html)
