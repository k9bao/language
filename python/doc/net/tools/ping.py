# coding=utf-8

import os,time

def pingIP(IPbegin,IPend):
    start_Time=int(time.time())
    IPhost = []
    IP1 =  IPbegin.split('.')[0]
    IP2 =  IPbegin.split('.')[1]
    IP3 =  IPbegin.split('.')[2]
    IP4 = IPbegin.split('.')[-1]
    IPend_last = IPend.split('.')[-1]
    ok_list = []
    fail_list = []
    for i in range(int(IP4)-1,int(IPend_last)):
        ip = str(IP1+'.'+IP2+'.'+IP3+'.'+IP4)
        int_IP4 = int(IP4)
        int_IP4 += 1
        IP4 = str(int_IP4)
        return1=os.system('ping -n 1 -w 1 %s'%ip)
        if return1:
            print('ping %s is fail'%ip)
            fail_list.append(ip)
        else:
            print('ping %s is ok'%ip)
            ok_list.append(ip)

    end_Time = int(time.time())
    print("time(秒)：",end_Time - start_Time,"s")
    print("ping通的ip数：",len(ok_list),"   ping不通的ip数：",len(fail_list))

    ip_True = open('ip_True.txt','w+')
    ip_False = open('ip_False.txt','w+')
    for ip in ok_list:
        print(ip)
        ip_True.write(ip+'\n')

    for ip in fail_list:
        ip_False.write(ip+'\n')

    ip_True.close()
    ip_False.close()

if __name__ == "__main__":
    IPbegin = "10.169.228.72" 
    IPend = "10.169.228.100"
    pingIP(IPbegin,IPend)