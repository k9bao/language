# 1. libuv文件系统

- [1. libuv文件系统](#1-libuv文件系统)
  - [1.1. 简介](#11-简介)
  - [1.2. API说明](#12-api说明)
    - [内容结构](#内容结构)
    - [1.2.1. 文件句柄操作相关](#121-文件句柄操作相关)
    - [1.2.2. 文件直接操作](#122-文件直接操作)
    - [1.2.3. 目录操作相关](#123-目录操作相关)
    - [1.2.4. 文件属性相关](#124-文件属性相关)
    - [1.2.5. 文件链接相关](#125-文件链接相关)
    - [1.2.6. 文件监听相关](#126-文件监听相关)
  - [Demo](#demo)
  - [1.3. 参考资料](#13-参考资料)

## 1.1. 简介

&emsp;&emsp;文件相关的函数以uv_fs_开头，结构统一使用uv_fs_t，文件操作和网络操作不同，网络操作是底层异步API支持非阻塞，而文件系统内部是阻塞调用，通过线程池实现异步通知。文件操作有两种模式：同步模式(回调为NULL)和异步模式。

&emsp;&emsp;所有操作回调可以设置为空，如果回调为空执行同步返回，一般对文件的频繁读写，需要使用同步模式。因为异步不保证内容写入的先后顺序，如果对一个文件多次写入有先后顺序要求的话，只能通过同步模式进行。文件的异步回调方式，每个函数的操作都会从线程池里关联一个线程进行操作。

## 1.2. API说明

### 内容结构

```C++
/* uv_fs_t is a subclass of uv_req_t. */
struct uv_fs_s {
  UV_REQ_FIELDS
  uv_fs_type fs_type;
  uv_loop_t* loop;
  uv_fs_cb cb;
  ssize_t result;
  void* ptr;
  const char* path;
  uv_stat_t statbuf;  /* Stores the result of uv_fs_stat() and uv_fs_fstat(). */
  UV_FS_PRIVATE_FIELDS
};
```

### 1.2.1. 文件句柄操作相关

```c++
//清空请求，每次发起请求后，最后必须调用的函数，用于释放libuv可能申请的内存。
void uv_fs_req_cleanup(uv_fs_t* req);

//功能：构造一个uv_buf_t结构，uv_buf_t结构含有数据和长度，一般用于uv_fs_write.
uv_buf_t uv_buf_init(char* base, unsigned int len);

// 功能：打开文件，linux下封装open，window下封装CreateFileW （window仅支持二进制模式打开）。
// 参数：
//      req：输入输出参数，req->result成功的话，返回文件句柄(>0)，否则返回错误码(<0),下同
//      Flags,mode：unix flags,libuv在window环境之中，会自动对应到window的参数。
int uv_fs_open(uv_loop_t* loop, uv_fs_t* req, const char* path, int flags, int mode, uv_fs_cb cb);

// 功能：读多组数据，等价于preadv.
// 参数：
//     file:打开的文件句柄，bufs：数据数组(含数据和长度)，nbufs：bufs个数，offset：偏移量
int uv_fs_read(uv_loop_t* loop,uv_fs_t* req,uv_file file,const uv_buf_t bufs[], unsigned int nbufs, int64_t offset,uv_fs_cb cb);
// 功能：写多组数据，等同于pwritev.
int uv_fs_write(uv_loop_t* loop,uv_fs_t* req,uv_file file,uv_buf_t bufs[],unsigned int nbufs,int64_t offset,uv_fs_cb cb);
// 功能：关闭文件，linux下封装close。
int uv_fs_close(uv_loop_t* loop,uv_fs_t* req,uv_file file,uv_fs_cb cb);

//刷新文件，类似flushes
int uv_fs_fsync(uv_loop_t* loop,uv_fs_t* req,uv_file file, uv_fs_cb cb);
int uv_fs_fdatasync(uv_loop_t* loop,uv_fs_t* req,uv_file file, uv_fs_cb cb);//同上
int uv_fs_ftruncate(uv_loop_t* loop,uv_fs_t* req,uv_file file, int64_t offset, uv_fs_cb cb);//截取文件或者补充文件到指定大小

//返回结构特uv_fs_t的各种参数，具体详见uv_fs_t
uv_fs_type uv_fs_get_type(const uv_fs_t*);
ssize_t uv_fs_get_result(const uv_fs_t*);
void* uv_fs_get_ptr(const uv_fs_t*);
const char* uv_fs_get_path(const uv_fs_t*);
uv_stat_t* uv_fs_get_statbuf(uv_fs_t*);
```

### 1.2.2. 文件直接操作

```C++
//改名字
int uv_fs_rename(uv_loop_t* loop, uv_fs_t* req, const char* path, const char* new_path, uv_fs_cb cb);
//删除文件，封装的unlink(2).
int uv_fs_unlink(uv_loop_t* loop,uv_fs_t* req, const char* path, uv_fs_cb cb);
//拷贝文件
// flags：
// UV_FS_COPYFILE_EXCL: 如果目标文件已经存在，返回UV_EEXIST，不执行。
// UV_FS_COPYFILE_FICLONE: 如果目标文件存在，尝试创建copy-on-write reflink，如果平台不支持copy-on-write reflink，那就先删除，后拷贝
// UV_FS_COPYFILE_FICLONE_FORCE:同上，但是，如果平台不支持copy-on-write reflink，那就返回失败
int uv_fs_copyfile(uv_loop_t* loop, uv_fs_t* req, const char* path, const char* new_path, int flags, uv_fs_cb cb);
//等同于copyfile
int uv_fs_sendfile(uv_loop_t* loop, uv_fs_t* req, uv_file out_fd, uv_file in_fd, int64_t in_offset, size_t length, uv_fs_cb cb);
```

### 1.2.3. 目录操作相关

```C++
//创建目录
int uv_fs_mkdir(uv_loop_t* loop,uv_fs_t* req,const char* path,int mode,uv_fs_cb cb);
int uv_fs_mkdtemp(uv_loop_t* loop,uv_fs_t* req,const char* tpl,uv_fs_cb cb);
//删除目录
int uv_fs_rmdir(uv_loop_t* loop,uv_fs_t* req,const char* path,uv_fs_cb cb);
//浏览目录文件
int uv_fs_scandir(uv_loop_t* loop,uv_fs_t* req,const char* path,int flags,uv_fs_cb cb);
//目录下一个文件
int uv_fs_scandir_next(uv_fs_t* req, uv_dirent_t* ent);
```

### 1.2.4. 文件属性相关

```C++
int uv_fs_stat(uv_loop_t* loop, uv_fs_t* req, const char* path, uv_fs_cb cb);//同linux的stat
int uv_fs_fstat(uv_loop_t* loop,uv_fs_t* req,uv_file file,uv_fs_cb cb);//同linux的fstat
int uv_fs_lstat(uv_loop_t* loop,uv_fs_t* req,const char* path,uv_fs_cb cb);//同linux的lstat
int uv_fs_access(uv_loop_t* loop, uv_fs_t* req, const char* path, int mode, uv_fs_cb cb);//检测是否可用。同linux access(2). Windows GetFileAttributesW().
int uv_fs_utime(uv_loop_t* loop,uv_fs_t* req,const char* path,double atime,double mtime,uv_fs_cb cb);//修改文件时间
int uv_fs_futime(uv_loop_t* loop, uv_fs_t* req, uv_file file, double atime, double mtime, uv_fs_cb cb);//同上
int uv_fs_chmod(uv_loop_t* loop,uv_fs_t* req,const char* path,int mode,uv_fs_cb cb);//检测执行权限，等同于chmod
int uv_fs_fchmod(uv_loop_t* loop, uv_fs_t* req, uv_file file, int mode, uv_fs_cb cb);//检测执行权限，等同于fchmod
int uv_fs_chown(uv_loop_t* loop,uv_fs_t* req,const char* path,uv_uid_t uid,uv_gid_t gid,uv_fs_cb cb);//更改所有者及所属组，等同于chown()
int uv_fs_fchown(uv_loop_t* loop, uv_fs_t* req, uv_file file, uv_uid_t uid, uv_gid_t gid, uv_fs_cb cb);//更改所有者及所属组，等同于fchown()
int uv_fs_lchown(uv_loop_t* loop, uv_fs_t* req, const char* path, uv_uid_t uid, uv_gid_t gid, uv_fs_cb cb);//更改所有者及所属组，等同于lchown()
```

### 1.2.5. 文件链接相关

```C++
int uv_fs_link(uv_loop_t* loop, uv_fs_t* req, const char* path, const char* new_path, uv_fs_cb cb);//创建硬链接，等同于link
int uv_fs_symlink(uv_loop_t* loop,uv_fs_t* req,const char* path,const char* new_path,int flags,uv_fs_cb cb);//创建符号链接,等同于symlink。
int uv_fs_readlink(uv_loop_t* loop, uv_fs_t* req, const char* path, uv_fs_cb cb);//读取链接，结果放在req->ptr.
//相对路径转换为绝对路径。同linux 的realpath(3). Windows 用 GetFinalPathNameByHandle. 结果放在req->ptr.
int uv_fs_realpath(uv_loop_t* loop, uv_fs_t* req, const char* path, uv_fs_cb cb);
```

### 1.2.6. 文件监听相关

FS Event在有些系统(AIX、z/OS)使用有限制，而pool不受系统影响。

```C++
int uv_fs_event_init(uv_loop_t* loop, uv_fs_event_t* handle);//初始化监听句柄
int uv_fs_event_start(uv_fs_event_t* handle,uv_fs_event_cb cb, const char* path, unsigned int flags);//开始监听
int uv_fs_event_stop(uv_fs_event_t* handle);//停止监听
//通过监听句柄，获取监听路径。buffer和size是输出参数，buffer需要提前申请好。
int uv_fs_event_getpath(uv_fs_event_t* handle, char* buffer, size_t* size);

int uv_fs_poll_init(uv_loop_t* loop, uv_fs_poll_t* handle);//参考event
int uv_fs_poll_start(uv_fs_poll_t* handle, uv_fs_poll_cb poll_cb, const char* path, unsigned int interval);//参考event
int uv_fs_poll_stop(uv_fs_poll_t* handle);//参考event
int uv_fs_poll_getpath(uv_fs_poll_t* handle, char* buffer, size_t* size);//参考event
```

## Demo

详见

## 1.3. 参考资料

1. [Filesystem](http://docs.libuv.org/en/v1.x/guide/filesystem.html)
2. [API-FS](http://docs.libuv.org/en/v1.x/fs.html)
3. [API-FSevent](http://docs.libuv.org/en/v1.x/fs_event.html)
4. [API-FSpoll](http://docs.libuv.org/en/v1.x/fs_poll.html)
