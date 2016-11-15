/*************************************************************************
 > File Name: condition.cc
 > Author:
 > Mail:
 > Created Time: 2016年11月15日 星期二 19时45分02秒
 ************************************************************************/
#include "condition.h"

#include <errno.h>
#include <stdint.h>
#include <time.h>

bool Condition::WaitForSeconds(double seconds)
{
    struct timespec absTime;
    clock_gettime(CLOCK_REALTIME, &absTime);

    const int64_t kNanoSeconsPerSencond = 1e9;
    int64_t nanoSeconds = static_cast<int64_t>(seconds * kNanoSeconsPerSencond);

    absTime.tv_sec += static_cast<time_t>((absTime.tv_nsec + nanoSeconds) / kNanoSeconsPerSencond);
    absTime.tv_nsec = static_cast<long>((absTime.tv_nsec + nanoSeconds) % kNanoSeconsPerSencond);

    MutexLock::UnassignGuard ug(mutexLock_);

    return ETIMEDOUT == pthread_cond_timedwait(&pcond_, mutexLock_.GetPthreadMutex(), &absTime);
}
