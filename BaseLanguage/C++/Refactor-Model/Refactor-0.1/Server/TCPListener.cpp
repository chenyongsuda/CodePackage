#include "TCPListener.h"
#include "ErrorNoDefine.h"
#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <pthread.h>
#include <stdlib.h>
#include <fcntl.h>

using namespace std::placeholders;

TCPListener::TCPListener()
:addr_(NULL),port_(NULL),fd_(0),listenEvent_()
{
    //ctor
}

TCPListener::~TCPListener()
{
    //dtor
}

int
TCPListener::init(const char* addr, const char* port){
    int ret = TC_SUCCESS;
    addr_ = addr;
    port_ = port;
    return ret;
}

int
TCPListener::start(){
    int ret = TC_SUCCESS;

    //Do listen
	int socketfd = socket(AF_INET, SOCK_STREAM, 0);
	if(TC_SOCKET_ERROR == socketfd){
        return TC_SOCKET_ERROR;
	}

    //set no block
	ret = fcntl(socketfd, F_SETFL, fcntl(socketfd, F_GETFD, 0)|O_NONBLOCK);
	if(TC_SOCKET_ERROR == ret){
        close(socketfd);
        return ret;
	}

    //Set reuse
    int reuse = 1;
    ret = setsockopt(socketfd, SOL_SOCKET, SO_REUSEADDR, (char*)&reuse, sizeof(reuse));
	if(TC_SOCKET_ERROR == ret)
	{
		close(socketfd);
		return ret;
	}

	//New socket
	sockaddr_in serv_addr;
	memset(&serv_addr, 0, sizeof(sockaddr_in));
	serv_addr.sin_family = AF_INET;
	serv_addr.sin_port = htons(atoi(port_));
	serv_addr.sin_addr.s_addr = inet_addr(addr_);

	//bind
	ret = ::bind(socketfd, (sockaddr*)&serv_addr, sizeof(serv_addr));
	if(TC_SOCKET_ERROR == ret)
	{
		close(socketfd);
		return ret;
	}

	//listern
	ret = listen(socketfd, 64);
	if(TC_SOCKET_ERROR == ret)
	{
		close(socketfd);
		return ret;
	}

    //Set FD
	fd_ = socketfd;

    //Set Event
    listenEvent_.fd = fd_;
    listenEvent_.type = Event::kEventRead;
    listenEvent_.read_cb = std::bind(&TCPListener::accept,this,_1,_2,_3);

    return ret;
}

Event*
TCPListener::getListenEvent(){
    return &listenEvent_;
}

int
TCPListener::accept(int fd, int type, void *args){
    //Log
    int ret = TC_SUCCESS;
//Accept a client we need to General a subclass
	sockaddr_in peer_addr;
	socklen_t addr_len = sizeof(sockaddr_in);
	if((ret = ::accept(fd_, (sockaddr*)&peer_addr, &addr_len)) == TC_SOCKET_ERROR)
	{
	    printf("TCPListener::accept error[%d] fd[%d] type[%d]\n",ret,fd,type);
		return ret;
	}
    printf("TCPListener::accept fd[%d] type[%d]\n",fd,type);
    return ret;
}


