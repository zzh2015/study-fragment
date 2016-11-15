/*************************************************************************
 > File Name: mutex.h
 > Author:
 > Mail:
 > Created Time: 2016年11月15日 星期二 17时13分12秒
 ************************************************************************/

#ifndef MUTEX_H_
#define MUTEX_H_

#include <assert.h>
#include <pthread.h>
#include <unistd.h>
#include <sys/syscall.h>

#define MCHECK(ret) ({__typeof__(ret) errnum=(ret);\
        assert(errnum == 0); (void) errnum;})

namespace {
    pid_t GetTid()
    {
        return static_cast<pid_t>(::syscall(SYS_gettid));
    }
}

class MutexLock {
    public:
        MutexLock() : holds_(0) {
            MCHECK(pthread_mutex_init(&mutex_, nullptr));
        }

        ~MutexLock() {
            assert(holds_ == 0);
            MCHECK(pthread_mutex_destroy(&mutex_));
        }

        bool IsLockedByThisThread() {
            return holds_ == GetTid();
        }

        void AssertLocked() {
            assert(IsLockedByThisThread());
        }

        void Lock() {
            MCHECK(pthread_mutex_lock(&mutex_));
            AssignHolder();
        }

        void Unlock() {
            UnassignHolder();
            MCHECK(pthread_mutex_unlock(&mutex_));
        }

        pthread_mutex_t *GetPthreadMutex() {
            return &mutex_;
        }

    private:
        friend class Condition;

        class UnassignGuard {
            public:
                UnassignGuard(MutexLock &owner) : owner_(owner) {
                    owner_.UnassignHolder();
                }

                ~UnassignGuard() {
                    owner_.AssignHolder();
                }
            private:
                MutexLock &owner_;
        };

        void AssignHolder() {
            holds_ = GetTid();
        }

        void UnassignHolder() {
            holds_ = 0;
        }

        pthread_mutex_t mutex_;
        pid_t holds_;
};

class MutexLockGuard {
    public:
        MutexLockGuard(MutexLock &mutexLock) : mutexLock_(mutexLock) {
            mutexLock_.Lock();
        }

        ~MutexLockGuard() {
            mutexLock_.Unlock();
        }
    private:
        MutexLock &mutexLock_;
};

#endif
