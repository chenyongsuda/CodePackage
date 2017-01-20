#ifndef EVENTCENTER_H
#define EVENTCENTER_H
#include "Event.h"

class EventCenter
{
    public:
        EventCenter();
        virtual ~EventCenter();
        int registe_event(Event ev);
        int unregiste_event(Event ev);
        int update_event(Event ev);
        int loop();
    protected:
    private:
};

#endif // EVENTCENTER_H
