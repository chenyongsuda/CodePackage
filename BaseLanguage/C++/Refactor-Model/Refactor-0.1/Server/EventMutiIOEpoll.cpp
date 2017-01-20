#include "EventMutiIOEpoll.h"
#include "ErrorNoDefine.h"
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <sys/socket.h>
#include <netdb.h>
#include <fcntl.h>
#include <string.h>
#include "EventCenter.h"

EventMutiIOEpoll::EventMutiIOEpoll(EventCenter* evc)
:epfd_(0),loop_size_(0),evc_(evc)
{
    //ctor
}

EventMutiIOEpoll::~EventMutiIOEpoll()
{
    destory();
}


int
EventMutiIOEpoll::init(void){
    int ret = TC_SUCCESS;
    epfd_ = epoll_create(EPOLL_CREATE_SIZE);
    if(-1 == epfd_){
        ret = TC_EPOLL_ERR;
    }
    return ret;
}

int
EventMutiIOEpoll::registe_event(int fd, int type){
    int ret = TC_SUCCESS;

    struct epoll_event event;
    event.events = 0;
    event.data.fd = fd;
    if (type & Event::kEventRead) {
        event.events |= EPOLLIN;
    }
    if (type & Event::kEventWrite) {
        event.events |= EPOLLOUT;
    }
    int epoll_ret = epoll_ctl(epfd_,EPOLL_CTL_ADD,fd,&event);
    if(-1 == epoll_ret){
        ret = TC_EPOLL_ERR;
    }
    return ret;
}

int
EventMutiIOEpoll::unregiste_event(int fd, int type){
    int ret = TC_SUCCESS;

    struct epoll_event event;
    event.events = 0;
    event.data.fd = fd;
    if (type & Event::kEventRead) {
        event.events |= EPOLLIN;
    }
    if (type & Event::kEventWrite) {
        event.events |= EPOLLOUT;
    }
    int epoll_ret = epoll_ctl(epfd_,EPOLL_CTL_DEL,fd,&event);
    if(-1 == epoll_ret){
        ret = TC_EPOLL_ERR;
    }
    return ret;
}

int
EventMutiIOEpoll::update_event(int fd, int type){
    int ret = TC_SUCCESS;

    struct epoll_event event;
    event.events = 0;
    event.data.fd = fd;
    if (type & Event::kEventRead) {
        event.events |= EPOLLIN;
    }
    if (type & Event::kEventWrite) {
        event.events |= EPOLLOUT;
    }
    int epoll_ret = epoll_ctl(epfd_,EPOLL_CTL_MOD,fd,&event);
    if(-1 == epoll_ret){
        ret = TC_EPOLL_ERR;
    }
    return ret;
}

int
EventMutiIOEpoll::loop(){
    int ret = TC_SUCCESS;
    //0 means no block
    //-1 means block
    int events_num = 0;
    struct epoll_event events[1024];
    events_num = epoll_wait(epfd_,events,1024,0);
    /*if(-1 == events_num){
        ret = TC_EPOLL_ERR;
        return ret;
    }*/

    for(int i = 0; i < events_num; i++){
        printf("LOOP Number %d",events_num);
        struct epoll_event *e = &events[i];
        int event_type = 0;
        if (e->events & EPOLLERR) {
            event_type |= Event::kEventError;
        }
        if (e->events & (EPOLLIN | EPOLLHUP)) {
            event_type |= Event::kEventRead;
        }
        if (e->events & EPOLLOUT) {
            event_type |= Event::kEventWrite;
        }
        //e->data.fd
        evc_->add_ready_event(e->data.fd,event_type);
    }
    return ret;
}

int
EventMutiIOEpoll::destory(){
int ret = TC_SUCCESS;
close(epfd_);
return ret;

}
