#ifndef EVENTMUTIIOAPI_H
#define EVENTMUTIIOAPI_H
#include "Event.h"

class EventMutiIOAPI
{
    public:
        EventMutiIOAPI();
        virtual ~EventMutiIOAPI();
        virtual int init(void) = 0;
        virtual int registe_event(int fd, int type) = 0;
        virtual int unregiste_event(int fd, int type) = 0;
        virtual int update_event(int fd, int type) = 0;
        virtual int loop() = 0;
        virtual int destory() = 0;
    protected:

    private:

};

#endif // EVENTMUTIIOAPI_H
