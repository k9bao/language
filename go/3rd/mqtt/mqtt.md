# 1. paho.mqtt.golang

- [1. paho.mqtt.golang](#1-pahomqttgolang)
  - [1.1. package](#11-package)
    - [1.1.1. 常量及全局函数](#111-常量及全局函数)
    - [1.1.2. ControlPacket](#112-controlpacket)
      - [1.1.2.1. FixHeader&Details](#1121-fixheaderdetails)
      - [1.1.2.2. Connect](#1122-connect)
        - [1.1.2.2.1. ConnectPacket](#11221-connectpacket)
        - [1.1.2.2.2. ConnackPacket](#11222-connackpacket)
        - [1.1.2.2.3. DisconnectPacket](#11223-disconnectpacket)
      - [1.1.2.3. pub](#1123-pub)
        - [1.1.2.3.1. PublishPacket](#11231-publishpacket)
        - [1.1.2.3.2. PubackPacket](#11232-pubackpacket)
        - [1.1.2.3.3. PubrecPacket](#11233-pubrecpacket)
        - [1.1.2.3.4. PubrelPacket](#11234-pubrelpacket)
        - [1.1.2.3.5. PubcompPacket](#11235-pubcomppacket)
      - [1.1.2.4. sub](#1124-sub)
        - [1.1.2.4.1. SubscribePacket](#11241-subscribepacket)
        - [1.1.2.4.2. SubackPacket](#11242-subackpacket)
        - [1.1.2.4.3. UnsubscribePacket](#11243-unsubscribepacket)
        - [1.1.2.4.4. UnsubackPacket](#11244-unsubackpacket)
      - [1.1.2.5. ping](#1125-ping)
        - [1.1.2.5.1. PingreqPacket](#11251-pingreqpacket)
        - [1.1.2.5.2. PingrespPacket](#11252-pingresppacket)
  - [1.2. Token interface](#12-token-interface)
    - [1.2.1. ConnectToken](#121-connecttoken)
    - [1.2.2. DisconnectToken](#122-disconnecttoken)
    - [1.2.3. PublishToken](#123-publishtoken)
    - [1.2.4. SubscribeToken](#124-subscribetoken)
    - [1.2.5. UnsubscribeToken](#125-unsubscribetoken)
    - [1.2.6. PacketAndToken](#126-packetandtoken)
  - [1.3. Store](#13-store)
    - [1.3.1. MemoryStore](#131-memorystore)
    - [1.3.2. FileStore](#132-filestore)
  - [1.4. Client](#14-client)
    - [1.4.1. ClientOptions](#141-clientoptions)
    - [1.4.2. ClientOptionsReader](#142-clientoptionsreader)
    - [1.4.3. Client interface](#143-client-interface)
  - [1.5. 参考资料](#15-参考资料)

## 1.1. package

### 1.1.1. 常量及全局函数

```golang
type component string
const (
    NET component = "[net]     "
    PNG component = "[pinger]  "
    CLI component = "[client]  "
    DEC component = "[decode]  "
    MES component = "[message] "
    STR component = "[store]   "
    MID component = "[msgids]  "
    TST component = "[test]    "
    STA component = "[state]   "
    ERR component = "[error]   "
)
func DefaultConnectionLostHandler(client Client, reason error)//链接丢失，仅打印日志
```

### 1.1.2. ControlPacket

#### 1.1.2.1. FixHeader&Details

#### 1.1.2.2. Connect

##### 1.1.2.2.1. ConnectPacket

##### 1.1.2.2.2. ConnackPacket

##### 1.1.2.2.3. DisconnectPacket

#### 1.1.2.3. pub

##### 1.1.2.3.1. PublishPacket

##### 1.1.2.3.2. PubackPacket

##### 1.1.2.3.3. PubrecPacket

##### 1.1.2.3.4. PubrelPacket

##### 1.1.2.3.5. PubcompPacket

#### 1.1.2.4. sub

##### 1.1.2.4.1. SubscribePacket

##### 1.1.2.4.2. SubackPacket

##### 1.1.2.4.3. UnsubscribePacket

##### 1.1.2.4.4. UnsubackPacket

#### 1.1.2.5. ping

##### 1.1.2.5.1. PingreqPacket

##### 1.1.2.5.2. PingrespPacket

## 1.2. Token interface

任务执行返回接口，比如 Connect,DisConnect,Pub,sub 等的返回值

```golang
type Token interface {
    Wait() bool
    WaitTimeout(time.Duration) bool
    Error() error
}
```

### 1.2.1. ConnectToken

### 1.2.2. DisconnectToken

### 1.2.3. PublishToken

### 1.2.4. SubscribeToken

### 1.2.5. UnsubscribeToken

### 1.2.6. PacketAndToken

## 1.3. Store

### 1.3.1. MemoryStore

### 1.3.2. FileStore

## 1.4. Client

### 1.4.1. ClientOptions

```golang
func NewClientOptions() *ClientOptions
```

### 1.4.2. ClientOptionsReader

ClientOptions的接口实现

### 1.4.3. Client interface

```golang
func NewClient(o *ClientOptions) Client
```

## 1.5. 参考资料

1. [github](https://github.com/eclipse/paho.mqtt.golang)
2. [自家博客](https://github.com/knowledgebao/knowledgebao.github.io/_posts/net/protocol/mqtt/mqtt.md)
3. [API](https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang@v1.2.0)
