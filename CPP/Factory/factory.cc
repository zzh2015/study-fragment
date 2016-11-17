/*************************************************************************
 > File Name: factory.cc
 > Author:
 > Mail:
 > Created Time: 2016年11月17日 星期四 11时39分21秒
 ************************************************************************/
#include "factory.h"

#include "product.h"

std::unique_ptr<Product> ProductFactory::ProductCreate(ProductOptions options)
{
    if (options == ProductOptions::ToyOption) {
            std::unique_ptr<Product> t1(new ToyProduct());
            return std::move(t1);
    } else if (options == ProductOptions::GameOption) {
            std::unique_ptr<Product> t2(new GameProduct());
            return std::move(t2);
    } else {
        return nullptr;
    }
}

