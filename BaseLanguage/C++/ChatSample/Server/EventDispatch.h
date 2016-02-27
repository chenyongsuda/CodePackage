#ifndef __EVENT_DISPATCH_H__
#define __EVENT_DISPATCH_H__

class EventDispatch
{
	private:
		EventDispatch();
		virtual ~EventDispatch();
	
	public:
		void AddEvent(int fd, int event_type);
		void RemoveEvent(int fd);
		void AddTimer();
		void RemoveTimer();
		void StartEventLoop(int wait_time = 100);
		static EventDispatch* GetInstance();
		void Init();
	private:
		int m_epfd;
		static EventDispatch* instance;
};

#endif