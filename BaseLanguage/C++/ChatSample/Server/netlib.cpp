#include "netlib.h"
#include "ostype.h"
#include "BaseSocket.h"

int netlib_listen(
		const char*	server_ip, 
		int	port)
{
	BaseSocket* bs = new BaseSocket(SOCKET_TYPE_LISTEN);
	bs->Listen(server_ip,port);
}