#ifndef __NETLIB_H__
#define __NETLIB_H__

#include "EventDispatch.h"

int netlib_listen(
	const char*	server_ip, 
	int	port
	);

#endif