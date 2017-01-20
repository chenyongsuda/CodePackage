#include "TCPServer.h"
#include "ErrorNoDefine.h"
#include "EventMutiIOEpoll.h"

TCPServer::TCPServer()
:listen_(),eventCenter_(),eventIO(NULL)
{

}

TCPServer::~TCPServer()
{
    //dtor
    if (NULL != eventIO){
        delete eventIO;
    }
}

int
TCPServer::init(const char* addr, const char* port){
    int ret = TC_SUCCESS;
    //1. init listen
    if (TC_SUCCESS != (ret = listen_.init(addr,port))){
        return ret;
    }
    //2. init event center
    eventIO = new EventMutiIOEpoll(&eventCenter_);
    if(TC_SUCCESS != (ret = eventCenter_.init(eventIO))){
        return ret;
    }

    return ret;
}

int
TCPServer::start(){
    int ret = TC_SUCCESS;
    //Start listen
    if (TC_SUCCESS != (ret = listen_.start())){
        return ret;
    }
    //add event to event center
    if(TC_SUCCESS != (ret = eventCenter_.registe_event(listen_.getListenEvent()))){
        return ret;
    }

    //Go loop
    if(TC_SUCCESS != (ret = eventCenter_.loop())){
        return ret;
    }
    return ret;
}
