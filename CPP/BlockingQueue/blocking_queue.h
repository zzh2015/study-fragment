#ifndef BLOCKING_QUEUE_H_
#define BLOCKING_QUEUE_H_
#include <assert.h>
#include <deque>

#include "condition.h"

template <typename T>
class BlockingQueue {
    public:
        BlockingQueue() : mutexLock_(), notEmpty_(mutexLock_), queue_() {}
        ~BlockingQueue() {}

        void Put(const T &e) {
            MutexLockGuard lock(mutexLock_);
            queue_.push_back(e);
            notEmpty_.Notify();
        }

        void Put (T &&e) {
            MutexLockGuard lock(mutexLock_);
            queue_.push_back(std::move(e));
            notEmpty_.Notify();
        }

        T Tack() {
            MutexLockGuard lock(mutexLock_);
            while (queue_.empty()) {
                notEmpty_.Wait();
            }

            T front(std::move(queue_.front()));
            queue_.pop_front();

            return front;
        }

        size_t Size() const {
            MutexLockGuard lock(mutexLock_);
            return queue_.size();
        }
    private:
        mutable MutexLock mutexLock_;
        Condition notEmpty_;
        std::deque<T> queue_;
};

#endif
