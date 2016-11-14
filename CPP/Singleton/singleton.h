/*************************************************************************
 > File Name: singleton.h
 > Author:
 > Mail:
 > Created Time: 2016年11月14日 星期一 13时41分47秒
 ************************************************************************/

#ifndef SINGLETON_H_
#define SINGLETON_H_

#include <pthread.h>
#include <stdlib.h>

template <typename T>
class Singleton {
    public:
        static T& Instance() {
            pthread_once(&ponce_, &Singleton::Init);
            return *values_;
        }

    private:
        Singleton();
        ~Singleton();

        static void Init()
        {
            values_ = new T();
            ::atexit(Destroy);
        }

        static void Destroy()
        {
            delete values_;
        }

        static pthread_once_t ponce_;
        static T *values_;
};

template<typename T> pthread_once_t Singleton<T>::ponce_ = PTHREAD_ONCE_INIT;
template<typename T> T* Singleton<T>::values_ = nullptr;
#endif
