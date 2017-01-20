#ifndef EVENTMUTIIOEPOLL_H
#define EVENTMUTIIOEPOLL_H

#include "EventMutiIOAPI.h"
#include "Event.h"
#include <sys/epoll.h>

class EventMutiIOEpoll : public EventMutiIOAPI
{
    public:
        EventMutiIOEpoll();
        virtual ~EventMutiIOEpoll();

        virtual int init(void);
        virtual int registe_event();
        virtual int unregiste_event();
        virtual int update_event();
    protected:

    private:
};

#endif // EVENTMUTIIOEPOLL_H
