#ifndef EVENTMUTIIOAPI_H
#define EVENTMUTIIOAPI_H


class EventMutiIOAPI
{
    public:
        EventMutiIOAPI();
        virtual ~EventMutiIOAPI();
        virtual int init(void) = 0;
        virtual int registe_event() = 0;
        virtual int unregiste_event() = 0;
        virtual int update_event() = 0;
    protected:

    private:
};

#endif // EVENTMUTIIOAPI_H
