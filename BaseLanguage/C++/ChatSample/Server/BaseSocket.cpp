#include "BaseSocket.h"
#include "EventDispatch.h"

typedef hash_map<int,BaseSocket*> SocketMap;
SocketMap	g_socket_map;

void AddBaseSocket(BaseSocket* pSocket)
{
	g_socket_map.insert(make_pair(pSocket->GetSocketfd(), pSocket));
}

void RemoveBaseSocket(BaseSocket* pSocket)
{
	g_socket_map.erase(pSocket->GetSocketfd());
}

BaseSocket* FindBaseSocket(int fd)
{
	BaseSocket* pSocket = NULL;
	SocketMap::iterator iter = g_socket_map.find(fd);
	if (iter != g_socket_map.end())
	{
		pSocket = iter->second;
	}

	return pSocket;
}

BaseSocket::BaseSocket(int type)
{
	m_socket_type = type;
}

BaseSocket::~BaseSocket()
{
	
}

void 
BaseSocket::Read()
{
	
}


void 
BaseSocket::Write()
{
	
}

int 
BaseSocket::AcceptConnection()
{
	//m_socket_fd
	int listenfd = m_socket_fd;
	//Accept a client we need to General a subclass
	int con_fd = 0;
	sockaddr_in peer_addr;
	socklen_t addr_len = sizeof(sockaddr_in);
	if((con_fd = accept(listenfd, (sockaddr*)&peer_addr, &addr_len)) == SOCKET_ERROR)
	{
		//printf("BaseSocket=================>fd: %d accept failed\n",con_fd);
		return SOCKET_ERROR;
	}
	
	//Set non block
	//int ret =setnonblocking(socketfd);
	int ret = 0;
	ret = fcntl(con_fd, F_SETFL, fcntl(con_fd, F_GETFD, 0)|O_NONBLOCK);
	if(SOCKET_ERROR == ret)
	{
		printf("BaseSocket=================>fd: %d set nonblock failed\n",con_fd);
		close(con_fd);
		return SOCKET_ERROR;
	}
	
	//Set reuse
	int reuse = 1;
	ret = setsockopt(con_fd, SOL_SOCKET, SO_REUSEADDR, (char*)&reuse, sizeof(reuse));
	if(SOCKET_ERROR == ret)
	{
		printf("BaseSocket=================>fd: %d setsockopt failed\n",con_fd);
		close(con_fd);
		return SOCKET_ERROR;
	}
	
	//add to event loop
	//Add Event
	//EventDispatch::GetInstance()->AddEvent(con_fd,(EPOLLIN|EPOLLOUT|EPOLLET|EPOLLPRI|EPOLLERR|EPOLLHUP));
	return con_fd;
}

int
BaseSocket::Listen(
		const char*	server_ip, 
		int	port)
{
	//What do when new 1.new socket 2. add to socket list 3.add to Event system 4.delete the container
	//Set the IP
	m_socket_ip = server_ip;
	m_socket_port = port;
	//Do listen
	int socketfd = socket(AF_INET, SOCK_STREAM, 0);
	
	//Set non block
	//int ret =setnonblocking(socketfd);
	int ret = 0;
	ret = fcntl(socketfd, F_SETFL, fcntl(socketfd, F_GETFD, 0)|O_NONBLOCK);
	if(SOCKET_ERROR == ret)
	{
		printf("BaseSocket=================>ip:%s,port:%d setnonblocking failed\n",server_ip,port);
		close(socketfd);
		return SOCKET_ERROR;
	}
	
	//Set reuse
	int reuse = 1;
	ret = setsockopt(socketfd, SOL_SOCKET, SO_REUSEADDR, (char*)&reuse, sizeof(reuse));
	if(SOCKET_ERROR == ret)
	{
		printf("BaseSocket=================>ip:%s,port:%d setsockopt failed\n",server_ip,port);
		close(socketfd);
		return SOCKET_ERROR;
	}
	
	//New socket
	sockaddr_in serv_addr;
	memset(&serv_addr, 0, sizeof(sockaddr_in));
	serv_addr.sin_family = AF_INET;
	serv_addr.sin_port = htons(port);
	serv_addr.sin_addr.s_addr = inet_addr(server_ip);
	
	//bind
	ret = bind(socketfd, (sockaddr*)&serv_addr, sizeof(serv_addr));
	if(SOCKET_ERROR == ret)
	{
		printf("BaseSocket=================>ip:%s,port:%d bind failed\n",server_ip,port);
		close(socketfd);
		return SOCKET_ERROR;
	}
	
	//listern
	ret = listen(socketfd, 64);
	if(SOCKET_ERROR == ret)
	{
		printf("BaseSocket=================>ip:%s,port:%d listen failed\n",server_ip,port);
		close(socketfd);
		return SOCKET_ERROR;
	}
	
	m_socket_fd = socketfd;
	printf("BaseSocket=================>ip:%s,port:%d listen success\n",server_ip,port);
	
	//Add To The g_map
	AddBaseSocket(this);
	//Add Event
	EventDispatch::GetInstance()->AddEvent(socketfd,(EPOLLIN|EPOLLET|EPOLLPRI|EPOLLERR|EPOLLHUP));

	return socketfd;
}

void 
BaseSocket::CloseSocket()
{
	//1. close the fd
	close(m_socket_fd);
	//2. remove from the map
	RemoveBaseSocket(this);
	//3. remove the event
	EventDispatch::GetInstance()->RemoveEvent(m_socket_fd);
	//4. delete this TODO
	printf("BaseSocket=================>fd:%d cloased\n",m_socket_fd);
}