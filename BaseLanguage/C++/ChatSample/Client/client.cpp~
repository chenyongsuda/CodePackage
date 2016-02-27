#include <sys/types.h>
#include <sys/socket.h>
#include <sys/ioctl.h>
#include <netinet/in.h>
#include <netinet/tcp.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <errno.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdint.h>		// define int8_t ...
#include <signal.h>
#include <unistd.h>

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>


int main(int argc, char *argv[])
{
	int sockfd;
	int n;
	struct sockaddr_in serv_addr;
	char buffer[256];

	sockfd = socket(AF_INET, SOCK_STREAM, 0);
	if (sockfd < 0)
	{
		printf("opening socket error! \n");
		exit(-1);
	}
	
	bzero(&serv_addr, sizeof(serv_addr));
	serv_addr.sin_family = AF_INET;
	serv_addr.sin_port = htons(10000);
	serv_addr.sin_addr.s_addr = inet_addr("127.0.0.1");
	if (connect(sockfd, (struct sockaddr *) &serv_addr, sizeof(serv_addr)) < 0)
	{
		printf("connecting error! \n");
		exit(-1);
	}
	
	//close(sockfd);
	
	
	while(1)
	{
		printf("Please enter the message: ");
		bzero(buffer, 256);
		fgets(buffer, 255, stdin);
		n = write(sockfd, buffer, strlen(buffer)-1);
		if (n < 0)
			printf("error writing to socket! \n");
	}
	bzero(buffer, 256);
	n = read(sockfd, buffer, 255);
	if (n < 0)
	printf("error reading from socket! \n");
	printf("%s \n",buffer);
	return 0;
}
