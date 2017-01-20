#include <iostream>
#include <stdio.h>
#include <functional>
#include "TCPServer.h"
#include "ErrorNoDefine.h"

/*
using namespace std;
using namespace std::placeholders;

typedef function<int(int,int)> TestCallback;

int Test(int x, int y){
    return x + y;
}
*/
int main()
{
    /*TestCallback tcb = bind(Test,_1,_2);
    printf("%d\n",tcb(1,2));
    cout << "Hello world!" << endl;
    return 0;*/
    TCPServer server;
    int ret = TC_SUCCESS;
    if(TC_SUCCESS != (ret=server.init("127.0.0.1","80000"))){
        printf("TCPServer init Error");
        return ret;
    }

    if(TC_SUCCESS != (ret=server.start())){
        printf("TCPServer init Error");
        return ret;
    }
    return ret;
}
