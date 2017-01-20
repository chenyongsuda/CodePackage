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

        static const int kEventRead = 1 << 0;
        static const int kEventWrite = 1 << 1;
        static const int kEventTimer = 1 << 2;
        static const int kEventError = 1 << 3;

        int fd;
        int type;
        EventCallBack read_cb;
        EventCallBack write_cb;
        EventCallBack timer_cb;
    protected:

    private:
};

#endif // EVENT_H
