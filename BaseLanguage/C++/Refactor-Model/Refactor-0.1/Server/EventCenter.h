#ifndef EVENTCENTER_H
#define EVENTCENTER_H
#include "Event.h"
#include <unordered_map>
#include <list>
#include "ErrorNoDefine.h"
using namespace std;

class EventMutiIOAPI;
class EventCenter
{
    public:
        EventCenter();
        virtual ~EventCenter();
        int init(EventMutiIOAPI* ioAPI);
        int registe_event(Event* ev);
        int unregiste_event(Event* ev);
        int update_event(Event* ev);
        int loop();
        void add_ready_event(int fd, int event_type);
    protected:
    private:
        EventMutiIOAPI*             ioAPI_;
        unordered_map<int, Event*>   regiested_events_;
        list<Event> ready_events_;
        bool stop_;
};

#endif // EVENTCENTER_H
