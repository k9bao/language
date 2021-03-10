import socket
import time
import datetime

def SendTcp(ip,port):
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    print("begin connect",ip,port)
    s.connect((ip,port))
    # 1118242,153
    data1 = bytes(1118242)
    data2 = bytes(153)
    count = 0
    while True:
        sc = 0
        # if count % 4 == 0:
        #     sc = s.send(data2)
        # else:
        #     sc = s.send(data1)
        # time.sleep(60/1000)
        sc = s.send(data1)
        count = count+1
        print("send ok",datetime.datetime.now(),count,sc)
    s.close()

if __name__ == "__main__":
    SendTcp("10.156.10.150",2362)