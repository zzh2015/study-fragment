/*************************************************************************
 > File Name: object_pool.h
 > Author:
 > Mail:
 > Created Time: 2016年11月17日 星期四 15时20分21秒
 ************************************************************************/
#pragma once

#include <functional>
#include <memory>
#include <stdint.h>
#include <vector>

template <typename T>
class ObjectPool {
    using DeleterType = std::function<void(T*)>;
    public:
        ObjectPool(uint32_t objSize) : objSize_(objSize) {
            for (uint32_t i=0; i<objSize_; ++i) { //?
                objPool_.push_back(std::unique_ptr<T>(new T));
            }
        }

        ~ObjectPool() {}

        void Add(std::unique_ptr<T> t) {
            objPool_.push_back(t);
        }

        std::unique_ptr<T, DeleterType> Tack() {
            if (objPool_.empty()) {
                throw std::logic_error("no more obj");
            }

            std::unique_ptr<T, DeleterType> ptr(objPool_.back().release(), [this](T *t) \
                    {
                        objPool_.push_back(std::unique_ptr<T>(t));
                    });
            objPool_.pop_back();

            return std::move(ptr);
        }

        bool Empty() const {
            return objPool_.empty();
        }

        size_t Size() const {
            return objPool_.size();
        }

    private:
        uint32_t objSize_;
        std::vector<std::unique_ptr<T>> objPool_;
};
