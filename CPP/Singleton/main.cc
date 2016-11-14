/*************************************************************************
  > File Name: main.cc
  > Author:
  > Mail:
  > Created Time: 2016年11月14日 星期一 13时54分58秒
 ************************************************************************/
#include <iostream>

#include "singleton.h"

class Test {
    public:
        Test() : name_("Test::") {}
        ~Test() {}

        void PrintInfo() { std::cout << name_ << " Test Singleton" << std::endl; }
    private:
        std::string name_;
};

int main(int argc, char **argv)
{
    Singleton<Test>::Instance().PrintInfo();

    return 0;
}
