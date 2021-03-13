# 1. 网络

- [1. 网络](#1-网络)
  - [1.1. 简介](#11-简介)
  - [1.2. UDP](#12-udp)
  - [1.3. stream](#13-stream)
    - [1.4. TCP](#14-tcp)
    - [1.5. PIPE](#15-pipe)
    - [1.6. TTY](#16-tty)
  - [参考资料](#参考资料)

## 1.1. 简介

&emsp;&emsp;libuv中的网络与直接使用BSD套接字接口没什么不同，有些东西更容易，都是非阻塞的，但概念保持不变。此外，libuv还提供实用程序功能来抽象恼人的，重复的和低级别的任务，例如使用BSD套接字结构设置套接字，DNS查找以及调整各种套接字参数。这里网络相关的主要涉及TCP和UDP，由于TCP是Stream流，而PIPE和TTY也是Stream类，所以他们有共同的结构体，libuv将TCP、PIPE、TTY归为Stream流所以这里一并对PIPE和TTY的相关API也做一个简单介绍。

## 1.2. UDP

UDP handles 为客户端和服务器封装UDP通信。[参考](http://docs.libuv.org/en/v1.x/udp.html)

```C++

typedef struct uv_udp_s uv_udp_t;
/* uv_udp_t is a subclass of uv_handle_t. */
struct uv_udp_s {
  UV_HANDLE_FIELDS
  size_t send_queue_size;//待发送块大小
  size_t send_queue_count;//待发送块个数
  UV_UDP_PRIVATE_FIELDS
};

int uv_udp_init(uv_loop_t*, uv_udp_t* handle);//初始化handle
int uv_udp_init_ex(uv_loop_t*, uv_udp_t* handle, unsigned int flags);
int uv_udp_open(uv_udp_t* handle, uv_os_sock_t sock);//打开init后的udp句柄
int uv_udp_bind(uv_udp_t* handle,const struct sockaddr* addr,unsigned int flags);//绑定一个ip和端口
int uv_udp_send(uv_udp_send_t* req,uv_udp_t* handle, const uv_buf_t bufs[], unsigned int nbufs, const struct sockaddr* addr, uv_udp_send_cb send_cb);//发送数据，发到缓存
int uv_udp_try_send(uv_udp_t* handle,const uv_buf_t bufs[], unsigned int nbufs, const struct sockaddr* addr);//直接发送数据
size_t uv_udp_get_send_queue_size(const uv_udp_t* handle);//缓存数据大小
size_t uv_udp_get_send_queue_count(const uv_udp_t* handle);//缓存数据个数

int uv_udp_recv_start(uv_udp_t* handle,uv_alloc_cb alloc_cb, uv_udp_recv_cb recv_cb);//接受数据
int uv_udp_recv_stop(uv_udp_t* handle);//停止接受数据
```

```C++
int uv_udp_getsockname(const uv_udp_t* handle,struct sockaddr* name,int* namelen);//获取本地ip和端口
int uv_udp_set_membership(uv_udp_t* handle,const char* multicast_addr,const char* interface_addr,uv_membership membership);//设置组播角色，接受还是发送

int uv_udp_set_multicast_loop(uv_udp_t* handle, int on);//是否接受组播
int uv_udp_set_multicast_ttl(uv_udp_t* handle, int ttl);//设置足膜的ttl(time to live)值
int uv_udp_set_multicast_interface(uv_udp_t* handle, const char* interface_addr);//设置组播发送或接受地址

int uv_udp_set_broadcast(uv_udp_t* handle, int on);//设置是否接受广播

int uv_udp_set_ttl(uv_udp_t* handle, int ttl);//设置udp的ttl
int uv_ip4_addr(const char* ip, int port, struct sockaddr_in* addr);//ipv4,port -> addr
int uv_ip6_addr(const char* ip, int port, struct sockaddr_in6* addr);//ipv6,port -> addr6
int uv_ip4_name(const struct sockaddr_in* src, char* dst, size_t size);//addr4 -> ipv4,prot
int uv_ip6_name(const struct sockaddr_in6* src, char* dst, size_t size);//addr6 -> ipv6,prot
int uv_inet_ntop(int af, const void* src, char* dst, size_t size);//将点分十进制的ip地址转化为用于网络传输的数值格式
int uv_inet_pton(int af, const char* src, void* dst);//将数值格式转化为点分十进制的ip地址格式
```

## 1.3. stream

&emsp;&emsp;下边是TCP、PIPE、TTY的结构体。他们有相同的结构 UV_HANDLE_FIELDS，UV_STREAM_FIELDS，所以TCP、PIPE、TTY可以强制转换成Stream类型。

```C++
#define UV_STREAM_FIELDS                                                      \
  /* number of bytes queued for writing */                                    \
  size_t write_queue_size;                                                    \
  uv_alloc_cb alloc_cb;                                                       \
  uv_read_cb read_cb;                                                         \
  /* private */                                                               \
  UV_STREAM_PRIVATE_FIELDS

/* uv_stream_t is a subclass of uv_handle_t.
 * uv_stream is an abstract class.
 * uv_stream_t is the parent class of uv_tcp_t, uv_pipe_t and uv_tty_t. */
struct uv_stream_s {
  UV_HANDLE_FIELDS
  UV_STREAM_FIELDS
};
/*uv_pipe_t is a subclass of uv_stream_t.
 * Representing a pipe stream or pipe server. On Windows this is a Named
 * Pipe. On Unix this is a Unix domain socket.*/
struct uv_pipe_s {
  UV_HANDLE_FIELDS
  UV_STREAM_FIELDS
  int ipc; /* non-zero if this pipe is used for passing handles */
  UV_PIPE_PRIVATE_FIELDS
};
/*
 * uv_tty_t is a subclass of uv_stream_t.
 * Representing a stream for the console.
 */
struct uv_tty_s {
  UV_HANDLE_FIELDS
  UV_STREAM_FIELDS
  UV_TTY_PRIVATE_FIELDS
};
/*
 * uv_tcp_t is a subclass of uv_stream_t.
 * Represents a TCP stream or TCP server.
 */
struct uv_tcp_s {
  UV_HANDLE_FIELDS
  UV_STREAM_FIELDS
  UV_TCP_PRIVATE_FIELDS
};

/* uv_connect_t is a subclass of uv_req_t. */
struct uv_connect_s {
  UV_REQ_FIELDS
  uv_connect_cb cb;
  uv_stream_t* handle;
  UV_CONNECT_PRIVATE_FIELDS
};
```

Stream基类相关
这里边的函数可以被PIPE、TTY、TCP调用，不同的类型可能会有一些差异，但是表示的功能相同。[官网API](http://docs.libuv.org/en/v1.x/stream.html)

```C++
int uv_shutdown(uv_shutdown_t* req, uv_stream_t* handle, uv_shutdown_cb cb)//关闭连接，此关闭会等待写操作完成。

int uv_listen(uv_stream_t* stream, int backlog, uv_connection_cb cb);//开始监听，backlog表示监听个数。同linux的listen(2)
int uv_accept(uv_stream_t* server, uv_stream_t* client);//接收请求。此接口一般在listen的回调里边处理

//1.38.0版本修改： 调用多次时，返回 UV_EALREADY，已经关闭时，返回UV_EINVAL。之前的版本window都返回UV_EALREADY，linux都返回UV_EINVAL
//参数：uv_read_cb 回调数据，回调数据的内存在 alloc_cb 中申请。出发 uv_read_cb 之前先触发 uv_alloc_cb 获取存放数据内存，如果内存小于缓存数据，则读取指定内存大小数据
int uv_read_start(uv_stream_t*,uv_alloc_cb alloc_cb,uv_read_cb read_cb);//开始读数据
int uv_read_stop(uv_stream_t*);//停止读数据

int uv_is_readable(const uv_stream_t* handle);//是否可读
int uv_is_writable(const uv_stream_t* handle);//是否可写
int uv_stream_set_blocking(uv_stream_t* handle, int blocking);//设置阻塞或非阻塞(discard，后续默认都是非阻塞)

//写数据，buffer空间在回调之前必须保持有效。
int uv_write(uv_write_t* req,uv_stream_t* handle,const uv_buf_t bufs[],unsigned int nbufs,uv_write_cb cb);
//uv_write的扩展,为了pipe发送handles，此pipe必须使用ipc==1初始化。
//Note： send_handle must be a TCP socket or pipe, which is a server or a connection (listening or connected state). Bound sockets or pipes will be assumed to be servers.
int uv_write2(uv_write_t* req,uv_stream_t* handle,const uv_buf_t bufs[],unsigned int nbufs,uv_stream_t* send_handle,uv_write_cb cb);

int uv_try_write(uv_stream_t* handle,const uv_buf_t bufs[],unsigned int nbufs);//同uv_write, 并且不关联req的同步写操作，类似于直接调用系统API。
size_t uv_stream_get_write_queue_size(const uv_stream_t* stream);//Returns stream->write_queue_size.
```

### 1.4. TCP

[参考网址](http://docs.libuv.org/en/v1.x/tcp.html)

TCP handles are used to represent both TCP streams and servers.
uv_tcp_t is a ‘subclass’ of uv_stream_t.

- Server sockets proceed by
  - `uv_tcp_init` the TCP handle.
  - `uv_tcp_bind` it.
  - Call `uv_listen` on the handle to have a callback invoked whenever a new connection is   established by a client.
  - Use `uv_accept` to accept the connection.
  - Use stream operations to communicate with the client.

```C++
//Initialize the handle. No socket is created as of yet.
int uv_tcp_init(uv_loop_t*, uv_tcp_t* handle);
//flags：AF_UNIX域、AF_INET域、AF_UNSPEC域等。AF_UNSPEC等同uv_tcp_init
int uv_tcp_init_ex(uv_loop_t*, uv_tcp_t* handle, unsigned int flags);
//Open an existing file descriptor or SOCKET as a TCP handle.
int uv_tcp_open(uv_tcp_t* handle, uv_os_sock_t sock);
//https://www.cnblogs.com/wajika/p/6573014.html
int uv_tcp_nodelay(uv_tcp_t* handle, int enable);//Enable TCP_NODELAY, which disables Nagle’s algorithm.
//Enable / disable TCP keep-alive. delay is the initial delay in seconds, ignored when enable is zero.
int uv_tcp_keepalive(uv_tcp_t* handle, int enable, unsigned int delay);
//uv_accept是否影响上次握手，true:不影响,false:影响.默认为true。不影响也有容量限制，取决于uv_listen的backlog参数
int uv_tcp_simultaneous_accepts(uv_tcp_t* handle, int enable);
//Bind the handle to an address and port. addr should point to an initialized struct sockaddr_in or struct sockaddr_in6.
//flags可以包含UV_TCP_IPV6ONLY，在这种情况下，禁用双栈支持并且仅使用IPv6。
int uv_tcp_bind(uv_tcp_t* handle, const struct sockaddr* addr, unsigned int flags);
////获取本地地址 name point to a valid and big enough chunk of memory, structsockaddr_storage is recommended for IPv4 and IPv6 support.
int uv_tcp_getsockname(const uv_tcp_t* handle, struct sockaddr* name, int* namelen);
int uv_tcp_getpeername(const uv_tcp_t* handle, struct sockaddr* name, int* namelen);//获取对方地址
//客户端与tcp的Server建立连接。Provide an initialized TCP handle and an uninitialized uv_connect_t.
int uv_tcp_connect(uv_connect_t* req,uv_tcp_t* handle, const struct sockaddr* addr, uv_connect_cb cb);
```

### 1.5. PIPE

PIPE
Pipe handles provide an abstraction over local domain sockets on Unix and named pipes on Windows.

uv_pipe_t is a ‘subclass’ of uv_stream_t.

```C++
int uv_pipe_init(uv_loop_t*, uv_pipe_t* handle, int ipc);//初始化一个pipe，ipc表示是否跨进程。
int uv_pipe_open(uv_pipe_t*, uv_file file);//打开已经存在file或HANDLE
int uv_pipe_bind(uv_pipe_t* handle, const char* name);//绑定路径(Unix)或者名称(Windows)
void uv_pipe_connect(uv_connect_t* req, uv_pipe_t* handle, const char* name, uv_connect_cb cb);//连接pipe
//获取pipe名字，buffer不含结束符，size不包含结束符。
int uv_pipe_getsockname(const uv_pipe_t* handle, char* buffer, size_t* size);
int uv_pipe_getpeername(const uv_pipe_t* handle, char* buffer, size_t* size);//获取对方pipe的名称
void uv_pipe_pending_instances(uv_pipe_t* handle, int count);//设置待处理句柄个数，仅适用于window

//First - call uv_pipe_pending_count(), if it’s > 0 then initialize a handle of the given type, returned by uv_pipe_pending_type() and call uv_accept(pipe, handle).
int uv_pipe_pending_count(uv_pipe_t* handle);//Pipe待处理任务个数，uv_accept使用之前调用。
uv_handle_type uv_pipe_pending_type(uv_pipe_t* handle);

int uv_pipe_chmod(uv_pipe_t* handle, int flags);//此函数是同步的，设置pipe的权限。
```

### 1.6. TTY

[TTY](http://docs.libuv.org/en/v1.x/tty.html)
[官网](http://docs.libuv.org/en/v1.x/tty.html)

TTY handles represent a stream for the console.
uv_tty_t is a ‘subclass’ of uv_stream_t.

&emsp;&emsp;在Linux中，TTY也许是跟终端有关系的最为混乱的术语。TTY是TeleTYpe的一个老缩写。Teletypes，或者teletypewriters，原来指的是电传打字机，是通过串行线用打印机键盘通过阅读和发送信息的东西，和古老的电报机区别并不是很大。之后，当计算机只能以批处理方式运行时（当时穿孔卡片阅读器是唯一一种使程序载入运行的方式），电传打字机成为唯一能够被使用的“实时”输入/输出设备。最终，电传打字机被键盘和显示器终端所取代，但在终端或TTY接插的地方，操作系统仍然需要一个程序来监视串行端口。一个getty“Get TTY”的处理过程是：一个程序监视物理的TTY/终端接口。对一个虚拟网络控制台（VNC）来说，一个伪装的TTY(Pseudo-TTY，即假冒的TTY，也叫做“PTY”）是等价的终端。当你运行一个xterm(终端仿真程序）或GNOME终端程序时，PTY对虚拟的用户或者如xterm一样的伪终端来说，就像是一个TTY在运行。“Pseudo”的意思是“duplicating in a fake way”（用伪造的方法复制），它相比“virtual”或“emulated”更能真实的说明问题。而在的计算中，它却处于被放弃的阶段。

&emsp;&emsp;串行端口终端（Serial Port Terminal）是使用计算机串行端口连接的终端设备。计算机把每个串行端口都看作是一个字符设备。有段时间这些串行端口设备通常被称为终端设备，因为那时它的最大用途就是用来连接终端。这些串行端口所对应的设备名称是/dev/tts/0（或/dev/ttyS0),/dev/tts/1（或/dev/ttyS1）等，设备号分别是（4,0），（4,1）等，分别对应于DOS系统下的COM1、COM2等。若要向一个端口发送数据，可以在命令行上把标准输出重定向到这些特殊文件名上即可。例如，在命令行提示符下键入：echo test > /dev/ttyS1会把单词”test”发送到连接在ttyS1(COM2）端口的设备上。

```C++
typedef enum {
  /* Initial/normal terminal mode */
  UV_TTY_MODE_NORMAL,
  /* Raw input mode (On Windows, ENABLE_WINDOW_INPUT is also enabled) */
  UV_TTY_MODE_RAW,
  /* Binary-safe I/O mode for IPC (Unix-only) */
  UV_TTY_MODE_IO
} uv_tty_mode_t;

// Initialize a new TTY stream with the given file descriptor. Usually the file descriptor will be:
// 0 = stdin      1 = stdout        2 = stderr
// readable, specifies if you plan on calling uv_read_start() with this stream. stdin is readable, stdout is not.

// On Unix this function will determine the path of the fd of the terminal using ttyname_r(3), open it, and use it if the passed file descriptor refers to a TTY. This lets libuv put the tty in non-blocking mode without affecting other processes that share the tty.

// This function is not thread safe on systems that don’t support ioctl TIOCGPTN or TIOCPTYGNAME, for instance OpenBSD and Solaris.

// Note: If reopening the TTY fails, libuv falls back to blocking writes for non-readable TTY streams.
// Changed in version 1.9.0:: the path of the TTY is determined by ttyname_r(3). In earlier versions libuv opened /dev/ttyinstead.
// Changed in version 1.5.0:: trying to initialize a TTY stream with a file descriptor that refers to a file returns UV_EINVAL on UNIX.
int uv_tty_init(uv_loop_t*, uv_tty_t*, uv_file fd, int readable);
int uv_tty_set_mode(uv_tty_t*, uv_tty_mode_t mode);//Set the TTY using the specified terminal mode.

//To be called when the program exits. Resets TTY settings to default values for the next process to take over.
//This function is async signal-safe on Unix platforms but can fail with error code UV_EBUSY if you call it when execution is inside uv_tty_set_mode().
int uv_tty_reset_mode(void);
int uv_tty_get_winsize(uv_tty_t*, int* width, int* height);//获取window大小,On success it returns 0.

uv_handle_type uv_guess_handle(uv_file file);
```

## 参考资料

1. [stream](http://docs.libuv.org/en/v1.x/stream.html)
2. [tty](http://docs.libuv.org/en/v1.x/tty.html)
3. [tcp](http://docs.libuv.org/en/v1.x/tcp.html)
4. [udp](http://docs.libuv.org/en/v1.x/udp.html)
5. [pipe](docs.libuv.org/en/v1.x/pipe.html)
