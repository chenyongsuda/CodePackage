/*************************************************************************
	> File Name: client.c
	> Author: fucang_zxx
	> Mail: fucang_zxx@163.com
	> Created Time: Sun 24 Jul 2016 03:45:21 PM CST
 ************************************************************************/

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <unistd.h>

const int port = 80000;
const char* ip = "127.0.0.1";

int main()
{
	  //创建套接字,即创建socket
	  int clt_sock = socket(AF_INET, SOCK_STREAM, 0);
	  if(clt_sock < 0)
	  {
		  //创建失败
		  perror("socket");
		  return 1;
	  }

	  //绑定信息，即命名socket
	  struct sockaddr_in addr;
	  addr.sin_family = AF_INET;
	  addr.sin_port = htons(port);
	  //inet_addr函数将用点分十进制字符串表示的IPv4地址转化为用网络
 	  //字节序整数表示的IPv4地址
	  addr.sin_addr.s_addr = inet_addr(ip);

	  //不需要监听

	  //发起连接
//	  struct sockaddr_in peer;
	  socklen_t addr_len = sizeof(addr);
	  int connect_fd = connect(clt_sock, (struct sockaddr*)&addr, addr_len);
	  if(connect_fd < 0)
	  {
		  perror("connect");
		  return 2;
	  }
	  char buf[1024];

	  while(1)
	  {
		  memset(buf, '\0', sizeof(buf));
		  printf("client please enter: ");
		  fflush(stdout);
		  ssize_t size = read(0, buf, sizeof(buf) - 1);
		  if(size > 0)
		  {
			  buf[size - 1] = '\0';
		  }
		  else if(size == 0)
		  {
			  printf("read is done...\n");
			  break;
		  }
		  else
		  {
			  perror("read");
			  return 4;
		  }
		 // printf("client: %s\n", buf);
          write(clt_sock, buf, strlen(buf));
		  size = read(clt_sock, buf, sizeof(buf));
		  if(size > 0)
		  {
		 	  buf[size] = '\0';
		  }
		  else if(size == 0)
		  {
			  printf("read is done...\n");
			  break;
		  }
		  else
		  {
			  perror("read");
			  return 5;
		  }
		  printf("server: %s\n", buf);
	   }
	  close(clt_sock);
	  return 0;
}
