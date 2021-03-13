#include <winsock2.h>
#include <windows.h>
#include <string>
#include <iostream>
using namespace std;

#pragma comment(lib,"ws2_32.lib")
#pragma comment(lib,"kernel32.lib")

HANDLE g_hIOCP;

enum IO_OPERATION{IO_READ,IO_WRITE};

struct IO_DATA{
    OVERLAPPED                  Overlapped;
    WSABUF                      wsabuf;
    int                         nBytes;
    IO_OPERATION                opCode;
    SOCKET                      client;
};

char buffer[1024];

//工作线程
DWORD WINAPI WorkerThread (LPVOID WorkThreadContext) {
    IO_DATA *lpIOContext = NULL; 
    DWORD nBytes = 0;
    DWORD dwFlags = 0; 
    int nRet = 0;

    DWORD dwIoSize = 0; 
    void * lpCompletionKey = NULL;
    LPOVERLAPPED lpOverlapped = NULL;

    while(1){
        GetQueuedCompletionStatus(g_hIOCP, &dwIoSize,(PULONG_PTR)&lpCompletionKey,(LPOVERLAPPED *)&lpOverlapped, INFINITE);

        lpIOContext = (IO_DATA *)lpOverlapped;
        if(dwIoSize == 0)
        {
            cout << "Client disconnect" << endl;
            closesocket(lpIOContext->client);
            delete lpIOContext;
            continue;
        }

        if(lpIOContext->opCode == IO_READ) {// a read operation complete
            ZeroMemory(&lpIOContext->Overlapped, sizeof(lpIOContext->Overlapped));
            lpIOContext->wsabuf.buf = buffer;
            lpIOContext->wsabuf.len = strlen(buffer)+1;
            lpIOContext->opCode = IO_WRITE;
            lpIOContext->nBytes = strlen(buffer)+1;
            dwFlags = 0;
            nBytes = strlen(buffer)+1;
            nRet = WSASend(
                lpIOContext->client,
                &lpIOContext->wsabuf, 1, &nBytes,
                dwFlags,
                &(lpIOContext->Overlapped), NULL);
            if( nRet == SOCKET_ERROR && (ERROR_IO_PENDING != WSAGetLastError()) ) {
                cout << "WASSend Failed::Reason Code::"<< WSAGetLastError() << endl;
                closesocket(lpIOContext->client);
                delete lpIOContext;
                continue;
            }
            memset(buffer, NULL, sizeof(buffer));
        }
        else if(lpIOContext->opCode == IO_WRITE) {//a write operation complete
            // Write operation completed, so post Read operation.
            lpIOContext->opCode = IO_READ; 
            nBytes = 1024;
            dwFlags = 0;
            lpIOContext->wsabuf.buf = buffer;
            lpIOContext->wsabuf.len = nBytes;
            lpIOContext->nBytes = nBytes;
            ZeroMemory(&lpIOContext->Overlapped, sizeof(lpIOContext->Overlapped));

            nRet = WSARecv(
                lpIOContext->client,
                &lpIOContext->wsabuf, 1, &nBytes,
                &dwFlags,
                &lpIOContext->Overlapped, NULL);
            if( nRet == SOCKET_ERROR && (ERROR_IO_PENDING != WSAGetLastError()) ) {
                cout << "WASRecv Failed::Reason Code1::"<< WSAGetLastError() << endl;
                closesocket(lpIOContext->client);
                delete lpIOContext;
                continue;
            } 
            cout<<lpIOContext->wsabuf.buf<<endl;
        }
    }
    return 0;
}
void main ()
{
    WSADATA wsaData;
    WSAStartup(MAKEWORD(2,2), &wsaData);

    SOCKET    m_socket = WSASocket(AF_INET,SOCK_STREAM, IPPROTO_TCP, NULL,0,WSA_FLAG_OVERLAPPED);

    sockaddr_in server;
    server.sin_family = AF_INET;
    server.sin_port = htons(6000);
    server.sin_addr.S_un.S_addr = htonl(INADDR_ANY);

    bind(m_socket ,(sockaddr*)&server,sizeof(server));

    listen(m_socket, 8);

    SYSTEM_INFO sysInfo;
    GetSystemInfo(&sysInfo);
    int g_ThreadCount = sysInfo.dwNumberOfProcessors * 2;

    //创建iocp句柄
    g_hIOCP = CreateIoCompletionPort(INVALID_HANDLE_VALUE,NULL,0,g_ThreadCount);
    //CreateIoCompletionPort((HANDLE)m_socket,g_hIOCP,0,0);

    for( int i=0;i < g_ThreadCount; ++i){
        HANDLE  hThread;
        DWORD   dwThreadId;
        hThread = CreateThread(NULL, 0, WorkerThread, 0, 0, &dwThreadId);
        CloseHandle(hThread);
    }

    while(1)
    {
        SOCKET client = accept( m_socket, NULL, NULL );
        cout << "Client connected." << endl;
        //socket 与 iocp 句柄关联
        if (CreateIoCompletionPort((HANDLE)client, g_hIOCP, 0, 0) == NULL){
            cout << "Binding Client Socket to IO Completion Port Failed::Reason Code::"<< GetLastError() << endl;
            closesocket(client);
        }else { //post a recv request
            IO_DATA * data = new IO_DATA;
            memset(buffer, NULL ,1024);
            memset(&data->Overlapped, 0 , sizeof(data->Overlapped));
            data->opCode = IO_READ;
            data->nBytes = 0;
            data->wsabuf.buf  = buffer;
            data->wsabuf.len  = sizeof(buffer);
            data->client = client;
            DWORD nBytes= 1024 ,dwFlags=0;
            int nRet = WSARecv(client,&data->wsabuf, 1, &nBytes,
                &dwFlags, &data->Overlapped, NULL);
            if(nRet == SOCKET_ERROR  && (ERROR_IO_PENDING != WSAGetLastError())){
                cout << "WASRecv Failed::Reason Code::"<< WSAGetLastError() << endl;
                closesocket(client);
                delete data;
            }
            cout<<data->wsabuf.buf<<endl;
        }
    }
    closesocket(m_socket);
    WSACleanup();
}