#include "EventCenter.h"
#include "EventMutiIOAPI.h"

EventCenter::EventCenter()
: regiested_events_(),ioAPI_(NULL),stop_(false),ready_events_()
{
    //ctor
}

EventCenter::~EventCenter()
{
    //dtor
}

int
EventCenter::init(EventMutiIOAPI* ioAPI){
    int ret = TC_SUCCESS;
    ioAPI_ = ioAPI;
    if (TC_SUCCESS != ioAPI_->init()){
        ret = TC_EPOLL_ERR;
    }
    return ret;
}


int
EventCenter::registe_event(Event *ev){
    int ret = TC_SUCCESS;
    if((ev->type&Event::kEventRead) || (ev->type&Event::kEventWrite)){
        if(TC_SUCCESS != (ret = ioAPI_->registe_event(ev->fd,ev->type))){
            //Error
            return ret;
        }
        regiested_events_[ev->fd] = ev;
    }
    return ret;
}

int
EventCenter::unregiste_event(Event *ev){
    int ret = TC_SUCCESS;
    if((ev->type&Event::kEventRead) || (ev->type&Event::kEventWrite)){
        if(TC_SUCCESS != (ret = ioAPI_->unregiste_event(ev->fd,ev->type))){
            //Error
            return ret;
        }
        regiested_events_.erase(ev->fd);
    }
    return ret;
}

int
EventCenter::update_event(Event *ev){
    int ret = TC_SUCCESS;
    if((ev->type&Event::kEventRead) || (ev->type&Event::kEventWrite)){
        if(TC_SUCCESS != (ret = ioAPI_->update_event(ev->fd,ev->type))){
            //Error
            return ret;
        }
        regiested_events_[ev->fd] = ev;
    }
    return ret;
}

void
EventCenter::add_ready_event(int fd, int event_type){
    Event ev = *regiested_events_[fd];
    ev.type = event_type;
    ready_events_.push_back(ev);
}

int
EventCenter::loop(){
    int ret = TC_SUCCESS;
    while(!stop_){
        //1. Loop inner
        if (TC_SUCCESS != (ret = ioAPI_->loop())){
            return ret;
        }
        //2. Call Function
        for (auto e = ready_events_.cbegin(); e != ready_events_.cend(); ++e) {
            Event* ev = regiested_events_[e->fd];
            if(ev->type & Event::kEventRead){
                ev->read_cb(ev->fd,Event::kEventRead,NULL);
            }
            if(ev->type & Event::kEventWrite){
                ev->write_cb(ev->fd,Event::kEventWrite,NULL);
            }
        }
        //Clean events
        ready_events_.clear();
    }
    return ret;
}
