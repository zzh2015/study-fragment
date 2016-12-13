#pragma once

#include <unistd.h>

enum class IODataType {
    UdpData,
    TcpData
};

class Sockets {
    public:
        Sockets(int fd) : fd_(fd) {}
        virtual ~Sockets() { ::close(fd_); };

        ssize_t ReadFd(char *, IODataType);
        ssize_t WriteFd(const char *, IODataType);
    private:
        int fd_;
};
