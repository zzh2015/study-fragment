/*************************************************************************
 > File Name: thread_product_test.cc
 > Author:
 > Mail:
 > Created Time: 2016年11月18日 星期五 11时31分31秒
 ************************************************************************/
#include <algorithm>
#include <stdint.h>
#include <stdio.h>
#include <thread>
#include <vector>

void func(uint8_t t) {
    for (uint8_t i=0; i<2; ++i) {
        printf("thread work no,%u\n", t);
    }
}

int main() {
    std::vector<std::thread> threads; //vector 移动敏感

    for (auto i=0; i<16; ++i) {
        threads.push_back(std::thread(func, i));
    }

    // std::for_each (threads.begin(), threads.end(), std::mem_fn(&std::thread::join));
    std::for_each (threads.begin(), threads.end(), [](std::thread &th){
                th.join();
            });

    return 0;
}
