#ifndef TCPSERVER_H
#define TCPSERVER_H

#include "TCPListener.h"
#include "EventCenter.h"

class TCPServer
{
    public:
        TCPServer();
        virtual ~TCPServer();
        int init(const char* addr, const char* port);
        int start();
    protected:

    private:
        TCPListener listen_;
        EventCenter eventCenter_;
        EventMutiIOAPI* eventIO;
};

#endif // TCPSERVER_H
