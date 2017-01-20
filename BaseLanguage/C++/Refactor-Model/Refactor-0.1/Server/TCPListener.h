#ifndef TCPLISTENER_H
#define TCPLISTENER_H
#include <functional>
#include "Event.h"

class TCPListener
{
    public:
        TCPListener();
        virtual ~TCPListener();
        int init(const char* addr, const char* port);
        int start();
        Event* getListenEvent();
    protected:

    private:
        int accept(int fd, int type, void *args);
        const char*  addr_;
        const char*  port_;
        int fd_;
        Event listenEvent_;
};

#endif // TCPLISTENER_H
