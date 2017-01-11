#ifndef EVENT_H
#define EVENT_H
#include <functional>

using namespace std;
using namespace std::placeholders;

typedef function<int(int,int,void *)> EventCallBack;

class Event
{
    public:
        Event();
        virtual ~Event();

        int fd;
        int type;
        EventCallBack read_cb;
        EventCallBack write_cb;
        EventCallBack timer_cb;
    protected:

    private:
};

#endif // EVENT_H
