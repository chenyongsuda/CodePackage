#ifndef __OS_TYPE_H__
#define __OS_TYPE_H__

#include <sys/epoll.h>


#include <pthread.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/ioctl.h>
#include <netinet/in.h>
#include <netinet/tcp.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <errno.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdint.h>		// define int8_t ...
#include <signal.h>
#include <unistd.h>



//UTILS
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>

//Define Map
#ifdef __GNUC__
    #include <ext/hash_map>
    using namespace __gnu_cxx;
    namespace __gnu_cxx {
        template<> struct hash<std::string> {
            size_t operator()(const std::string& x) const {
                return hash<const char*>()(x.c_str());
            }
        };
    }
#else
    #include <hash_map>
    using namespace stdext;
#endif

using namespace std;

//set option and connect result
const int SOCKET_ERROR	= -1;

//There have 2 type the listen and the con
enum
{
	SOCKET_TYPE_LISTEN = 0X0001,
	SOCKET_TYPE_CON = 0X0002,
};

#endif