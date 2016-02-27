#ifndef __BASESOCKET_H__
#define __BASESOCKET_H__
#include "ostype.h"


class BaseSocket
{
	public:
		BaseSocket(int type);
		~BaseSocket();
		
		int Listen(
			const char*	server_ip, 
			int	port);
		
		int GetSocketfd() { return m_socket_fd; };
		int GetSocketType() { return m_socket_type; };
		int AcceptConnection();
		virtual void Read();
		virtual void Write();
		void CloseSocket();
	protected:
		int m_socket_type;
		string m_socket_ip;
		int m_socket_port;
		int m_socket_fd;
};

BaseSocket* FindBaseSocket(int fd);
void AddBaseSocket(BaseSocket* pSocket);
void RemoveBaseSocket(BaseSocket* pSocket);


#endif