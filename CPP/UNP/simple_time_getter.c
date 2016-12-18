/*************************************************************************
 > File Name: simple_time_getter.c
 > Author:
 > Mail:
 > Created Time: 2016年12月13日 星期二 16时08分07秒
 ************************************************************************/
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#include <sys/socket.h>
#include <sys/types.h>

#include <arpa/inet.h>

#ifndef BUF_SIZE
#define BUF_SIZE 1024
#endif

int
main(int argc, char **argv)
{
    int sockfd;
    char recv_buf[BUF_SIZE];
    struct sockaddr_in server_addr;

    if (argc != 2) {
        perror("arg");
        return 1;
    }

    if ( (sockfd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("socket");
        return 1;
    }

    memset(&server_addr, 0, sizeof(struct sockaddr_in));
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(13);

    // x.x.x.x -> addr[4]
    if (inet_pton(AF_INET, argv[1], &server_addr.sin_addr) <= 0) {
        perror("inet_pton");
        return 1;
    }

    if (connect(sockfd, (struct sockaddr *)&(server_addr), sizeof(server_addr)) < 0) {
        perror("connect");
        return 1;
    }

    ssize_t r;
    while ( (r = read(sockfd, recv_buf, BUF_SIZE-1)) > 0) {
        recv_buf[r] = '\0';
        if (fputs(recv_buf, stdout) == EOF) {
            return 1;
        }
    }

    if (r < 0) {
        perror("read");
    }

    return 0;
}
