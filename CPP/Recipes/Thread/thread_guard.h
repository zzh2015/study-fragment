/*************************************************************************
 > File Name: thread_guard.h
 > Author:
 > Mail:
 > Created Time: 2016年11月18日 星期五 10时53分56秒
 ************************************************************************/
#pragma once

#include <thread>

class ThreadGuard {
    public:
        ThreadGuard(std::thread thread) : thread_(std::move(thread)) {
            if (!thread_.joinable()) {
                throw std::logic_error("no thread");
            }
        }

        ~ThreadGuard() {
            thread_.join();
        }

        ThreadGuard(const ThreadGuard &) = delete;
        ThreadGuard &operator=(const ThreadGuard &) = delete;
    private:
            std::thread thread_;
};
