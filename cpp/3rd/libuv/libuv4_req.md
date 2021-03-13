# request handle

## 1.1. request对象

```C++
#define UV_REQ_FIELDS                                                         \
  /* public */                                                                \
  void* data;                                                                 \
  /* read-only */                                                             \
  uv_req_type type;                                                           \
  /* private */                                                               \
  void* reserved[6];                                                          \
  UV_REQ_PRIVATE_FIELDS                                                       \

/* Abstract base class of all requests. */
struct uv_req_s {
  UV_REQ_FIELDS
};
```
