# 1. libuv-进程

- [1. libuv-进程](#1-libuv-进程)
  - [1.1. 简介](#11-简介)
  - [1.2. API](#12-api)
  - [1.3. 创建进程](#13-创建进程)
  - [1.4. 父子进程分离](#14-父子进程分离)
  - [1.5. 向进程发送信号](#15-向进程发送信号)
  - [1.6. 子进程I/O](#16-子进程io)
  - [1.7. 管道](#17-管道)
  - [1.8. 参考文章](#18-参考文章)

## 1.1. 简介

&emsp;&emsp;libuv提供子进程的管理工作，抽象平台的差异性，允许通过stream和pipe进行线程之间的交互

&emsp;&emsp;unix建议每个进程只做一件事并且做好。所以一个进程常常使用多个子进程来完成各种任务，类似于shell中使用多个pipe。与线程共享内存类似，含有消息的多进程模型也很简单。

&emsp;&emsp;基于事件的程序想利用现代计算机的多核功能是有一定挑战性的。在多线程程序里，内核可以通过调度把不同的线程分配到不同的系统核core上来提供core的利用率。但是libuv的loop只有一个线程，通过变通的使用多进程，每个进程上运行一个loop，每个进程关联到不同的core上以达到利用多核系统。

## 1.2. API

```C++
typedef enum {
  UV_IGNORE         = 0x00,
  UV_CREATE_PIPE    = 0x01,
  UV_INHERIT_FD     = 0x02,
  UV_INHERIT_STREAM = 0x04,

  UV_READABLE_PIPE  = 0x10,//只读 pipe, 前提是 UV_CREATE_PIPE 被设置
  UV_WRITABLE_PIPE  = 0x20,//只写 pipe, 前提是 UV_CREATE_PIPE 被设置

  UV_NONBLOCK_PIPE  = 0x40,//非阻塞 pipe, 前提是 UV_CREATE_PIPE 被设置
  UV_OVERLAPPED_PIPE = 0x40 /* old name, for compatibility */
} uv_stdio_flags;

typedef struct uv_stdio_container_s {
  uv_stdio_flags flags;

  union {
    uv_stream_t* stream;
    int fd;
  } data;
} uv_stdio_container_t;

enum uv_process_flags {
  UV_PROCESS_SETUID = (1 << 0),//手动设置 UID ,window无效
  UV_PROCESS_SETGID = (1 << 1),//手动设置 GID ,window无效

  UV_PROCESS_WINDOWS_VERBATIM_ARGUMENTS = (1 << 2),//参数不使用引号，并且没有任何转义，window专用
  
  UV_PROCESS_DETACHED = (1 << 3),//主子进程分离

  UV_PROCESS_WINDOWS_HIDE = (1 << 4),//window子进程隐藏创建
  UV_PROCESS_WINDOWS_HIDE_CONSOLE = (1 << 5),//window子进程隐藏 console
  UV_PROCESS_WINDOWS_HIDE_GUI = (1 << 6)//window子进程隐藏 gui
};
```

```c++
typedef struct uv_process_options_s {
  uv_exit_cb exit_cb; /* 回调*/
  const char* file;   /* 进程执行的可执行文件*/
  char** args; //参数，第一个参数必须同file,最后一个参数一NULL结束
  char** env;//环境变量，为NULL，继承父进程的环境变量。类似于 VAR=VALUE 格式的字符数组
  const char* cwd;//设置运行目录,默认是当前目录
  unsigned int flags;//具体详见 enum uv_process_flags，比如linux指定UID，GID，window隐藏console和GUI,detache等

  int stdio_count;
  uv_stdio_container_t* stdio;//指定 stdin,stdout,stderr

  uv_uid_t uid;//flags包含 UV_PROCESS_SETUID 时，手动指定uid，linux有效
  uv_gid_t gid;//flags包含 UV_PROCESS_SETGID 时，手动指定gid，linux有效
} uv_process_options_t;
```

```C++
//创建子进程
UV_EXTERN int uv_spawn(uv_loop_t* loop,
                       uv_process_t* handle,
                       const uv_process_options_t* options);
//杀掉子进程,如果调用uv_process_kill关闭进程的话，别忘记调用uv_close(handle)函数。
UV_EXTERN int uv_process_kill(uv_process_t* handle, int signum);
UV_EXTERN int uv_kill(int pid, int signum);//杀掉子进程
UV_EXTERN uv_pid_t uv_process_get_pid(const uv_process_t*);//获取子进程ID，实际就是 uv_process_t.pid
```

- signum
  - SIGINT : 程序终止(interrupt)信号, 在用户键入INTR字符(通常是Ctrl-C)时发出，用于通知前台进程组终止进程。
  - SIGQUIT : 和SIGINT类似, 但由QUIT字符(通常是Ctrl-)来控制. 进程在因收到SIGQUIT退出时会产生core文件, 在这个意义上类似于一个程序错误信号。
  - SIGTERM : 程序结束(terminate)信号, 与SIGKILL不同的是该信号可以被阻塞和处理。通常用来要求程序自己正常退出，shell命令kill缺省产生这个信号。如果进程终止不了，我们才会尝试SIGKILL。
  - SIGSTOP : 停止(stopped)进程的执行. 注意它和terminate以及interrupt的区别:该进程还未结束, 只是暂停执行. 本信号不能被阻塞, 处理或忽略.
  - ...

## 1.3. 创建进程

使用uv_spawn来创建一个进程。代码如下：

`UV_EXTERN int uv_spawn(uv_loop_t* loop,uv_process_t* handle,const uv_process_options_t* options);`

```c++
uv_loop_t *loop;
uv_process_t child_req;
uv_process_options_t options;
void on_exit(uv_process_t *req, int64_t exit_status, int term_signal) {
    fprintf(stderr, "Process exited with status %" PRId64 ", signal %d\n", exit_status, term_signal);
    uv_close((uv_handle_t*) req, NULL);
}
int main() {
    loop = uv_default_loop();

    char* args[3];
    args[0] = "mkdir";
    args[1] = "test-dir";
    args[2] = NULL;

    options.exit_cb = on_exit;
    options.file = "mkdir";
    options.args = args;

    int r;
    if ((r = uv_spawn(loop, &child_req, &options))) {
        fprintf(stderr, "%s\n", uv_strerror(r));
        return 1;
    } else {
        fprintf(stderr, "Launched process with ID %d\n", child_req.pid);
    }

    return uv_run(loop, UV_RUN_DEFAULT);
}
```

- options 必须被初始化位 0，上边代码因为 options 是全局的，所以默认会被置为 0.
- args 的数组最后一个参数必须是 NULL。
- 完成调用之后， uv_process_t.pid 包含的就是进程ID.
- on_exit()在进程退出时，会触发。
- 在进程运行前，可以改变参数 uv_process_options_t

## 1.4. 父子进程分离

&emsp;&emsp;可以创建和父进程无关的子进程，使用UV_PROCESS_DETACHED标记。不受父进程的影响，父进程关闭不影响子进程的运行。可以用来写守候等不受父进程影响的子进程。举例如下：

```C++
int main() {
    loop = uv_default_loop();

    char* args[3];
    args[0] = "sleep";
    args[1] = "100";
    args[2] = NULL;

    options.exit_cb = NULL;
    options.file = "sleep";
    options.args = args;
    options.flags = UV_PROCESS_DETACHED;

    int r;
    if ((r = uv_spawn(loop, &child_req, &options))) {
        fprintf(stderr, "%s\n", uv_strerror(r));
        return 1;
    }
    fprintf(stderr, "Launched sleep with PID %d\n", child_req.pid);
    uv_unref((uv_handle_t*) &child_req);

    return uv_run(loop, UV_RUN_DEFAULT);
}
```

&emsp;&emsp;虽然使用了UV_PROCESS_DETACHED，但是句柄还在父进程那里，如果想完全失去读子进程的控制，需要调用uv_unref接口，完全释放对子进程的控制。

## 1.5. 向进程发送信号

&emsp;&emsp;在进程执行的过程中，可以通过发送信号来控制进程的执行。

UV_EXTERN int uv_process_kill(uv_process_t*, int signum);
UV_EXTERN int uv_kill(int pid, int signum);
如果调用uv_process_kill关闭进程的话，别忘记调用uv_close函数。

信号
libuv针对Unix提供了一些信号，有些也适用于Windows.

使用uv_signal_init()初始化，它与一个循环关联。要侦听该处理程序上的特定信号，使用uv_signal_start()处理程序函数。每个处理程序只能与一个信号编号相关联，连续调用 uv_signal_start()后边的覆盖先前信号。使用uv_signal_stop()停止信号监听。下边提供一个小例子供参考：

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <uv.h>

uv_loop_t* create_loop()
{
    uv_loop_t *loop = malloc(sizeof(uv_loop_t));
    if (loop) {
      uv_loop_init(loop);
    }
    return loop;
}

void signal_handler(uv_signal_t *handle, int signum)
{
    printf("Signal received: %d\n", signum);
    uv_signal_stop(handle);
}

// two signal handlers in one loop
void thread1_worker(void *userp)
{
    uv_loop_t *loop1 = create_loop();

    uv_signal_t sig1a, sig1b;
    uv_signal_init(loop1, &sig1a);
    uv_signal_start(&sig1a, signal_handler, SIGUSR1);

    uv_signal_init(loop1, &sig1b);
    uv_signal_start(&sig1b, signal_handler, SIGUSR1);

    uv_run(loop1, UV_RUN_DEFAULT);
}

// two signal handlers, each in its own loop
void thread2_worker(void *userp)
{
    uv_loop_t *loop2 = create_loop();
    uv_loop_t *loop3 = create_loop();

    uv_signal_t sig2;
    uv_signal_init(loop2, &sig2);
    uv_signal_start(&sig2, signal_handler, SIGUSR1);

    uv_signal_t sig3;
    uv_signal_init(loop3, &sig3);
    uv_signal_start(&sig3, signal_handler, SIGUSR1);

    while (uv_run(loop2, UV_RUN_NOWAIT) || uv_run(loop3, UV_RUN_NOWAIT)) {
    }
}

int main()
{
    printf("PID %d\n", getpid());

    uv_thread_t thread1, thread2;

    uv_thread_create(&thread1, thread1_worker, 0);
    uv_thread_create(&thread2, thread2_worker, 0);

    uv_thread_join(&thread1);
    uv_thread_join(&thread2);
    return 0;
}
上述监听回调会触发4次监听。

## 1.6. 子进程I/O

&emsp;&emsp;子进程有其自身的一组文件描述符，用0，1和2的分别表示stdin，stdout和stderr。有时我们可能希望子进程与父进程共享文件描述符。libuv 支持继承文件描述符。

```c++
int main() {
    loop = uv_default_loop();

    /* ... */

    options.stdio_count = 3;
    uv_stdio_container_t child_stdio[3];
    child_stdio[0].flags = UV_IGNORE;
    child_stdio[1].flags = UV_IGNORE;
    child_stdio[2].flags = UV_INHERIT_FD;
    child_stdio[2].data.fd = 2;
    options.stdio = child_stdio;

    options.exit_cb = on_exit;
    options.file = args[0];
    options.args = args;

    int r;
    if ((r = uv_spawn(loop, &child_req, &options))) {
        fprintf(stderr, "%s\n", uv_strerror(r));
        return 1;
    }

    return uv_run(loop, UV_RUN_DEFAULT);
}
```

UV_IGNORE 被重定向到/dev/null
UV_INHERIT_FD 继承父进程的I/O,需要在data.fd中，设置父进程的描述符。
同理，子进程可以继承父进程的Stream，举例如下：


    args[1] = NULL;

    /* ... finding the executable path and setting up arguments ... */

    options.stdio_count = 3;
    uv_stdio_container_t child_stdio[3];
    child_stdio[0].flags = UV_IGNORE;
    child_stdio[1].flags = UV_INHERIT_STREAM;
    child_stdio[1].data.stream = (uv_stream_t*) client;
    child_stdio[2].flags = UV_IGNORE;
    options.stdio = child_stdio;

    options.exit_cb = cleanup_handles;
    options.file = args[0];
    options.args = args;

    // Set this so we can close the socket after the child process exits.
    child_req.data = (void*) client;
    int r;
    if ((r = uv_spawn(loop, &child_req, &options))) {
        fprintf(stderr, "%s\n", uv_strerror(r));

## 1.7. 管道

管道不像linux下的符号“|”或者pipe，管道是用于两个进程之间进行通信的方法。具体如下：

父子管道：
父进程在创建子进程的时候uv_spawn，可以通过设置 uv_stdio_container_t.flags 参数，来设置父子进程之间的单向或双向通信， uv_stdio_container_t.flags可以是UV_CREATE_PIPE 、 UV_READABLE_PIPE or UV_WRITABLE_PIPE的一个或多个。可以同时设置多个，占不同的二进制位。UV_CREATE_PIPE 、 UV_READABLE_PIPE or UV_WRITABLE_PIPE都是从子进程的角度来看。

Arbitrary process IPC
Since domain sockets [1] can have a well known name and a location in the file-system they can be used for IPC between unrelated processes. The D-BUS system used by open source desktop environments uses domain sockets for event notification. Various applications can then react when a contact comes online or new hardware is detected. The MySQL server also runs a domain socket on which clients can interact with it.

When using domain sockets, a client-server pattern is usually followed with the creator/owner of the socket acting as the server. After the initial setup, messaging is no different from TCP, so we’ll re-use the echo server example.

pipe-echo-server/main.c

```C++
void remove_sock(int sig) {
    uv_fs_t req;
    uv_fs_unlink(loop, &req, PIPENAME, NULL);
    exit(0);
}

int main() {
    loop = uv_default_loop();

    uv_pipe_t server;
    uv_pipe_init(loop, &server, 0);

    signal(SIGINT, remove_sock);

    int r;
    if ((r = uv_pipe_bind(&server, PIPENAME))) {
        fprintf(stderr, "Bind error %s\n", uv_err_name(r));
        return 1;
    }
    if ((r = uv_listen((uv_stream_t*) &server, 128, on_new_connection))) {
        fprintf(stderr, "Listen error %s\n", uv_err_name(r));
        return 2;
    }
    return uv_run(loop, UV_RUN_DEFAULT);
}
```

We name the socket echo.sock which means it will be created in the local directory. This socket now behaves no different from TCP sockets as far as the stream API is concerned. You can test this server using socat:

$ socat - /path/to/socket
A client which wants to connect to a domain socket will use:

void uv_pipe_connect(uv_connect_t *req, uv_pipe_t *handle, const char *name, uv_connect_cb cb);
where name will be echo.sock or similar. On Unix systems, name must point to a valid file (e.g. /tmp/echo.sock). On Windows, name follows a \\?\pipe\echo.sock format.

## 1.8. 参考文章

1. [uv_process_t](http://docs.libuv.org/en/v1.x/process.html)
