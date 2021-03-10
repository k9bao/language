# http包

## httpClient长连接

[httpClient结构](http://docscn.studygolang.com/pkg/go/build/)

```golang
type Client struct {
    Transport RoundTripper//连接池，client处理公用此连接池，设置长连接等参数
    CheckRedirect func(req *Request, via []*Request) error
    Jar CookieJar//cook设置
    Timeout time.Duration // Go 1.3 设置超时时间
}
```

[RoundTripper](http://docscn.studygolang.com/pkg/net/http/#Transport)

```golang
type RoundTripper interface {
    RoundTrip(*Request) (*Response, error)
}
//Transport implement RoundTripper interface
type Transport struct {
    Proxy func(*Request) (*url.URL, error)
    DialContext func(ctx context.Context, network, addr string) (net.Conn, error) // Go 1.7
    Dial func(network, addr string) (net.Conn, error)
    DialTLSContext func(ctx context.Context, network, addr string) (net.Conn, error) // Go 1.14
    DialTLS func(network, addr string) (net.Conn, error) // Go 1.4
    TLSClientConfig *tls.Config
    TLSHandshakeTimeout time.Duration // Go 1.3 限制 TLS握手的时间
    DisableKeepAlives bool //是否开启http keepalive功能
    DisableCompression bool
    MaxIdleConns int // Go 1.7 所有host的连接池最大链接数量,keep-alive长连接
    MaxIdleConnsPerHost int //每个host的最大链接数量<=MaxIdleConns,keep-alive长连接,默认是 DefaultMaxIdleConnsPerHost=2
    MaxConnsPerHost int // Go 1.11 每个host可以发出的最大连接个数（包括keep-alive长链接和短链接）
    IdleConnTimeout time.Duration // Go 1.7 socket在该时间内没有交互则自动关闭连接
    ResponseHeaderTimeout time.Duration // Go 1.1 限制读取response header的时间,默认 timeout + 5*time.Second
    ExpectContinueTimeout time.Duration // Go 1.6 限制client在发送包含 Expect: 100-continue的header到收到继续发送body的response之间的时间等待。
    TLSNextProto map[string]func(authority string, c *tls.Conn) RoundTripper // Go 1.6
    ProxyConnectHeader Header // Go 1.8
    MaxResponseHeaderBytes int64 // Go 1.7
    WriteBufferSize int // Go 1.13
    ReadBufferSize int // Go 1.13
    ForceAttemptHTTP2 bool // Go 1.13
}
```

举例

```golang
package main

import (
    "bytes"
    "io/ioutil"
    "log"
    "net"
    "net/http"
    "time"
)

var (
    httpClient *http.Client
)

// init HTTPClient
func init() {
    httpClient = createHTTPClient()
}

const (
    MaxIdleConns int = 100
    MaxIdleConnsPerHost int = 100
    IdleConnTimeout int = 90
)

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
    client := &http.Client{
        Transport: &http.Transport{
            Proxy: http.ProxyFromEnvironment,
            DialContext: (&net.Dialer{
                Timeout:   30 * time.Second,
                KeepAlive: 30 * time.Second,
            }).DialContext,
            MaxIdleConns:        MaxIdleConns,
            MaxIdleConnsPerHost: MaxIdleConnsPerHost,
            IdleConnTimeout:     time.Duration(IdleConnTimeout)* time.Second,
        },
    Timeout: 20 * time.Second,
    }
    return client
}

func main() {
    var endPoint string = "https://localhost:8080/doSomething"

    req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer([]byte("Post this data")))
    if err != nil {
        log.Fatalf("Error Occured. %+v", err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    // use httpClient to send request
    response, err := httpClient.Do(req)
    if err != nil && response == nil {
        log.Fatalf("Error sending request to API endpoint. %+v", err)
    } else {
        // Close the connection to reuse it
        defer response.Body.Close()

        // Let's check if the work actually is done
        // We have seen inconsistencies even when we get 200 OK response
        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatalf("Couldn't parse response body. %+v", err)
        }

        log.Println("Response Body:", string(body))
    }

}
```

## 参考资料

1. [go-http](https://studygolang.com/articles/12040)
