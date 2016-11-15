/*************************************************************************
 > File Name: mutex_text.cc
 > Author:
 > Mail:
 > Created Time: 2016年11月15日 星期二 18时41分16秒
 ************************************************************************/
#include <iostream>
#include <string>

#include "../condition.h"
#include "../blocking_queue.h"

int main(int atgc,char **argv)
{
    MutexLock mutex;
    std::cout << "mutex" << std::endl;

    MutexLockGuard guard(mutex);
    std::cout << "guard" << std::endl;

    Condition cond(mutex);
    cond.WaitForSeconds(3.0);

    std::cout << "blocking queue" << std::endl;
    BlockingQueue<std::string> bq;
    bq.put("hello");

    return 0;
}
