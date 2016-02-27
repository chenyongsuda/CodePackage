#include "BaseConn.h"
#define BUFF_SIZE 1024

BaseConn::BaseConn(int fd)
: BaseSocket(SOCKET_TYPE_CON)
{
	m_socket_fd = fd;
}

BaseConn::~BaseConn()
{
	
}

void 
BaseConn::Read()
{
	printf("BaseConn=================>fd:%d\n",m_socket_fd);
	char buf[BUFF_SIZE];
	bzero(buf, BUFF_SIZE);
	int n = 0;
	int nread = 0;
	while ((nread = read(m_socket_fd, buf + n, BUFF_SIZE-1)) > 0) {
		n += nread;
	}
	if (nread == -1 && errno != EAGAIN) {
		
	}
	
	printf("Client:%d 	%s\n",m_socket_fd,buf);
	
}

void 
BaseConn::Write()
{
	
}
