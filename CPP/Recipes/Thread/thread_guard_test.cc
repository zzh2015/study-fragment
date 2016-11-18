/*************************************************************************
 > File Name: thread_guard_test.cc
 > Author:
 > Mail:
 > Created Time: 2016年11月18日 星期五 10时59分28秒
 ************************************************************************/
#include "thread_guard.h"

#include <stdint.h>
#include <stdio.h>

void func(uint8_t t) {
    for (uint8_t i=0; i<t; ++i) {
        printf("thread no.%u\n", t);
    }
}

int main()
{
    uint8_t times = 8;

    ThreadGuard tg(std::thread(func, times));

    func(7);

    return 0;
}
