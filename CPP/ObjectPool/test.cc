/*************************************************************************
 > File Name: test.cc
 > Author:
 > Mail:
 > Created Time: 2016年11月17日 星期四 15时50分09秒
 ************************************************************************/
#include "object_pool.h"

#include <iostream>
#include <string>

class Test {
    public:
        Test() {}
        ~Test() {}

        void ShowInfo(const std::string &tag) { std::cout << tag << " Test..." << std::endl; }
};

int main()
{
    ObjectPool<Test> pool(8);

    auto t1 = pool.Tack();
    t1->ShowInfo("obj::t1");
    {
        auto t2 = pool.Tack();
        t1->ShowInfo("obj::t2");
        std::cout << pool.Size() <<std::endl;
    }
    std::cout << pool.Size() <<std::endl;

    return 0;
}
