#include "EventDispatch.h"
#include "BaseSocket.h"
#include "BaseConn.h"
#include "ostype.h"
#define MAX_SOCKETS_LEN 1024
#define MAX_WAIT_TIME 500

EventDispatch* EventDispatch::instance = NULL;

void 
EventDispatch::AddEvent(int fd, int event_type)
{
	struct epoll_event ev;
	ev.events = event_type;
	ev.data.fd = fd;
	if (epoll_ctl(m_epfd, EPOLL_CTL_ADD, fd, &ev) != 0)
	{
		printf("EventDispatch=================>fd:%d AddEvent Failed\n",fd);
	}
	else
	{
	}
}


void 
EventDispatch::RemoveEvent(int fd)
{
	if (epoll_ctl(m_epfd, EPOLL_CTL_DEL, fd, NULL) != 0)
	{
		printf("EventDispatch=================>fd:%d DelEvent %d\n",fd);
	}
}


void 
EventDispatch::AddTimer()
{
	
}


void 
EventDispatch::RemoveTimer()
{
	
}

EventDispatch*
EventDispatch::GetInstance()
{
	if(NULL == instance)
	{
		instance = new EventDispatch();
	}
	return instance;
}

EventDispatch::EventDispatch()
{
	m_epfd = epoll_create(MAX_SOCKETS_LEN);
	if(m_epfd == -1)
	{
		//error
		printf("EventDispatch=================>epoll_create failed\n");
	}
}


EventDispatch::~EventDispatch()
{
	
}
void 
EventDispatch::Init()
{
	
}

void 
EventDispatch::StartEventLoop(int wait_time)
{
	char buf[BUFSIZ];
	printf("EventDispatch=================>StartEventLoop\n");
	struct epoll_event events[MAX_SOCKETS_LEN];
	for ( ; ; ) 
	{
		int count = epoll_wait(m_epfd, events, MAX_SOCKETS_LEN,wait_time);
		
		for(int i = 0; i < count; i++)
		{
			//
			int fd = events[i].data.fd;
			BaseSocket* pFindSocket = FindBaseSocket(fd);
			if(NULL == pFindSocket)
			{
				printf("EventDispatch=================>FIND NULL\n");
				continue;
			}
			
			if(SOCKET_TYPE_LISTEN == pFindSocket->GetSocketType())
			{
				int con_fd = 0;
				//Accept The Connect
				while ( (con_fd = pFindSocket->AcceptConnection()) != SOCKET_ERROR ) 
				{
					BaseSocket* socket = new BaseConn(con_fd);
					AddBaseSocket(socket);
					EventDispatch::GetInstance()->AddEvent(con_fd,(EPOLLIN|EPOLLOUT|EPOLLET|EPOLLPRI|EPOLLERR|EPOLLHUP));
					//Do CallBack Set Call Back Data as  con_fds
					printf("EventDispatch=================>Accept Success new fd:%d\n",con_fd);
				}
			}
			else
			{
				
				if(events[i].events & EPOLLIN)
				{
					u_long avail = 0;
					//the Socket Client closed
					if(ioctl(events[i].data.fd, FIONREAD, &avail) == SOCKET_ERROR || avail == 0)
					{
						pFindSocket->CloseSocket();
					}
					//read the data only
					else
					{
						pFindSocket->Read();
					}
				}
				
				if(events[i].events & EPOLLOUT)
				{
					//Write The data 
					pFindSocket->Write();
					printf("EventDispatch=================>EPOLLOUT FD:%d\n",events[i].data.fd);
				}
				
				if(events[i].events & (EPOLLPRI|EPOLLERR|EPOLLHUP))
				{
					//UNNormal Reset
					//当socket的一端认为对方发来了一个不存在的4元组请求的时候,会回复一个RST响应,在epoll上会响应为EPOLLHUP事件,目前我已知的两种情况会发响应RST
					//[1] 当客户端向一个没有在listen的服务器端口发送的connect的时候 服务器会返回一个RST 因为服务器根本不知道这个4元组的存在 
					//[2] 当已经建立好连接的一对客户端和服务器,客户端突然操作系统崩溃,或者拔掉电源导致操作系统重新启动(kill pid或者正常关机不行的,因为操作系统会发送FIN给对方).这时服务器在原有的4元组上发送数据,会收到客户端返回的RST,因为客户端根本不知道之前这个4元组的存在
					printf("EventDispatch=================>EPOLLClose FD:%d\n",events[i].data.fd);
					pFindSocket->CloseSocket();
				}
			}
			
			
		}
	}
}

/*
//main sample
int main()
{
	struct epoll_event events[MAX_SOCKETS_LEN];
	int epfd = epoll_create(MAX_SOCKETS_LEN);
	for ( ; ; ) 
	{
		int count = epoll_wait(epfd, events, MAX_SOCKETS_LEN,MAX_WAIT_TIME);
		
		for(int i = 0; i < count; i++)
		{
			if(events[i].events & EPOLLIN)
			{
				//IN
				printf("[EventDispatcher]========================= Event in fd:"+events[i].data.fd);
			}
			
			if(events[i].events & EPOLLOUT)
			{
				//OUT
			}
			
			if(events[i].events & (EPOLLPRI|EPOLLERR|EPOLLHUP))
			{
				//CLOSE
			}
			
		}
	}
}*/