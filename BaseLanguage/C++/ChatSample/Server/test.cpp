#include "netlib.h"
#include "ostype.h"

int main(int argc, char *argv[])
{
	//listern port
	netlib_listen("127.0.0.1", 10000);
	//Start event loop
	EventDispatch::GetInstance()->StartEventLoop();
}