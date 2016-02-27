#ifndef __BASE_CONN_H__
#define __BASE_CONN_H__
#include "BaseSocket.h"

class BaseConn : public BaseSocket
{
	public:
		BaseConn(int fd);
		~BaseConn();
		virtual void Read();
		virtual void Write();
};

#endif