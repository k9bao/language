# windows的socket解析

```C++
// Windows Socket API 缩写为 WSA
//功能：创建完成端口和关联完成端口
HANDLE WINAPI CreateIoCompletionPort(
   __in   HANDLE FileHandle,              // 已经打开的文件句柄或者空句柄，一般是客户端的句柄
   __in   HANDLE ExistingCompletionPort,  // 已经存在的IOCP句柄
   __in   ULONG_PTR CompletionKey,        // 完成键，包含了指定I/O完成包的指定文件
   __in   DWORD NumberOfConcurrentThreads // 真正并发同时执行最大线程数，一般推介是CPU核心数*2
);
/*
功能：获取队列完成状态
返回值：调用成功，则返回非零数值，相关数据存于lpNumberOfBytes、lpCompletionKey、lpoverlapped变量中。失败则返回零值。
*/
BOOL GetQueuedCompletionStatus(
    HANDLE   CompletionPort,           //完成端口句柄
    LPDWORD   lpNumberOfBytes,         //一次I/O操作所传送的字节数
    PULONG_PTR   lpCompletionKey,      //当文件I/O操作完成后，用于存放与之关联的callback
    LPOVERLAPPED   *lpOverlapped,      //IOCP特定的结构体，用户对象指针
    DWORD   dwMilliseconds);           //调用者的等待时间

//用于IOCP的特定函数
typedef struct _OVERLAPPEDPLUS{
    OVERLAPPED ol;      //一个固定的用于处理网络消息事件返回值的结构体变量
    SOCKET s, sclient;  
    int OpCode;  //用来区分本次消息的操作类型（在完成端口的操作里面，是以消息通知系统，读数据/写数据，都是要发这样的消息结构体过去的）
    WSABUF wbuf;　　　　 //读写缓冲区结构体变量 
    DWORD dwBytes, dwFlags; //一些在读写时用到的标志性变量 
}OVERLAPPEDPLUS;
//功能：投递一个队列完成状态
BOOL PostQueuedCompletionStatus( 
　　HANDLE CompletlonPort, //指定想向其发送一个完成数据包的完成端口对象
　　DW0RD dwNumberOfBytesTrlansferred, //指定—个值,直接传递给GetQueuedCompletionStatus 函数中对应的参数 
　　DWORD dwCompletlonKey, //指定—个值,直接传递给GetQueuedCompletionStatus函数中对应的参数
　　LPOVERLAPPED lpoverlapped, ); //指定—个值,直接传递给GetQueuedCompletionStatus 函数中对应的参数
```

```C++
typedef struct _WSABUF {
    ULONG len;     /* the length of the buffer */
    _Field_size_bytes_(len) CHAR FAR *buf; /* the pointer to the buffer */
} WSABUF, FAR * LPWSABUF;
```

- `len` : 长度
- `buf` : buf指针

## WSARecv

```C++
#if INCL_WINSOCK_API_PROTOTYPES
WINSOCK_API_LINKAGE
int
WSAAPI
WSARecv(
    _In_ SOCKET s,
    _In_reads_(dwBufferCount) __out_data_source(NETWORK) LPWSABUF lpBuffers,
    _In_ DWORD dwBufferCount,
    _Out_opt_ LPDWORD lpNumberOfBytesRecvd,
    _Inout_ LPDWORD lpFlags,
    _Inout_opt_ LPWSAOVERLAPPED lpOverlapped,
    _In_opt_ LPWSAOVERLAPPED_COMPLETION_ROUTINE lpCompletionRoutine
    );
#endif /* INCL_WINSOCK_API_PROTOTYPES */
```

- `s [in]` : socket描述符
- `lpBuffers [in]` : 一个指向 WSABUF 结构数组的指针。
- `dwBufferCount` : lpBuffers 数组中 WSABUF 结构的大小。
- `lpNumberOfBytesRecvd [out]` : 如果发送操作立即完成，则为一个指向所发送数据字节数的指针。
- `dwFlags [in]` : 标志位。
- `lpOverlapped [in]` : 指向WSAOVERLAPPED结构的指针（对于非重叠套接口则忽略）。
- `lpCompletionRoutine [in]` : 一个指向发送操作完成后调用的完成例程的指针。（对于非重叠套接口则忽略）。
