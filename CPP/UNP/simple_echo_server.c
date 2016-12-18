/*************************************************************************
 > File Name: simple_echo_server.c
 > Author:
 > Mail:
 > Created Time: 2016年12月18日 星期日 15时03分31秒
 ************************************************************************/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <errno.h>
#include <unistd.h>

#include <sys/socket.h>
#include <sys/types.h>
#include <arpa/inet.h>

#ifndef SERVER_PORT
#define SERVER_PORT 9877 // general 1023 < port < 49152
#endif

#ifndef BUF_SIZE
#define BUF_SIZE 1024
#endif

// write func
static ssize_t writen(int fd, const void *vptr, size_t count);
static void str_echo(int fd);

int
main(int argc, char *argv[])
{
    int listenfd, connfd;
    pid_t child_pid;
    struct sockaddr_in cliaddr, seraddr;

    // ipv4, 0 mean sys will auto set protocol base family & type
    if ( (listenfd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("socket");
        return 1;
    }
    // fill server addr info
    memset(&seraddr, 0, sizeof(seraddr));
    seraddr.sin_family = AF_INET;
    seraddr.sin_addr.s_addr = htonl(INADDR_ANY);
    seraddr.sin_port = htons(SERVER_PORT);

    if (bind(listenfd, (struct sockaddr *)&(seraddr), sizeof(seraddr)) < 0) {
        perror("bind");
        return 1;
    }

    if (listen(listenfd, SOMAXCONN) < 0) {
        perror("listen");
        return 1;
    }

    socklen_t clilen;
    // while(1)
    for ( ; ; ) {
        clilen = sizeof(cliaddr);
        connfd = accept(listenfd, (struct sockaddr *)&(cliaddr), &clilen);
        if (connfd == -1) {
            perror("accept");
            continue;
        }
        if ( (child_pid = fork()) == 0) {
            close(listenfd);
            //
            str_echo(connfd);
            close(connfd);
            exit(0);
        }
        close(connfd);
    }

    return 0;
}

static ssize_t
writen(int fd, const void *vptr, size_t count)
{
    size_t nleft;
    ssize_t nwrite;
    const char *ptr;

    ptr = vptr;
    nleft = count;

    while (nleft > 0) {
        if ( (nwrite = write(fd, ptr, nleft)) <= 0) {
            if (nwrite < 0 && errno == EINTR) {
                nwrite = 0;
            } else {
                return -1;
            }
        }
        nleft -= (size_t)nwrite;
        ptr += nwrite;
    }

    return (ssize_t)count;
}

static void
str_echo(int fd)
{
    ssize_t n;
    char buf[BUF_SIZE];

again:
    while ( (n = read(fd, buf, BUF_SIZE)) > 0) {
        writen(fd, buf, (size_t)n);
    }

    if (n < 0 && errno == EINTR) {
        goto again;
    } else if (n < 0) {
        perror("read");
        return;
    }
}
