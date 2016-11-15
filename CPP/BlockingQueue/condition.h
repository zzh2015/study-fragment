/*************************************************************************
 > File Name: condition.h
 > Author:
 > Mail:
 > Created Time: 2016年11月15日 星期二 19时17分12秒
 ************************************************************************/

#ifndef CONDITION_H_
#define CONDITION_H_

#include "mutex.h"

class Condition {
    public:
        Condition(MutexLock &mutexLock) : mutexLock_(mutexLock) {
            MCHECK(pthread_cond_init(&pcond_, nullptr));
        }

        ~Condition() {
            MCHECK(pthread_cond_destroy(&pcond_));
        }

        void Wait() {
            MutexLock::UnassignGuard ug(mutexLock_);
            MCHECK(pthread_cond_wait(&pcond_, mutexLock_.GetPthreadMutex()));
        }

        bool WaitForSeconds(double);

        void Notify() {
            MCHECK(pthread_cond_signal(&pcond_));
        }

        void NotifyAll() {
            MCHECK(pthread_cond_broadcast(&pcond_));
        }
    private:
        MutexLock &mutexLock_;
        pthread_cond_t pcond_;
};

#endif
