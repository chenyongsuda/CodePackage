#ifndef EVENTMUTIIOEPOLL_H
#define EVENTMUTIIOEPOLL_H

#include "EventMutiIOAPI.h"
#include <sys/epoll.h>
#include <fcntl.h>

class EventCenter;
class EventMutiIOEpoll : public EventMutiIOAPI
{
    public:
        EventMutiIOEpoll(EventCenter* evc);
        virtual ~EventMutiIOEpoll();

        virtual int init(void);
        virtual int registe_event(int fd, int type);
        virtual int unregiste_event(int fd, int type);
        virtual int update_event(int fd, int type);
        virtual int loop();
        virtual int destory();
    protected:

    private:
        int epfd_;
        //struct epoll_event *events_;
        int loop_size_;
        EventCenter* evc_;
};

#endif // EVENTMUTIIOEPOLL_H
